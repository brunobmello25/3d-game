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
	camera   rl.Camera3D
	position rl.Vector3
	rotation rl.Vector2
}

func NewPlayer() *Player {
	return &Player{
		camera: rl.NewCamera3D(
			rl.NewVector3(0, 0, 0),
			rl.NewVector3(1, 0, 0),
			rl.NewVector3(0, 1, 0),
			60,
			rl.CameraPerspective,
		),
		position: rl.NewVector3(0, 0, 0),
		rotation: rl.NewVector2(0, 0),
	}
}

func (p *Player) Init() {
	rl.DisableCursor()
}

func (p *Player) Update() {
	// Handle mouse movement
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
		p.position = rl.Vector3Add(p.position, rl.Vector3Scale(forward, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.position = rl.Vector3Subtract(p.position, rl.Vector3Scale(forward, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.position = rl.Vector3Subtract(p.position, rl.Vector3Scale(right, PLAYER_SPEED))
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.position = rl.Vector3Add(p.position, rl.Vector3Scale(right, PLAYER_SPEED))
	}

	// Update camera position and target
	p.camera.Position = p.position
	p.camera.Target = rl.Vector3Add(p.position, forward)
}

func (p *Player) GetCamera() rl.Camera3D {
	return p.camera
}

func (p *Player) GetPosition() rl.Vector3 {
	return p.position
}
