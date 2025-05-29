package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth, screenHeight := int32(800), int32(600)
	rl.InitWindow(screenWidth, screenHeight, "Minecraft Clone")

	camera := rl.NewCamera3D(
		rl.NewVector3(32, 24, 32), // position
		rl.NewVector3(0, 0, 0),    // target
		rl.NewVector3(0, 1, 0),    // up
		60,                        // fov
		rl.CameraPerspective,
	)

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	chunk := NewChunk(32, 32, 32)

	for !rl.WindowShouldClose() {
		UpdateCamera(&camera)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		chunk.Draw()

		rl.EndMode3D()

		rl.DrawText("Move with WASD", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
