package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (b *Block) GetColor() rl.Color {
	switch b.Type {
	case BlockAir:
		return rl.Blank
	case BlockDirt:
		return rl.Brown
	case BlockStone:
		return rl.NewColor(143, 143, 143, 255)
	case BlockGrass:
		return rl.Green
	}

	panic(fmt.Sprintf("Unknown block type: %d", b.Type))
}

func (b *Block) Draw(pos rl.Vector3) {
	if b.Type == BlockAir {
		return
	}

	// TODO: this probably should be in the block struct... maybe?
	// or maybe not, fuck it
	dimensions := rl.NewVector3(1, 1, 1)

	// Convert logical coordinates to world coordinates
	x := pos.X * dimensions.X
	y := pos.Y * dimensions.Y
	z := pos.Z * dimensions.Z

	rl.Begin(rl.Quads)

	// ===================================
	// =========== Top Face ==============
	// ===================================
	rl.SetTexture(b.Textures[0].baseTexture.ID)
	color := b.Textures[0].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(0.0, 1.0, 0.0) // Normal Pointing Up
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad.
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad Bottom Face

	// ===================================
	// ========== Bottom Face ============
	// ===================================
	rl.SetTexture(b.Textures[1].baseTexture.ID)
	color = b.Textures[1].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(0.0, -1.0, 0.0) // Normal Pointing Down
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad

	// ===================================
	// ========== Front Face =============
	// ===================================
	rl.SetTexture(b.Textures[2].baseTexture.ID)
	color = b.Textures[2].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(0.0, 0.0, 1.0) // Normal Pointing Towards Viewer
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Left Of The Texture and Quad

	if b.Textures[2].overlayTexture.ID != 0 {
		// ===================================
		// ========== Front Overlay ==========
		// ===================================
		rl.SetTexture(b.Textures[2].overlayTexture.ID)
		color = b.Textures[2].overlayTint
		rl.Color4ub(color.R, color.G, color.B, color.A)

		rl.Normal3f(0.0, 0.0, 1.0) // Normal Pointing Towards Viewer
		rl.TexCoord2f(0.0, 1.0)
		rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1.0, 1.0)
		rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1.0, 0.0)
		rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0.0, 0.0)
		rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Left Of The Texture and Quad
	}

	// ===================================
	// ========== Right Face =============
	// ===================================
	rl.SetTexture(b.Textures[3].baseTexture.ID)
	color = b.Textures[3].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(1.0, 0.0, 0.0) // Normal Pointing Right
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad

	if b.Textures[3].overlayTexture.ID != 0 {
		// ====================================
		// ========== Right Overlay ===========
		// ====================================
		rl.SetTexture(b.Textures[3].overlayTexture.ID)
		color = b.Textures[3].overlayTint
		rl.Color4ub(color.R, color.G, color.B, color.A)

		rl.Normal3f(1.0, 0.0, 0.0) // Normal Pointing Right
		rl.TexCoord2f(1.0, 1.0)
		rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1.0, 0.0)
		rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0.0, 0.0)
		rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0.0, 1.0)
		rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Left Of The Texture and Quad
	}

	// ===================================
	// ========== Back Face ==============
	// ===================================
	rl.SetTexture(b.Textures[4].baseTexture.ID)
	color = b.Textures[4].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(0.0, 0.0, -1.0) // Normal Pointing Away From Viewer
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Left Of The Texture and Quad

	if b.Textures[4].overlayTexture.ID != 0 {
		// ====================================
		// ========== Back Overlay ============
		// ====================================
		rl.SetTexture(b.Textures[4].overlayTexture.ID)
		color = b.Textures[4].overlayTint
		rl.Color4ub(color.R, color.G, color.B, color.A)

		rl.Normal3f(0.0, 0.0, -1.0) // Normal Pointing Away From Viewer
		rl.TexCoord2f(1.0, 1.0)
		rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1.0, 0.0)
		rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0.0, 0.0)
		rl.Vertex3f(x+dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0.0, 1.0)
		rl.Vertex3f(x+dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Left Of The Texture and Quad
	}

	// ===================================
	// ========== Left Face ==============
	// ===================================
	rl.SetTexture(b.Textures[5].baseTexture.ID)
	color = b.Textures[5].baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)

	rl.Normal3f(-1.0, 0.0, 0.0) // Normal Pointing Left
	rl.TexCoord2f(0.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Left Of The Texture and Quad
	rl.TexCoord2f(1.0, 1.0)
	rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad
	rl.TexCoord2f(1.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Right Of The Texture and Quad
	rl.TexCoord2f(0.0, 0.0)
	rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad

	if b.Textures[5].overlayTexture.ID != 0 {
		// ====================================
		// ========== Left Overlay ============
		// ====================================
		rl.SetTexture(b.Textures[5].overlayTexture.ID)
		color = b.Textures[5].overlayTint
		rl.Color4ub(color.R, color.G, color.B, color.A)

		rl.Normal3f(-1.0, 0.0, 0.0) // Normal Pointing Left
		rl.TexCoord2f(0.0, 1.0)
		rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z-dimensions.Z/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1.0, 1.0)
		rl.Vertex3f(x-dimensions.X/2, y-dimensions.Y/2, z+dimensions.Z/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1.0, 0.0)
		rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z+dimensions.Z/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0.0, 0.0)
		rl.Vertex3f(x-dimensions.X/2, y+dimensions.Y/2, z-dimensions.Z/2) // Top Left Of The Texture and Quad
	}

	rl.End()

	rl.SetTexture(0)
}
