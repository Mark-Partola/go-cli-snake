package game

type field [][]int

func NewField(dimensions dimensions) field {
	var field [][]int

	for range dimensions.h {
		var row []int
		for range dimensions.w {
			row = append(row, 0)
		}
		field = append(field, row)
	}

	return field
}
