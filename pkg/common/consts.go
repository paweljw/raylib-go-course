package common

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080

	WorldWidth  = 16
	WorldHeight = 9

	TargetFps   = 60
	PlayerSpeed = 0.05
)

var (
	BackgroundFillColor = rl.NewColor(147, 211, 196, 255)
)

func WorldWidthToScreen(w float32) float32 {
	return w / WorldWidth * ScreenWidth
}

func WorldHeightToScreen(w float32) float32 {
	return w / WorldHeight * ScreenHeight
}

func ScreenWidthToWorld(w float32) float32 {
	return w / ScreenWidth * WorldWidth
}

func ScreenHeightToWorld(w float32) float32 {
	return w / ScreenHeight * WorldHeight
}

func WorldRectToScreen(rect rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(
		WorldWidthToScreen(rect.X),
		WorldHeightToScreen(rect.Y),
		WorldWidthToScreen(rect.Width),
		WorldHeightToScreen(rect.Height),
	)
}

func ScreenRectToWorld(rect rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(
		ScreenWidthToWorld(rect.X),
		ScreenHeightToWorld(rect.Y),
		ScreenWidthToWorld(rect.Width),
		ScreenHeightToWorld(rect.Height),
	)
}

func Upscale(rect rl.Rectangle, scale float32) rl.Rectangle {
	return rl.NewRectangle(
		rect.X*scale,
		rect.Y*scale,
		rect.Width*scale,
		rect.Height*scale,
	)
}
