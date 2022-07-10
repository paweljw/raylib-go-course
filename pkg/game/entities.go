package game

import (
	engoecs "github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
)

func NewPlayerEntity() *ecs.PlayerEntity {
	playerEntity := ecs.PlayerEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src: rl.NewRectangle(0, 0, 48, 48),

			TexturePath: "res/sproutlands/Characters/character_spritesheet.png",
		},
		RenderComponent: &ecs.RenderComponent{
			Dest:     rl.NewRectangle(3, 3, 2, 2),
			Rotation: 0,
			Tint:     rl.White,
		},
		InputComponent: &ecs.InputComponent{Speed: common.PlayerSpeed},
	}

	return &playerEntity
}

func NewGrassEntity() *ecs.BackgroundEntity {
	grassEntity := ecs.BackgroundEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src:         rl.NewRectangle(0, 0, 160, 128),
			TexturePath: "res/sproutlands/Tilesets/Grass.png",
		},
		RenderComponent: &ecs.RenderComponent{
			Dest:     rl.NewRectangle(16, 9, 16, 9),
			Rotation: 0,
			Tint:     rl.White,
		},
	}

	return &grassEntity
}

func NewMusicEntity() *ecs.MusicEntity {
	musicEntity := ecs.MusicEntity{
		BasicEntity:    engoecs.BasicEntity{},
		MusicComponent: &ecs.MusicComponent{MusicPath: "res/averys_farm.mp3"},
	}
	return &musicEntity
}
