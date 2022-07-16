package game

import (
	engoecs "github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
	"github.com/paweljw/raylib-go-course/pkg/game/tilemap"
	"github.com/solarlune/ldtkgo"
	"log"
)

func NewPlayerEntity() *ecs.PlayerEntity {
	levelProject, err := tilemap.LoadWorld("./res/", "level.ldtk")
	if err != nil {
		log.Fatalln(err)
	}

	levelEntities, err := levelProject.LevelEntities("Level_0")
	if err != nil {
		log.Fatalln(err)
	}
	var playerStart *ldtkgo.Entity
	for _, entity := range levelEntities {
		if entity.Identifier == "Player_start" {
			playerStart = entity
			break
		}
	}
	log.Printf("Position: x %d, y %d", playerStart.Position[0], playerStart.Position[1])

	playerEntity := ecs.PlayerEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src: rl.NewRectangle(0, 0, 48, 48),

			TexturePath: "res/sproutlands/Characters/character_spritesheet.png",
		},
		RenderComponent: &ecs.RenderComponent{
			Dest: rl.NewRectangle(float32(playerStart.Position[0]-playerStart.Width), float32(playerStart.Position[1]-playerStart.Height),
				48, 48),
			Rotation: 0,
			Tint:     rl.White,
		},
		InputComponent:  &ecs.InputComponent{Speed: common.PlayerSpeed},
		CameraComponent: &ecs.CameraComponent{Camera: rl.NewCamera2D(rl.NewVector2(common.ScreenWidth/2, common.ScreenHeight/2), rl.NewVector2(0, 0), 0, 5.0)},
		AnimationComponent: &ecs.AnimationComponent{
			Frame:                   0,
			State:                   0,
			MaxFrames:               2,
			FramesPerAnimationFrame: 8,
			Playing:                 false,
			PlayingOffset:           2,
		},
	}

	return &playerEntity
}

func NewGrassEntity() *ecs.BackgroundEntity {
	levelProject, err := tilemap.LoadWorld("./res/", "level.ldtk")
	if err != nil {
		log.Fatalln(err)
	}

	texture, err := levelProject.LevelTiles("Level_0")
	if err != nil {
		log.Fatalln(err)
	}

	sourceRect := rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	grassEntity := ecs.BackgroundEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src:           sourceRect,
			TexturePath:   "",
			Texture:       *texture,
			TextureLoaded: true,
		},
		RenderComponent: &ecs.RenderComponent{
			Dest:     sourceRect,
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
