package ecs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureComponent struct {
	Src           rl.Rectangle
	TexturePath   string
	Texture       rl.Texture2D
	TextureLoaded bool
}

func (r *TextureComponent) GetTextureComponent() *TextureComponent {
	return r
}

type TextureFace interface {
	GetTextureComponent() *TextureComponent
}
