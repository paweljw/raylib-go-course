package game

import (
	engoecs "github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
	"github.com/paweljw/raylib-go-course/pkg/game/ecs"
	"log"
)

func NewPlayerEntity() *ecs.PlayerEntity {
	playerEntity := ecs.PlayerEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src: rl.NewRectangle(0, 0, 48, 48),

			TexturePath: "res/sproutlands/Characters/character_spritesheet.png",
		},
		RenderComponent: &ecs.RenderComponent{
			Dest:     rl.NewRectangle(8, 4.5, 2, 2),
			Rotation: 0,
			Tint:     rl.White,
		},
		InputComponent:  &ecs.InputComponent{Speed: common.PlayerSpeed},
		CameraComponent: &ecs.CameraComponent{Camera: rl.NewCamera2D(rl.NewVector2(common.ScreenWidth/2, common.ScreenHeight/2), rl.NewVector2(0, 0), 0, 1.0)},
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
	sourceTexture := rl.LoadTexture("res/sproutlands/Tilesets/Grass.png")

	tilemap := []int{
		4, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
		3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
		4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
		5, 5, 5, 5, 5, 5, 5, 5, 5, 5,
		10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	}
	tilemapWidth := 10
	tilemapHeight := 10
	tileWidth := 16
	tileHeight := 16
	sourceWidth := 160

	target := rl.LoadRenderTexture(int32(tilemapWidth*tileWidth), int32(tilemapHeight*tileHeight))

	rl.BeginTextureMode(target)
	rl.ClearBackground(rl.Black)
	for i := 0; i < tilemapWidth; i++ {
		for j := 0; j < tilemapHeight; j++ {
			// tilemap[i+j*tilemapWidth]
			sourceX := tilemap[i+j*tilemapWidth] * tileWidth % sourceWidth
			sourceY := (tilemap[i+j*tilemapWidth] * tileWidth / sourceWidth) * tileHeight
			log.Printf("i:%d, j:%d, sourceX: %d, sourceY: %d", i, j, sourceX, sourceY)

			rl.DrawTexturePro(
				sourceTexture,
				rl.NewRectangle(
					float32(sourceX),
					float32(sourceY),
					float32(-tileWidth),
					float32(tileHeight),
				),
				rl.NewRectangle(
					float32(i*tileWidth),
					float32((tilemapHeight*tileHeight)-(j+1)*tileHeight),
					float32(tileWidth),
					float32(tileHeight),
				),
				rl.NewVector2(float32(tileWidth), float32(tileHeight)),
				180,
				rl.White,
			)
		}
	}
	rl.EndTextureMode()

	sourceRect := rl.NewRectangle(0, 0, float32(tileWidth*tilemapWidth), float32(tileHeight*tilemapHeight))
	destRect := common.Upscale(sourceRect, 5)
	destRect.X = common.ScreenWidth / 2
	destRect.Y = common.ScreenHeight / 2
	log.Println(destRect)

	grassEntity := ecs.BackgroundEntity{
		BasicEntity: engoecs.BasicEntity{},
		TextureComponent: &ecs.TextureComponent{
			Src:           sourceRect,
			TexturePath:   "",
			Texture:       target.Texture,
			TextureLoaded: true,
		},
		RenderComponent: &ecs.RenderComponent{
			Dest:     common.ScreenRectToWorld(destRect),
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
