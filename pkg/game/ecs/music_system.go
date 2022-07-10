package ecs

import (
	"github.com/EngoEngine/ecs"
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

type musicSystemEntity struct {
	ecs.BasicEntity
	*MusicComponent
}

type Musicable interface {
	ecs.BasicFace
	MusicFace
}

type MusicSystem struct {
	MusicPlaying bool
	entity       *musicSystemEntity
}

func (m *MusicSystem) New(w *ecs.World) {
	m.MusicPlaying = false
}

func (m *MusicSystem) Quit() {
	m.Remove(m.entity.BasicEntity)
}

func (m *MusicSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Musicable)
	m.Add(*obj.GetBasicEntity(), obj.GetMusicComponent())
}

func (m *MusicSystem) Add(basic ecs.BasicEntity, music *MusicComponent) {
	if music.musicLoaded == false {
		music.music = rl.LoadMusicStream(music.MusicPath)
		music.musicLoaded = true
		log.Printf("Loading music: %s", music.MusicPath)
	}

	if m.entity != nil {
		if m.entity.Playing {
			rl.StopMusicStream(m.entity.music)
			m.Remove(m.entity.BasicEntity)
		}
	}

	m.entity = &musicSystemEntity{basic, music}
}

func (m *MusicSystem) Remove(basic ecs.BasicEntity) {
	if m.entity == nil || m.entity.BasicEntity.ID() != basic.ID() {
		return
	}

	if m.entity.musicLoaded == true {
		rl.UnloadMusicStream(m.entity.music)
		m.entity.musicLoaded = false
		log.Printf("Unloading music: %s", m.entity.MusicComponent)
	}

	m.entity = nil
}

func (m *MusicSystem) Update(dt float32) {
	// TODO: hmm, I'm not happy with this, but wiring with input might be hard
	//if rl.IsKeyPressed(rl.KeyQ) {
	//	m.MusicPlaying = !m.MusicPlaying
	//}

	rl.UpdateMusicStream(m.entity.music)

	if m.MusicPlaying {
		if m.entity.Playing {
			rl.ResumeMusicStream(m.entity.music)
		} else {
			log.Println("Starting new playback " + m.entity.MusicPath)
			rl.PlayMusicStream(m.entity.music)
			m.entity.Playing = true
		}
	} else {
		rl.PauseMusicStream(m.entity.music)
	}
}

func AddMusicSystemToWorld(w *ecs.World) {
	var musicable *Musicable
	w.AddSystemInterface(&MusicSystem{}, musicable, nil)
}
