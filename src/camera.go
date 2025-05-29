package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	CAMERA_SPEED      = 0.18
	MOUSE_SENSITIVITY = 0.006
)

func UpdateCamera(camera *rl.Camera) {
	var mousePositionDelta = rl.GetMouseDelta()

	var moveInWorldPlane uint8
	moveInWorldPlane = 1

	var rotateAroundTarget uint8

	var lockView uint8
	lockView = 1

	var rotateUp uint8

	rl.CameraYaw(camera, -mousePositionDelta.X*MOUSE_SENSITIVITY, rotateAroundTarget)
	rl.CameraPitch(camera, -mousePositionDelta.Y*MOUSE_SENSITIVITY, lockView, rotateAroundTarget, rotateUp)

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
