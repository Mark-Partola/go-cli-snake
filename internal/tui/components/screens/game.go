package screens

import (
	"fuzzy-snake/internal/tui/components"

	"github.com/Broderick-Westrope/charmutils"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type screen struct {
	field      [][]int
	components map[string]components.Component
	viewport   viewport.Model
}

func NewGameScreen(field [][]int) components.Component {
	padding := 2
	height := len(field)
	width := len(field[0])
	viewport := viewport.New(width*2+padding, height+padding)
	viewport.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Background(lipgloss.Color("#000"))

	return screen{
		field:      field,
		components: map[string]components.Component{"field": components.NewField()},
		viewport:   viewport,
	}
}

func (s screen) Render(any) string {
	field := s.components["field"].Render(s.field)

	output, err := charmutils.Overlay(s.viewport.View(), field, 1, 1, true)
	if err != nil {
		return err.Error()
	}

	return output
}
