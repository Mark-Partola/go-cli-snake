package components

import "github.com/charmbracelet/lipgloss"

type FilledCell struct {
	color   string
	content string
}

func NewFilledCell() Component {
	return &FilledCell{
		color:   "#663399",
		content: "██",
	}
}

func (f *FilledCell) Render(any) string {
	return render(f.color, f.content)
}

type EmptyCell struct {
	color   string
	content string
}

func NewEmptyCell() Component {
	return &EmptyCell{
		color:   "#000000",
		content: "░░",
	}
}

func (e *EmptyCell) Render(any) string {
	return render(e.color, e.content)
}

type AccentCell struct {
	color   string
	content string
}

func NewAccentCell() Component {
	return &AccentCell{
		color:   "#ff9100",
		content: "██",
	}
}

func (a *AccentCell) Render(any) string {
	return render(a.color, a.content)
}

func render(color, content string) string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color(color)).Render(content)
}
