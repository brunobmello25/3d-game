package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	screenWidth, screenHeight := int32(1200), int32(675)
	rl.InitWindow(screenWidth, screenHeight, "Minecraft Clone")

	camera := rl.NewCamera3D(
		rl.NewVector3(32, 70, 32), // position
		rl.NewVector3(0, 0, 0),    // target
		rl.NewVector3(0, 1, 0),    // up
		60,                        // fov
		rl.CameraPerspective,
	)

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	chunks := buildChunks()

	for !rl.WindowShouldClose() {
		UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		for _, chunk := range chunks {
			chunk.Draw()
		}

		rl.EndMode3D()

		drawFPS(screenWidth)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func buildChunks() []*Chunk {
	width, depth := 5, 5

	chunks := []*Chunk{}

	for x := range width {
		for z := range depth {
			chunk := NewChunk(x, 0, z)
			chunks = append(chunks, chunk)
		}
	}

	return chunks
}

func drawFPS(screenW int32) {
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
