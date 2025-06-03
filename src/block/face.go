package block

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockFace struct {
	direction FacingDirection
	texture   rl.Texture2D
}

func NewFace(direction FacingDirection) BlockFace {
	return BlockFace{
		direction: direction,
	}
}

func (f BlockFace) getVertexData(blockCenter rl.Vector3) [6]rl.Vector3 {
	faceNormal := f.direction.getNormal()

	faceCenter := rl.Vector3Add(blockCenter, rl.Vector3(faceNormal))

	u, v := faceNormal.getUV()

	bottomLeft := rl.Vector3Subtract(rl.Vector3Subtract(faceCenter, u), v) // C - U - V
	bottomRight := rl.Vector3Subtract(rl.Vector3Add(faceCenter, u), v)     // C + U - V
	topRight := rl.Vector3Add(rl.Vector3Add(faceCenter, u), v)             // C + U + V
	topLeft := rl.Vector3Add(rl.Vector3Subtract(faceCenter, u), v)         // C - U + V

	return [6]rl.Vector3{
		// triangle 1: BL, BR, TR
		bottomLeft,
		bottomRight,
		topRight,

		// triangle 2: BL, TR, TL
		bottomLeft,
		topRight,
		topLeft,
	}
}

type Normal rl.Vector3

func (n Normal) getUV() (rl.Vector3, rl.Vector3) {
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

func (fd FacingDirection) getNormal() Normal {
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
