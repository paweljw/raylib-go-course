package ecs

import (
	"github.com/EngoEngine/ecs"
	"log"
)

func NewWorld() *ecs.World {
	w := &ecs.World{}
	AddBackgroundFloodSystemToWorld(w)
	AddRenderSystemToWorld(w)
	AddInputSystemToWorld(w)
	AddMusicSystemToWorld(w)

	return w
}

func QuitWorld(world *ecs.World) {
	for _, system := range world.Systems() {
		switch t := system.(type) {
		case QuittableSystemFace:
			log.Printf("Quitting out of system: %s", t)
			system.(QuittableSystemFace).Quit()
		default:
			log.Println(t)
		}
	}
}
