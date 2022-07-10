package ecs

import "github.com/EngoEngine/ecs"

func NewWorld() *ecs.World {
	w := &ecs.World{}
	AddBackgroundFloodSystemToWorld(w)
	AddRenderSystemToWorld(w)
	AddInputSystemToWorld(w)

	return w
}
