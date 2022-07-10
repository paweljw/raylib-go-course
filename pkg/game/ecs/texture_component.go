package ecs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureComponent struct {
	Src           rl.Rectangle
	TexturePath   string
	texture       rl.Texture2D
	textureLoaded bool
}

func (r *TextureComponent) GetTextureComponent() *TextureComponent {
	return r
}

type TextureFace interface {
	GetTextureComponent() *TextureComponent
}
