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

func WorldRectToScreen(rect rl.Rectangle) rl.Rectangle {
	return rl.NewRectangle(
		WorldWidthToScreen(rect.X),
		WorldHeightToScreen(rect.Y),
		WorldWidthToScreen(rect.Width),
		WorldHeightToScreen(rect.Height),
	)
}
