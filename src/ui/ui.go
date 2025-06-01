package ui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (r *UI) DrawFPS(screenW int32) {
	fps := rl.GetFPS()
	text := fmt.Sprintf("%d FPS", fps)
	fontSize := int32(20)

	// Measure the width of the text so we can right-align it
	textWidth := rl.MeasureText(text, fontSize)

	// draw it with a little padding from the edges:
	x := screenW - textWidth - 10
	y := int32(10)

	rl.DrawText(text, x, y, fontSize, rl.Black)
}

func (r *UI) DrawPosition(position rl.Vector3) {
	text := fmt.Sprintf("Position: %.1f, %.1f, %.1f", position.X, position.Y, position.Z)
	fontSize := int32(20)

	// Draw position in the top-left corner with padding
	x := int32(10)
	y := int32(10)

	rl.DrawText(text, x, y, fontSize, rl.Black)
}
