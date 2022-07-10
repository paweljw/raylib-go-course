package ecs

type InputComponent struct {
	Speed float32
}

func (r *InputComponent) GetInputComponent() *InputComponent {
	return r
}

type InputFace interface {
	GetInputComponent() *InputComponent
}
