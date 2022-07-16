package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type cameraSystemEntity struct {
	ecs.BasicEntity
	*CameraComponent
	*RenderComponent
}

type Camerable interface {
	ecs.BasicFace
	CameraFace
	RenderFace
}

type CameraSystemFace interface {
	GetCamera() *rl.Camera2D
}

type CameraSystem struct {
	entity *cameraSystemEntity
}

func (m *CameraSystem) GetCamera() *rl.Camera2D {
	if m.entity == nil {
		return nil
	}

	return &m.entity.Camera
}

func (m *CameraSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Camerable)
	m.Add(*obj.GetBasicEntity(), obj.GetCameraComponent(), obj.GetRenderComponent())
}

func (m *CameraSystem) Add(basic ecs.BasicEntity, camera *CameraComponent, render *RenderComponent) {
	m.entity = &cameraSystemEntity{basic, camera, render}
}

func (m *CameraSystem) Remove(basic ecs.BasicEntity) {
	m.entity = nil
}

func (m *CameraSystem) Update(dt float32) {
	if rl.IsKeyDown(rl.KeyEqual) {
		m.entity.Camera.Zoom = rl.Clamp(m.entity.Camera.Zoom+0.01, 0, 3)
	}

	if rl.IsKeyDown(rl.KeyMinus) {
		m.entity.Camera.Zoom = rl.Clamp(m.entity.Camera.Zoom-0.01, 0.3, 3)
	}

	if rl.IsKeyPressed(rl.KeyZero) {
		m.entity.Camera.Zoom = 1.0
	}

	m.entity.Camera.Target = rl.NewVector2(
		m.entity.Dest.X-m.entity.Dest.Width/2,
		m.entity.Dest.Y-m.entity.Dest.Height/2,
	)
}

func AddCameraSystemToWorld(w *ecs.World) {
	var camerable *Camerable
	w.AddSystemInterface(&CameraSystem{}, camerable, nil)
}
