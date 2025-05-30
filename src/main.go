package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	stoneTexture            rl.Texture2D
	dirtTexture             rl.Texture2D
	grassTopTexture         rl.Texture2D
	grassSideOverlayTexture rl.Texture2D
)

func main() {
	screenWidth, screenHeight := int32(1200), int32(675)
	rl.InitWindow(screenWidth, screenHeight, "MC Clone")

	stoneTexture = rl.LoadTexture("assets/blocks/stone.png")
	dirtTexture = rl.LoadTexture("assets/blocks/dirt.png")
	grassTopTexture = rl.LoadTexture("assets/blocks/grass_top.png")
	grassSideOverlayTexture = rl.LoadTexture("assets/blocks/grass_side_overlay.png")

	camera := rl.NewCamera3D(
		rl.NewVector3(32, 70, 32), // position
		rl.NewVector3(0, 0, 0),    // target
		rl.NewVector3(0, 1, 0),    // up
		60,                        // fov
		rl.CameraPerspective,
	)

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	// chunks := buildChunks()
	block1 := NewBlock(BlockGrass)
	position1 := rl.NewVector3(28, 65, 28)

	block2 := NewBlock(BlockStone)
	position2 := rl.NewVector3(50, 65, 28)

	for !rl.WindowShouldClose() {
		UpdateCamera(&camera)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		block1.Draw(position1)
		block2.Draw(position2)

		// for _, chunk := range chunks {
		// 	chunk.Draw()
		// }

		rl.EndMode3D()

		drawFPS(screenWidth)

		rl.EndDrawing()
	}

	rl.UnloadTexture(stoneTexture)
	rl.UnloadTexture(dirtTexture)
	rl.UnloadTexture(grassTopTexture)
	rl.UnloadTexture(grassSideOverlayTexture)

	rl.CloseWindow()
}

func buildChunks() []*Chunk {
	width, depth := 1, 1

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
