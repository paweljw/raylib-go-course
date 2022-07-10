package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Animationable interface {
	ecs.BasicFace
	TextureFace
	AnimationFace
}

type animationSystemEntity struct {
	ecs.BasicEntity
	*TextureComponent
	*AnimationComponent
}

type AnimationSystem struct {
	entities []animationSystemEntity
}

func (m *AnimationSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Animationable)
	m.Add(*obj.GetBasicEntity(), obj.GetTextureComponent(), obj.GetAnimationComponent())
}

func (m *AnimationSystem) Add(basic ecs.BasicEntity, texture *TextureComponent, animation *AnimationComponent) {
	m.entities = append(m.entities, animationSystemEntity{basic, texture, animation})
}

func (m *AnimationSystem) Remove(basic ecs.BasicEntity) {
	var del = -1
	for index, entity := range m.entities {
		if entity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		m.entities = append(m.entities[:del], m.entities[del+1:]...)
	}
}

func (m *AnimationSystem) Update(dt float32) {
	for _, entity := range m.entities {
		offset := 0
		if entity.Playing {
			offset = entity.PlayingOffset
		}

		entity.accumulator += 1 // Add a render frame

		framesPer := entity.FramesPerAnimationFrame

		if !entity.Playing {
			framesPer = framesPer * 6
		}

		if entity.accumulator >= framesPer {
			entity.accumulator = 0
			entity.Frame += 1
			if entity.Frame >= entity.MaxFrames {
				entity.Frame = 0
			}
		}

		entity.Src = rl.NewRectangle(
			float32(entity.Frame+offset)*entity.Src.Width,
			float32(entity.State)*entity.Src.Height,
			entity.Src.Width,
			entity.Src.Height,
		)
	}
}

func AddAnimationSystemToWorld(w *ecs.World) {
	var animationable *Animationable
	w.AddSystemInterface(&AnimationSystem{}, animationable, nil)
}
