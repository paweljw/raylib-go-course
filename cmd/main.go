package main

import (
	"github.com/paweljw/raylib-go-course/pkg/game"
)

func main() {
	game.Initialize()

	for game.IsRunning() {
		game.Update()
	}

	game.Quit()
}
