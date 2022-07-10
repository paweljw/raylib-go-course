package ecs

import "github.com/EngoEngine/ecs"

type MusicEntity struct {
	ecs.BasicEntity
	*MusicComponent
}
