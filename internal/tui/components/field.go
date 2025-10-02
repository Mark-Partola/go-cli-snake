package components

import (
	"log"
	"strings"
)

type Field struct {
	components map[string]Component
}

func NewField() Component {
	return &Field{
		components: map[string]Component{
			"empty":  NewEmptyCell(),
			"filled": NewFilledCell(),
			"accent": NewAccentCell(),
		},
	}
}

func (m *Field) Render(props any) string {
	field, ok := props.([][]int)
	if !ok {
		log.Fatal("TypeError: props must be of type [][]int")
	}

	var builder strings.Builder
	for i := range field {
		for j := range field[i] {
			switch field[i][j] {
			case 1:
				builder.WriteString(m.components["filled"].Render(nil))
			case 2:
				builder.WriteString(m.components["accent"].Render(nil))
			default:
				builder.WriteString(m.components["empty"].Render(nil))
			}
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
