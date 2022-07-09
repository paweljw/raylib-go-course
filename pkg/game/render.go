package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Render() {
	rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)

	drawScene()

	rl.EndDrawing()
}
