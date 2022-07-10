package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/paweljw/raylib-go-course/pkg/common"
	"log"
)

type renderSystemEntity struct {
	ecs.BasicEntity
	*TextureComponent
	*RenderComponent
}

type Renderable interface {
	ecs.BasicFace
	TextureFace
	RenderFace
}

type RenderSystem struct {
	camera   *rl.Camera2D
	entities []renderSystemEntity
	world    *ecs.World
}

func (m *RenderSystem) New(w *ecs.World) {
	m.world = w
	m.camera = GetWorldCamera(w)
}

func (m *RenderSystem) Quit() {
	for _, entity := range m.entities {
		m.Remove(entity.BasicEntity)
	}
}

func (m *RenderSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Renderable)
	m.Add(
		*obj.GetBasicEntity(),
		obj.GetTextureComponent(),
		obj.GetRenderComponent(),
	)
}

func (m *RenderSystem) Add(basic ecs.BasicEntity, texture *TextureComponent, render *RenderComponent) {
	if texture.textureLoaded == false {
		texture.texture = rl.LoadTexture(texture.TexturePath)
		texture.textureLoaded = true
		log.Printf("Loading texture: %s", texture.TexturePath)
	}

	m.entities = append(m.entities, renderSystemEntity{basic, texture, render})
}
func (m *RenderSystem) Remove(basic ecs.BasicEntity) {
	var del = -1
	for index, entity := range m.entities {
		if entity.ID() == basic.ID() {
			del = index
			break
		}
	}
	if del >= 0 {
		if m.entities[del].textureLoaded == true {
			rl.UnloadTexture(m.entities[del].texture)
			m.entities[del].textureLoaded = false
			log.Printf("Unloading texture: %s", m.entities[del].TexturePath)
		}

		m.entities = append(m.entities[:del], m.entities[del+1:]...)
	}
}

func (m *RenderSystem) Update(dt float32) {
	if m.camera == nil {
		m.camera = GetWorldCamera(m.world)
		if m.camera == nil {
			return
		}
	}

	rl.BeginDrawing()
	rl.BeginMode2D(*m.camera)

	for _, entity := range m.entities {
		screenRect := common.WorldRectToScreen(entity.Dest)

		rl.DrawTexturePro(
			entity.texture,
			entity.Src, screenRect,
			rl.NewVector2(screenRect.Width, screenRect.Height),
			entity.Rotation,
			entity.Tint,
		)
	}

	rl.EndMode2D()
	rl.DrawFPS(0, 0)
	rl.EndDrawing()
}

func AddRenderSystemToWorld(w *ecs.World) {
	var renderable *Renderable
	w.AddSystemInterface(&RenderSystem{}, renderable, nil)
}
