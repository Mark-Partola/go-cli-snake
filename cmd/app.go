package main

import (
	"fuzzy-snake/internal/game"
	"fuzzy-snake/internal/tui"
	"fuzzy-snake/internal/tui/components"
	"fuzzy-snake/internal/tui/components/screens"
)

func main() {
	state := game.New()

	tui.Run(tui.Config{
		Router: tui.Router{
			Screens: map[string]components.Component{
				"game":  screens.NewGameScreen(state.Field()),
				"score": components.NewGameOver(state.Score()),
			},
			Route: func() string {
				if state.GameOver() {
					return "score"
				}
				return "game"
			},
		},
		OnTick: state.Update,
		Rps:    15,
		OnKeyPress: func(key string) {
			switch key {
			case "up":
				state.TurnUp()
			case "down":
				state.TurnDown()
			case "left":
				state.TurnLeft()
			case "right":
				state.TurnRight()
			}
		},
	})
}
