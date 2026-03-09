// ui/stars.go (stub — replaced in Task 8)
package ui

import (
	"github.com/Null-Phnix/ghboard/api"
	"github.com/Null-Phnix/ghboard/store"
	tea "github.com/charmbracelet/bubbletea"
)

type StarsModel struct {
	rest *api.RESTClient
	tags *store.Store
}

func NewStarsModel(rest *api.RESTClient, tags *store.Store) StarsModel {
	return StarsModel{rest: rest, tags: tags}
}
func (m StarsModel) Init() tea.Cmd                    { return nil }
func (m StarsModel) Update(msg tea.Msg) (StarsModel, tea.Cmd) { return m, nil }
func (m StarsModel) View(w, h int) string             { return "stars coming soon" }
