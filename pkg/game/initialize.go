package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Initialize() {
	// This needs to come before any textures are loaded
	rl.InitWindow(screenWidth, screenHeight, "raylib-go course")

	rl.SetExitKey(rl.KeyF10)
	rl.SetTargetFPS(targetFps)

	playerSrc = rl.NewRectangle(0, 0, 48, 48)

	playerDest = rl.NewRectangle(100, 100, 100, 100)

	grassTexture = rl.LoadTexture("res/sproutlands/Tilesets/Grass.png")
	playerTexture = rl.LoadTexture("res/sproutlands/Characters/character_spritesheet.png")
}
