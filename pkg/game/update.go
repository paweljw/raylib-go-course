package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	running = !rl.WindowShouldClose()

	rl.BeginDrawing()

	world.Update(rl.GetFrameTime())

	rl.DrawFPS(0, 0)

	rl.EndDrawing()
}
