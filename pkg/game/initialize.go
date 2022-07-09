package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Initialize() {
	rl.InitWindow(screenWidth, screenHeight, "raylib-go course")
	rl.SetExitKey(rl.KeyF10)
	rl.SetTargetFPS(targetFps)
}
