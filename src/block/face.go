package block

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockFace struct {
	direction FacingDirection
}

type Normal rl.Vector3

func (n Normal) GetUV() (rl.Vector3, rl.Vector3) {
	switch n {
	case Normal(rl.NewVector3(1, 0, 0)): // Right
		return rl.NewVector3(0, 1, 0), rl.NewVector3(0, 0, 1)
	case Normal(rl.NewVector3(-1, 0, 0)): // Left
		return rl.NewVector3(0, 1, 0), rl.NewVector3(0, 0, -1)
	case Normal(rl.NewVector3(0, 1, 0)): // Up
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 0, 1)
	case Normal(rl.NewVector3(0, -1, 0)): // Down
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 0, -1)
	case Normal(rl.NewVector3(0, 0, 1)): // Front
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 1, 0)
	case Normal(rl.NewVector3(0, 0, -1)): // Back
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, -1, 0)
	}

	panic(fmt.Sprintf("unknown normal vector: %v", n))
}

type FacingDirection string

func (fd FacingDirection) GetNormal() Normal {
	switch fd {
	case "up":
		return Normal(rl.NewVector3(0, 1, 0))
	case "down":
		return Normal(rl.NewVector3(0, -1, 0))
	case "front":
		return Normal(rl.NewVector3(0, 0, 1))
	case "back":
		return Normal(rl.NewVector3(0, 0, -1))
	case "right":
		return Normal(rl.NewVector3(1, 0, 0))
	case "left":
		return Normal(rl.NewVector3(-1, 0, 0))
	}

	panic("unknown facing direction: " + string(fd))
}
