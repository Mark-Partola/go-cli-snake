package game

import (
	"fuzzy-snake/internal/system/utils"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

type game struct {
	direction direction
	intention direction
	field     field
	body      []Point
	food      Point
}

type dimensions struct {
	w int
	h int
}

type Point struct {
	x int
	y int
}

func New() game {
	return game{
		direction: right,
		field:     NewField(dimensions{50, 20}),
		body:      []Point{{2, 5}, {3, 5}, {4, 5}, {5, 5}},
		food:      Point{10, 10},
		intention: right,
	}
}

func (g *game) dimensions() dimensions {
	return dimensions{
		w: len(g.field[0]),
		h: len(g.field),
	}
}

func (g *game) Update() {
	tail := g.body[0]

	if !g.alive() {
		return
	}

	g.tryToEat()
	g.move()

	g.field[tail.y][tail.x] = 0
	for i := 0; i < len(g.body); i++ {
		g.field[g.body[i].y][g.body[i].x] = 1
	}
	g.field[g.food.y][g.food.x] = 2
}

func (g *game) alive() bool {
	head := g.body[len(g.body)-1]
	for i := 0; i < len(g.body)-1; i++ {
		if g.body[i].x == head.x && g.body[i].y == head.y {
			return false
		}
	}
	return true
}

func (g *game) move() {
	for i := 0; i < len(g.body)-1; i++ {
		g.body[i].x = g.body[i+1].x
		g.body[i].y = g.body[i+1].y
	}

	g.direction = g.intention
	head := g.body[len(g.body)-1]
	switch g.direction {
	case right:
		head.x = utils.Remainder(head.x+1, g.dimensions().w)
	case left:
		head.x = utils.Remainder(head.x-1, g.dimensions().w)
	case up:
		head.y = utils.Remainder(head.y-1, g.dimensions().h)
	case down:
		head.y = utils.Remainder(head.y+1, g.dimensions().h)
	}

	g.body[len(g.body)-1] = head
}

func (g *game) tryToEat() {
	head := g.body[len(g.body)-1]
	if head.x == g.food.x && head.y == g.food.y {
		g.body = append(g.body, Point{head.x, head.y})
		g.addFood()
	}
}

func (g *game) addFood() {
	g.food.x = utils.FastRand(g.dimensions().w)
	g.food.y = utils.FastRand(g.dimensions().h)

	for i := range g.body {
		if g.body[i].x == g.food.x && g.body[i].y == g.food.y {
			g.addFood()
		}
	}
}

func (g *game) TurnRight() {
	if g.direction != left {
		g.intention = right
	}
}

func (g *game) TurnLeft() {
	if g.direction != right {
		g.intention = left
	}
}

func (g *game) TurnUp() {
	if g.direction != down {
		g.intention = up
	}
}

func (g *game) TurnDown() {
	if g.direction != up {
		g.intention = down
	}
}

func (g *game) Field() field {
	return g.field
}

func (g *game) GameOver() bool {
	return !g.alive()
}

func (g *game) Score() int {
	return len(g.body)
}
