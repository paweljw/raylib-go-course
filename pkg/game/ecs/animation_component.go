package ecs

type AnimationComponent struct {
	Frame                   int // Left to right
	State                   int // Top to bottom
	MaxFrames               int
	FramesPerAnimationFrame float32
	Playing                 bool
	accumulator             float32
	PlayingOffset           int
}

func (r *AnimationComponent) GetAnimationComponent() *AnimationComponent {
	return r
}

type AnimationFace interface {
	GetAnimationComponent() *AnimationComponent
}
