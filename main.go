package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	screenWidth, screenHeight := int32(800), int32(600)
	rl.InitWindow(screenWidth, screenHeight, "Minecraft Clone")

	rl.SetTargetFPS(60)

	camera := rl.NewCamera3D(
		rl.NewVector3(16, 16, 16), // position
		rl.NewVector3(0, 0, 0),    // target
		rl.NewVector3(0, 1, 0),    // up
		45,                        // fov
		rl.CameraPerspective,
	)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		blockPos := rl.NewVector3(0, 0, 0)
		blockSize := rl.NewVector3(1, 1, 1)

		rl.DrawCubeV(blockPos, blockSize, rl.Brown)
		rl.DrawCubeWiresV(blockPos, blockSize, rl.Black)

		rl.EndMode3D()

		rl.DrawText("Move with WASD", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
