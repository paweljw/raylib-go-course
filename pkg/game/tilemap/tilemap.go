package tilemap

import (
	"errors"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/solarlune/ldtkgo"
	"image/color"
	"log"
	"math"
	"path/filepath"
)

type World struct {
	world     *ldtkgo.Project
	assetPath string
}

func LoadWorld(path, filename string) (*World, error) {
	w := World{}
	w.assetPath = path
	project, err := ldtkgo.Open(filepath.Join(path, filename))
	if err != nil {
		return nil, err
	}
	w.world = project

	return &w, nil
}

func (w *World) LevelEntities(levelId string) ([]*ldtkgo.Entity, error) {
	level := w.world.LevelByIdentifier(levelId)
	if level == nil {
		return []*ldtkgo.Entity{}, errors.New("level not found: " + levelId)
	}

	for _, layer := range level.Layers {
		if layer.Type == ldtkgo.LayerTypeEntity {
			return layer.Entities, nil
		}
	}

	return []*ldtkgo.Entity{}, errors.New("entity layer not found")
}

func (w *World) LevelTiles(levelId string) (*rl.Texture2D, error) {
	level := w.world.LevelByIdentifier(levelId)
	if level == nil {
		return nil, errors.New("level not found: " + levelId)
	}

	target := rl.LoadRenderTexture(int32(level.Width), int32(level.Height))

	rl.BeginTextureMode(target)
	r, g, b, a := level.BGColor.RGBA()

	rl.ClearBackground(color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})

	for i := len(level.Layers) - 1; i >= 0; i-- {
		layer := level.Layers[i]
		if layer.Type != ldtkgo.LayerTypeTile {
			continue
		}
		sourceTexture := rl.LoadTexture(filepath.Join(w.assetPath, layer.Tileset.Path))
		size := layer.Tileset.GridSize

		for _, tile := range layer.Tiles {
			rl.DrawTexturePro(
				sourceTexture,
				rl.NewRectangle(
					float32(tile.Src[0]),
					float32(tile.Src[1]),
					float32(size), //or maybe -size
					float32(-size),
				),
				rl.NewRectangle(
					float32(size+tile.Position[0]),
					float32(math.Abs(float64(level.Height-tile.Position[1]))),
					float32(size), // TODO: support for  tile.FlipX, Y etc
					float32(size),
				),
				rl.NewVector2(float32(size), float32(size)),
				0, // or maybe 180...
				rl.White,
			)
		}
	}

	rl.EndTextureMode()

	texture := target.Texture
	return &texture, nil
}

func LoadLdtk() {
	level, err := ldtkgo.Open("./res/level.ldtk")
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("WorldLayout: %v", level.WorldLayout)
	log.Printf("WorldGridWidth: %v", level.WorldGridWidth)
	log.Printf("WorldGridHeight: %v", level.WorldGridHeight)
	log.Printf("BGColorString: %v", level.BGColorString)
	log.Printf("BGColor: %v", level.BGColor)
	log.Printf("JSONVersion: %v", level.JSONVersion)

	log.Println("--- LEVELS")

	for _, item := range level.Levels {
		log.Printf("+ Identifier: %v", item.Identifier)
		log.Printf("+ BGImage: %v", item.BGImage)
		log.Printf("+ BGColor: %v", item.BGColor)
		log.Printf("+ BGColorString: %v", item.BGColorString)
		log.Printf("+ WorldY: %v", item.WorldY)
		log.Printf("+ WorldX: %v", item.WorldX)
		log.Printf("+ Width: %v", item.Width)
		log.Printf("+ Height: %v", item.Height)

		log.Printf("---- LAYERS:")

		for _, layer := range item.Layers {
			log.Printf("Layer ID: %v", layer.Identifier)
			log.Printf("GridSize: %v", layer.GridSize)
			log.Printf("CellWidth: %v", layer.CellWidth)
			log.Printf("CellHeight: %v", layer.CellHeight)
			log.Printf("OffsetX: %v", layer.OffsetX)
			log.Printf("OffsetY: %v", layer.OffsetY)
			log.Printf("Type: %v", layer.Type)
			log.Printf("Tileset: %v", layer.Tileset)
		}
		//
		//log.Printf("---- PROPERTIES:")
		//
		//for _, property := range item.Properties {
		//
		//}
	}

	//WorldLayout     string
	//WorldGridWidth  int
	//WorldGridHeight int
	//BGColorString   string      `json:"defaultLevelBgColor"`
	//BGColor         color.Color `json:"-"`
	//JSONVersion     string
	//Levels          []*Level
	//Tilesets        []*Tileset
	//IntGridNames    []string
}
