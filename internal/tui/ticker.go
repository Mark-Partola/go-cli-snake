package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickCmd time.Time

type ticker struct {
	ms time.Duration
}

func NewTicker(rps int) *ticker {
	return &ticker{
		ms: time.Duration(1000 / rps),
	}
}

func (t *ticker) tick() tea.Cmd {
	return tea.Tick(time.Millisecond*t.ms, func(t time.Time) tea.Msg {
		return tickCmd(t)
	})
}
