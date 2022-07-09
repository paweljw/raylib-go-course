package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Update() {
	running = !rl.WindowShouldClose()
}
