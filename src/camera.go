package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	CAMERA_SPEED = 0.09
)

func UpdateCamera(camera *rl.Camera) {
	var mousePositionDelta = rl.GetMouseDelta()

	var moveInWorldPlane uint8
	moveInWorldPlane = 1

	var rotateAroundTarget uint8

	var lockView uint8
	lockView = 1

	var rotateUp uint8

	// Camera rotation
	if rl.IsKeyDown(rl.KeyDown) {
		rl.CameraPitch(camera, -0.03, lockView, rotateAroundTarget, rotateUp)
	}
	if rl.IsKeyDown(rl.KeyUp) {
		rl.CameraPitch(camera, 0.03, lockView, rotateAroundTarget, rotateUp)
	}
	if rl.IsKeyDown(rl.KeyRight) {
		rl.CameraYaw(camera, -0.03, rotateAroundTarget)
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		rl.CameraYaw(camera, 0.03, rotateAroundTarget)
	}
	if rl.IsKeyDown(rl.KeyQ) {
		rl.CameraRoll(camera, -0.03)
	}
	if rl.IsKeyDown(rl.KeyE) {
		rl.CameraRoll(camera, 0.03)
	}

	rl.CameraYaw(camera, -mousePositionDelta.X*0.003, rotateAroundTarget)
	rl.CameraPitch(camera, -mousePositionDelta.Y*0.003, lockView, rotateAroundTarget, rotateUp)

	// Keyboard support
	if rl.IsKeyDown(rl.KeyW) {
		rl.CameraMoveForward(camera, CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyA) {
		rl.CameraMoveRight(camera, -CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyS) {
		rl.CameraMoveForward(camera, -CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyD) {
		rl.CameraMoveRight(camera, CAMERA_SPEED, moveInWorldPlane)
	}
}
