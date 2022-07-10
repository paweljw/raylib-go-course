package ecs

import rl "github.com/gen2brain/raylib-go/raylib"

type CameraComponent struct {
	Camera rl.Camera2D
}

func (r *CameraComponent) GetCameraComponent() *CameraComponent {
	return r
}

type CameraFace interface {
	GetCameraComponent() *CameraComponent
}
