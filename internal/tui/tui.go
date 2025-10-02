package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"fuzzy-snake/internal/tui/components"
)

type Router struct {
	Screens map[string]components.Component
	Route   func() string
}

type Config struct {
	Rps        int
	OnTick     func()
	OnKeyPress func(string)
	Router     Router
}

type model struct {
	config Config
	ticker *ticker
	route  string
}

func instantiate(config Config) model {
	return model{
		config: config,
		ticker: NewTicker(config.Rps),
		route:  config.Router.Route(),
	}
}

func (m model) Init() tea.Cmd {
	return m.ticker.tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickCmd:
		m.route = m.config.Router.Route()
		m.config.OnTick()
		return m, m.ticker.tick()
	case tea.KeyMsg:
		keys := msg.String()
		switch keys {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		default:
			m.config.OnKeyPress(keys)
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.config.Router.Screens[m.route].Render(nil)
}

func Run(config Config) {
	p := tea.NewProgram(instantiate(config), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
