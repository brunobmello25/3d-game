package block

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/texture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockFace struct {
	Direction   FacingDirection
	textureName string
}

func NewFace(direction FacingDirection, textureName string) BlockFace {
	return BlockFace{
		Direction:   direction,
		textureName: textureName,
	}
}

func (f BlockFace) GetVertexCoords(blockCenter rl.Vector3) []float32 {
	faceNormal := f.Direction.GetNormal()
	halfNormal := rl.Vector3Scale(faceNormal.ToVector3(), 0.5) // Half of the normal vector

	faceCenter := rl.Vector3Add(blockCenter, rl.Vector3(halfNormal))

	u, v := faceNormal.getUV()
	hu := rl.Vector3Scale(u, 0.5) // Half of the u vector
	hv := rl.Vector3Scale(v, 0.5) // Half of the v vector

	bl := rl.Vector3Subtract(rl.Vector3Subtract(faceCenter, hu), hv) // C - U - V
	br := rl.Vector3Subtract(rl.Vector3Add(faceCenter, hu), hv)      // C + U - V
	tr := rl.Vector3Add(rl.Vector3Add(faceCenter, hu), hv)           // C + U + V
	tp := rl.Vector3Add(rl.Vector3Subtract(faceCenter, hu), hv)      // C - U + V

	return []float32{
		bl.X, bl.Y, bl.Z, // bottom-left
		br.X, br.Y, br.Z, // bottom-right
		tr.X, tr.Y, tr.Z, // top-right
		tp.X, tp.Y, tp.Z, // top-left
	}
}

func (f BlockFace) GetVertexNormals(blockCenter rl.Vector3) []float32 {
	faceNormal := f.Direction.GetNormal()

	return []float32{
		faceNormal.X, faceNormal.Y, faceNormal.Z, // bottom-left
		faceNormal.X, faceNormal.Y, faceNormal.Z, // bottom-right
		faceNormal.X, faceNormal.Y, faceNormal.Z, // top-right
		faceNormal.X, faceNormal.Y, faceNormal.Z, // top-left
	}
}

func (f BlockFace) GetTextureCoords() []float32 {
	// Get the UV coordinates from the texture atlas
	uv := texture.GetTextureUV(f.textureName)

	// Texture coordinates are defined in a counter clockwise
	// manner starting from the bottom-left corner
	return []float32{
		uv.X, uv.Y + uv.Height, // bottom-left
		uv.X + uv.Width, uv.Y + uv.Height, // bottom-right
		uv.X + uv.Width, uv.Y, // top-right
		uv.X, uv.Y, // top-left
	}
}

type Normal rl.Vector3

func (n Normal) ToVector3() rl.Vector3 {
	return rl.Vector3(n)
}

func (n Normal) getUV() (rl.Vector3, rl.Vector3) {
	switch n {
	case Normal(rl.NewVector3(1, 0, 0)): // Right
		return rl.NewVector3(0, 0, -1), rl.NewVector3(0, 1, 0)
	case Normal(rl.NewVector3(-1, 0, 0)): // Left
		return rl.NewVector3(0, 0, 1), rl.NewVector3(0, 1, 0)
	case Normal(rl.NewVector3(0, 1, 0)): // Up
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 0, -1)
	case Normal(rl.NewVector3(0, -1, 0)): // Down
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 0, 1)
	case Normal(rl.NewVector3(0, 0, 1)): // Front
		return rl.NewVector3(1, 0, 0), rl.NewVector3(0, 1, 0)
	case Normal(rl.NewVector3(0, 0, -1)): // Back
		return rl.NewVector3(-1, 0, 0), rl.NewVector3(0, 1, 0)
	}

	panic(fmt.Sprintf("unknown normal vector: %v", n))
}

type FacingDirection string

const (
	FacingDirectionUp    FacingDirection = "up"
	FacingDirectionDown  FacingDirection = "down"
	FacingDirectionFront FacingDirection = "front"
	FacingDirectionBack  FacingDirection = "back"
	FacingDirectionRight FacingDirection = "right"
	FacingDirectionLeft  FacingDirection = "left"
)

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
