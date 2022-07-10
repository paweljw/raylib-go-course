package ecs

import "github.com/EngoEngine/ecs"

type PlayerEntity struct {
	ecs.BasicEntity
	*TextureComponent
	*RenderComponent
	*InputComponent
	*CameraComponent
	*AnimationComponent
}
