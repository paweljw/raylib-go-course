package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const targetFrameTime = 1 / 60.0

type Inputable interface {
	ecs.BasicFace
	InputFace
	RenderFace
	AnimationFace
}

type inputSystemEntity struct {
	ecs.BasicEntity
	*InputComponent
	*RenderComponent
	*AnimationComponent
}

type InputSystem struct {
	entities []inputSystemEntity
}

func (m *InputSystem) Quit() {
	for _, entity := range m.entities {
		m.Remove(entity.BasicEntity)
	}
}

func (m *InputSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Inputable)
	m.Add(*obj.GetBasicEntity(), obj.GetInputComponent(), obj.GetRenderComponent(), obj.GetAnimationComponent())
}

func (m *InputSystem) Add(basic ecs.BasicEntity, input *InputComponent, render *RenderComponent, animation *AnimationComponent) {
	m.entities = append(m.entities, inputSystemEntity{basic, input, render, animation})
}
func (m *InputSystem) Remove(basic ecs.BasicEntity) {
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

func (m *InputSystem) Update(dt float32) {
	for _, entity := range m.entities {
		realSpeed := entity.Speed * (dt / targetFrameTime)
		moving := false

		if rl.IsKeyDown(rl.KeyD) {
			entity.Dest.X += realSpeed
			moving = true
			entity.State = 3
		}

		if rl.IsKeyDown(rl.KeyA) {
			entity.Dest.X -= realSpeed
			moving = true
			entity.State = 2
		}

		if rl.IsKeyDown(rl.KeyW) {
			entity.Dest.Y -= realSpeed
			moving = true
			entity.State = 1
		}

		if rl.IsKeyDown(rl.KeyS) {
			entity.Dest.Y += realSpeed
			moving = true
			entity.State = 0
		}

		if moving && !entity.Playing {
			entity.accumulator = entity.FramesPerAnimationFrame + 1
		}

		if !moving && entity.Playing {
			entity.accumulator = 0
			entity.Frame = 0
		}

		entity.Playing = moving

		if rl.IsKeyDown(rl.KeyUp) {
			entity.Dest.Height += realSpeed
			entity.Dest.Width += realSpeed
		}

		if rl.IsKeyDown(rl.KeyDown) {
			entity.Dest.Height -= realSpeed
			entity.Dest.Width -= realSpeed
		}
	}
}

func AddInputSystemToWorld(w *ecs.World) {
	var inputable *Inputable
	w.AddSystemInterface(&InputSystem{}, inputable, nil)
}
