package ecs

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MusicComponent struct {
	MusicPath   string
	Playing     bool
	music       rl.Music
	musicLoaded bool
}

func (r *MusicComponent) GetMusicComponent() *MusicComponent {
	return r
}

type MusicFace interface {
	GetMusicComponent() *MusicComponent
}
