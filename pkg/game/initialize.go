package game

import (
	engoecs "github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
)

var (
	world *engoecs.World
)

func Initialize() {
	world = ecs.NewWorld()

	// This needs to come before any textures are loaded
	rl.InitWindow(common.ScreenWidth, common.ScreenHeight, "raylib-go course")
	rl.InitAudioDevice()
	rl.SetExitKey(rl.KeyQ)
	rl.SetTargetFPS(common.TargetFps)

	playerEntity := NewPlayerEntity()
	grassEntity := NewGrassEntity()
	musicEntity := NewMusicEntity()

	world.AddEntity(grassEntity)
	world.AddEntity(playerEntity)
	world.AddEntity(musicEntity)
}
