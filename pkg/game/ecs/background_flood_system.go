package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
)

type BackgroundFloodSystem struct {
	Color rl.Color
}

func (m *BackgroundFloodSystem) Update(dt float32) {
	rl.ClearBackground(m.Color)
}

func (m *BackgroundFloodSystem) Remove(e ecs.BasicEntity) {
}

func AddBackgroundFloodSystemToWorld(w *ecs.World) {
	w.AddSystem(&BackgroundFloodSystem{Color: common.BackgroundFillColor})
}
