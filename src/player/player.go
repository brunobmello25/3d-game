package player

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	PLAYER_SPEED      = 0.1
	MOUSE_SENSITIVITY = 0.1
)

type Player struct {
	Camera   rl.Camera3D
	Position rl.Vector3

	rotation rl.Vector2
}

func NewPlayer() *Player {
	return &Player{
		Camera: rl.NewCamera3D(
			rl.NewVector3(0, 0, 0), // pos
			rl.NewVector3(1, 0, 0), // target
			rl.NewVector3(0, 1, 0), // up
			60,
			rl.CameraPerspective,
		),
		Position: rl.NewVector3(0, 0, 0),
		rotation: rl.NewVector2(0, 0),
	}
}

func (p *Player) Update() {
	mouseDelta := rl.GetMouseDelta()

	p.rotation.X -= float32(mouseDelta.X) * MOUSE_SENSITIVITY
	p.rotation.Y -= float32(mouseDelta.Y) * MOUSE_SENSITIVITY

	// Clamp vertical rotation
	if p.rotation.Y < -89 {
		p.rotation.Y = -89
	} else if p.rotation.Y > 89 {
		p.rotation.Y = 89
	}

	// Calculate forward and right vectors
	forward := rl.Vector3{
		X: float32(math.Cos(float64(rl.Deg2rad * p.rotation.X))),
		Y: 0,
		Z: float32(math.Sin(float64(rl.Deg2rad * p.rotation.X))),
	}
	right := rl.Vector3{
		X: float32(math.Cos(float64(rl.Deg2rad * (p.rotation.X + 90)))),
		Y: 0,
		Z: float32(math.Sin(float64(rl.Deg2rad * (p.rotation.X + 90)))),
	}

	// Handle keyboard input
	if rl.IsKeyDown(rl.KeyW) {
		p.Position = rl.Vector3Add(p.Position, rl.Vector3Scale(forward, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.Position = rl.Vector3Subtract(p.Position, rl.Vector3Scale(forward, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.Position = rl.Vector3Subtract(p.Position, rl.Vector3Scale(right, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.Position = rl.Vector3Add(p.Position, rl.Vector3Scale(right, PLAYER_SPEED))
	}

	// Update camera position and target
	p.Camera.Position = p.Position
	p.Camera.Target = rl.Vector3Add(p.Position, forward)
}
