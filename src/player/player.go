package player

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_SPEED      = 0.18
	MOUSE_SENSITIVITY = 0.006
)

type Player struct {
	Camera   rl.Camera3D
	Position rl.Vector3
}

func NewPlayer() *Player {
	position := rl.NewVector3(0, 0, 0)
	return &Player{
		Camera: rl.NewCamera3D(
			position,
			rl.NewVector3(0, 0, -1), // target
			rl.NewVector3(0, 1, 0),  // up
			60,
			rl.CameraPerspective,
		),
		Position: position,
	}
}

// TODO: should probably handle camera movement and rotation
// manually instead of using raylib's camera functions.
// Ideally, we want to update player position and rotation
// based on user input, and then just update camera values
// accordingly.
func (p *Player) Update() {
	var mousePositionDelta = rl.GetMouseDelta()

	var moveInWorldPlane uint8
	moveInWorldPlane = 1

	var rotateAroundTarget uint8

	var lockView uint8
	lockView = 1

	var rotateUp uint8

	rl.CameraYaw(&p.Camera, -mousePositionDelta.X*MOUSE_SENSITIVITY, rotateAroundTarget)
	rl.CameraPitch(&p.Camera, -mousePositionDelta.Y*MOUSE_SENSITIVITY, lockView, rotateAroundTarget, rotateUp)

	// Keyboard support
	if rl.IsKeyDown(rl.KeyW) {
		rl.CameraMoveForward(&p.Camera, CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyA) {
		rl.CameraMoveRight(&p.Camera, -CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyS) {
		rl.CameraMoveForward(&p.Camera, -CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeyD) {
		rl.CameraMoveRight(&p.Camera, CAMERA_SPEED, moveInWorldPlane)
	}
	if rl.IsKeyDown(rl.KeySpace) {
		rl.CameraMoveUp(&p.Camera, CAMERA_SPEED)
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		rl.CameraMoveUp(&p.Camera, -CAMERA_SPEED)
	}

	p.Position = p.Camera.Position
}
