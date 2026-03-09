// ui/app.go
package ui

import (
	"github.com/Null-Phnix/ghboard/api"
	"github.com/Null-Phnix/ghboard/config"
	"github.com/Null-Phnix/ghboard/store"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tab int

const (
	tabHeatmap tab = iota
	tabStars
	tabNotifications
)

var tabNames = []string{"  Heatmap  ", "  Stars  ", "  Notifications  "}

type App struct {
	cfg           *config.Config
	rest          *api.RESTClient
	gql           *api.GraphQLClient
	tags          *store.Store
	activeTab     tab
	width         int
	height        int
	heatmap       HeatmapModel
	stars         StarsModel
	notifications NotificationsModel
	showHelp      bool
}

func NewApp(cfg *config.Config) *App {
	rest := api.NewRESTClient(cfg.Token, "")
	gql := api.NewGraphQLClient(cfg.Token, "")
	tags := store.New(store.DefaultPath())

	return &App{
		cfg:           cfg,
		rest:          rest,
		gql:           gql,
		tags:          tags,
		activeTab:     tabHeatmap,
		heatmap:       NewHeatmapModel(gql),
		stars:         NewStarsModel(rest, tags),
		notifications: NewNotificationsModel(rest),
	}
}

func (a *App) Init() tea.Cmd {
	return a.heatmap.Init()
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height

	case tea.KeyMsg:
		if a.showHelp {
			a.showHelp = false
			return a, nil
		}
		switch msg.String() {
		case "q", "ctrl+c":
			a.tags.Save()
			return a, tea.Quit
		case "?":
			a.showHelp = true
			return a, nil
		case "1":
			a.activeTab = tabHeatmap
			return a, a.heatmap.Init()
		case "2":
			a.activeTab = tabStars
			return a, a.stars.Init()
		case "3":
			a.activeTab = tabNotifications
			return a, a.notifications.Init()
		case "tab":
			a.activeTab = (a.activeTab + 1) % 3
		}
	}

	var cmd tea.Cmd
	switch a.activeTab {
	case tabHeatmap:
		a.heatmap, cmd = a.heatmap.Update(msg)
	case tabStars:
		a.stars, cmd = a.stars.Update(msg)
	case tabNotifications:
		a.notifications, cmd = a.notifications.Update(msg)
	}
	return a, cmd
}

var (
	activeTabStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF7F")).
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#00FF7F"))

	inactiveTabStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#666666"))

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Padding(0, 1)
)

func (a *App) View() string {
	if a.showHelp {
		return a.helpView()
	}

	// Tab bar
	tabs := ""
	for i, name := range tabNames {
		if tab(i) == a.activeTab {
			tabs += activeTabStyle.Render(name)
		} else {
			tabs += inactiveTabStyle.Render(name)
		}
	}

	tabBar := lipgloss.NewStyle().
		BorderBottom(true).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#333333")).
		Width(a.width).
		Render(tabs)

	statusBar := statusBarStyle.Render("1/2/3: switch tabs  •  Tab: cycle  •  ?: help  •  q: quit")

	contentHeight := a.height - 4
	var content string
	switch a.activeTab {
	case tabHeatmap:
		content = a.heatmap.View(a.width, contentHeight)
	case tabStars:
		content = a.stars.View(a.width, contentHeight)
	case tabNotifications:
		content = a.notifications.View(a.width, contentHeight)
	}

	return tabBar + "\n" + content + "\n" + statusBar
}

func (a *App) helpView() string {
	help := `
  ghboard — GitHub Terminal Dashboard

  GLOBAL
    1 / 2 / 3   Switch tabs
    Tab         Cycle tabs
    ?           Toggle help
    q           Quit

  HEATMAP  (tab 1)
    ← → ↑ ↓    Navigate days  (or h j k l)
    [ ]         Previous / next year
    ctrl+r      Refresh

  STARS  (tab 2)
    ↑ ↓         Move cursor  (or j k)
    /           Fuzzy search
    t           Add / edit tags  (comma-separated)
    f           Filter by tag or language
    u           Unstar  (confirm with y)
    o           Open in browser
    ctrl+r      Refresh

  NOTIFICATIONS  (tab 3)
    ↑ ↓         Move cursor  (or j k)
    r           Mark as read
    R           Mark ALL as read
    d           Dismiss
    o           Open in browser
    f           Cycle type filter
    ctrl+r      Refresh

  Press any key to close help
`
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render(help)
}
