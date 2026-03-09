// ui/notifications.go (stub — replaced in Task 9)
package ui

import (
	"github.com/Null-Phnix/ghboard/api"
	tea "github.com/charmbracelet/bubbletea"
)

type NotificationsModel struct{ rest *api.RESTClient }

func NewNotificationsModel(rest *api.RESTClient) NotificationsModel {
	return NotificationsModel{rest: rest}
}
func (m NotificationsModel) Init() tea.Cmd                           { return nil }
func (m NotificationsModel) Update(msg tea.Msg) (NotificationsModel, tea.Cmd) { return m, nil }
func (m NotificationsModel) View(w, h int) string                    { return "notifications coming soon" }
