package game

import rl "github.com/gen2brain/raylib-go/raylib"

func Update() {
	running = !rl.WindowShouldClose()

	if rl.IsKeyDown(rl.KeyD) {
		playerDest.X += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyW) {
		playerDest.Y -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyS) {
		playerDest.Y += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyA) {
		playerDest.X -= playerSpeed
	}

	if rl.IsKeyDown(rl.KeyUp) {
		playerDest.Height += playerSpeed
		playerDest.Width += playerSpeed
	}

	if rl.IsKeyDown(rl.KeyDown) {
		playerDest.Height -= playerSpeed
		playerDest.Width -= playerSpeed
	}
}
