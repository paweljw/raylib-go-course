package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

func NewWorld() *ecs.World {
	w := &ecs.World{}
	AddBackgroundFloodSystemToWorld(w)
	AddInputSystemToWorld(w)
	AddCameraSystemToWorld(w)
	AddRenderSystemToWorld(w)
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

func GetWorldCamera(world *ecs.World) *rl.Camera2D {
	for _, system := range world.Systems() {
		switch system.(type) {
		case CameraSystemFace:
			return system.(CameraSystemFace).GetCamera()
		}
	}

	return nil
}
