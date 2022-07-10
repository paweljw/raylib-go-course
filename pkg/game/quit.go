package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
)

func Quit() {
	ecs.QuitWorld(world)
	rl.CloseAudioDevice()
	rl.CloseWindow() // This needs to come after all textures are unloaded
}
