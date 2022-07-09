package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Quit() {
	rl.UnloadTexture(grassTexture)
	rl.UnloadTexture(playerTexture)

	rl.CloseWindow() // This needs to come after all textures are unloaded
}
