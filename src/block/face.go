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

func (f BlockFace) GetVertexColors() []uint8 {
	// Get base lighting color for this face direction
	baseColor := getLightingColor(f.Direction)

	// Create gradient within the face for better visual distinction
	// The gradient varies based on the face direction to create consistent lighting
	var colors [4]rl.Color

	switch f.Direction {
	case FacingDirectionUp:
		// Top face: gradient from front to back (darker in back)
		colors[0] = modulateColor(baseColor, 0.9)  // bottom-left (front): slightly darker
		colors[1] = modulateColor(baseColor, 0.95) // bottom-right (front): slightly darker
		colors[2] = modulateColor(baseColor, 1.0)  // top-right (back): base brightness
		colors[3] = modulateColor(baseColor, 0.85) // top-left (back): darker

	case FacingDirectionDown:
		// Bottom face: gradient from back to front (darker in front)
		colors[0] = modulateColor(baseColor, 1.0)  // bottom-left: base brightness
		colors[1] = modulateColor(baseColor, 0.9)  // bottom-right: darker
		colors[2] = modulateColor(baseColor, 0.8)  // top-right: darkest
		colors[3] = modulateColor(baseColor, 0.95) // top-left: slightly darker

	case FacingDirectionFront, FacingDirectionBack:
		// Front/Back faces: gradient from top to bottom (darker at bottom)
		colors[0] = modulateColor(baseColor, 0.8)  // bottom-left: darker
		colors[1] = modulateColor(baseColor, 0.85) // bottom-right: darker
		colors[2] = modulateColor(baseColor, 1.0)  // top-right: brightest
		colors[3] = modulateColor(baseColor, 0.95) // top-left: bright

	case FacingDirectionLeft, FacingDirectionRight:
		// Side faces: gradient from top to bottom and slight left-right variation
		colors[0] = modulateColor(baseColor, 0.75) // bottom-left: darkest
		colors[1] = modulateColor(baseColor, 0.8)  // bottom-right: dark
		colors[2] = modulateColor(baseColor, 0.95) // top-right: bright
		colors[3] = modulateColor(baseColor, 0.9)  // top-left: medium-bright
	}

	// Return colors for all 4 vertices (RGBA format)
	return []uint8{
		colors[0].R, colors[0].G, colors[0].B, colors[0].A, // bottom-left
		colors[1].R, colors[1].G, colors[1].B, colors[1].A, // bottom-right
		colors[2].R, colors[2].G, colors[2].B, colors[2].A, // top-right
		colors[3].R, colors[3].G, colors[3].B, colors[3].A, // top-left
	}
}

func modulateColor(color rl.Color, factor float32) rl.Color {
	// Clamp factor to prevent overflow
	if factor > 1.0 {
		factor = 1.0
	}
	if factor < 0.0 {
		factor = 0.0
	}

	return rl.Color{
		R: uint8(float32(color.R) * factor),
		G: uint8(float32(color.G) * factor),
		B: uint8(float32(color.B) * factor),
		A: color.A,
	}
}

func getLightingColor(direction FacingDirection) rl.Color {
	// Different lighting levels for each face direction
	switch direction {
	case FacingDirectionUp:
		return rl.Color{R: 255, G: 255, B: 255, A: 255} // Full brightness for top faces
	case FacingDirectionDown:
		return rl.Color{R: 100, G: 100, B: 100, A: 255} // Darkest for bottom faces
	case FacingDirectionFront, FacingDirectionBack:
		return rl.Color{R: 200, G: 200, B: 200, A: 255} // Medium-bright for front/back faces
	case FacingDirectionLeft, FacingDirectionRight:
		return rl.Color{R: 150, G: 150, B: 150, A: 255} // Medium-dark for left/right faces
	default:
		return rl.Color{R: 255, G: 255, B: 255, A: 255} // Default to full brightness
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
