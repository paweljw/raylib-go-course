package game

import rl "github.com/gen2brain/raylib-go/raylib"

func drawScene() {
	rl.DrawTexture(grassTexture, 10, 10, rl.White)

	rl.DrawTexturePro(
		playerTexture,
		playerSrc, playerDest,
		rl.NewVector2(playerDest.Width, playerDest.Height),
		0,
		rl.White,
	)
}
