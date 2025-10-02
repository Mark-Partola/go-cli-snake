package components

type Component interface {
	Render(props any) string
}
