package ecs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RenderComponent struct {
	Dest     rl.Rectangle
	Rotation float32
	Tint     rl.Color
}

func (r *RenderComponent) GetRenderComponent() *RenderComponent {
	return r
}

type RenderFace interface {
	GetRenderComponent() *RenderComponent
}
