package ecs

import "github.com/EngoEngine/ecs"

type BackgroundEntity struct {
	ecs.BasicEntity
	*TextureComponent
	*RenderComponent
	Whatever string
}
