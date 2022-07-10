package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
	"log"
)

func Quit() {
	for _, system := range world.Systems() {
		switch t := system.(type) {
		case ecs.QuittableSystemFace:
			log.Printf("Quitting out of system: %s", t)
			system.(ecs.QuittableSystemFace).Quit()
		default:
			log.Println(t)
		}
	}

	rl.CloseWindow() // This needs to come after all textures are unloaded
}
