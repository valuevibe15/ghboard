// ui/heatmap.go (stub — replaced in Task 7)
package ui

import (
	"github.com/Null-Phnix/ghboard/api"
	tea "github.com/charmbracelet/bubbletea"
)

type HeatmapModel struct{ gql *api.GraphQLClient }

func NewHeatmapModel(gql *api.GraphQLClient) HeatmapModel { return HeatmapModel{gql: gql} }
func (m HeatmapModel) Init() tea.Cmd                      { return nil }
func (m HeatmapModel) Update(msg tea.Msg) (HeatmapModel, tea.Cmd) {
	return m, nil
}
func (m HeatmapModel) View(w, h int) string { return "heatmap coming soon" }
