package game

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	running         = true
	backgroundColor = rl.NewColor(147, 211, 196, 255)
)

func IsRunning() bool {
	return running
}
