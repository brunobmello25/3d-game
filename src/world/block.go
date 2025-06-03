package world

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/texture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockVisibility int

func BlockVisibilityFromType(blockType BlockType) BlockVisibility {
	switch blockType {
	case BlockAir:
		return BlockVisibilityEmpty
	case BlockStone, BlockDirt, BlockGrass:
		return BlockVisibilityOpaque
	}
	panic(fmt.Sprintf("Unknown block type for visibility: %d", blockType))
}

func (bv BlockVisibility) IsEmpty() bool {
	return bv == BlockVisibilityEmpty
}

func (bv BlockVisibility) IsTransparent() bool {
	return bv == BlockVisibilityTransparent
}

func (bv BlockVisibility) IsOpaque() bool {
	return bv == BlockVisibilityOpaque
}

const (
	BlockVisibilityEmpty       BlockVisibility = iota // no block, usually air
	BlockVisibilityTransparent                        // glass, leaves, etc
	BlockVisibilityOpaque                             // solid blocks like stone, dirt, grass, etc
)

type BlockType int

const (
	BlockAir BlockType = iota
	BlockStone
	BlockDirt
	BlockGrass
)

type Block struct {
	Type       BlockType
	Visibility BlockVisibility
	Faces      [6]Face // top, bottom, front, right, back, left
}

func NewBlock(blockType BlockType) Block {
	visibility := BlockVisibilityFromType(blockType)
	block := Block{
		Type:       blockType,
		Visibility: visibility,
	}

	switch blockType {
	case BlockStone:
		stoneColor := rl.NewColor(143, 143, 143, 255)
		stoneTexture := texture.GetTexture(texture.TEXTURE_NAME_STONE)
		stoneFaceTexture := NewBlockTextureBase(stoneTexture, stoneColor)

		for i := range 6 {
			block.Faces[i] = NewFace(rl.Vector3{}, *stoneFaceTexture, faceDirections[i])
		}
		return block

	case BlockDirt:
		dirtTexture := texture.GetTexture(texture.TEXTURE_NAME_DIRT)
		dirtFaceTexture := NewBlockTextureBase(dirtTexture, rl.Brown)

		for i := range 6 {
			block.Faces[i] = NewFace(rl.Vector3{}, *dirtFaceTexture, faceDirections[i])
		}
		return block

	case BlockGrass:
		// Top face
		grassTopTexture := texture.GetTexture(texture.TEXTURE_NAME_GRASS_TOP)
		block.Faces[0] = NewFace(rl.Vector3{}, *NewBlockTextureBase(grassTopTexture, rl.DarkGreen), faceDirections[0])

		// Bottom face
		dirtTexture := texture.GetTexture(texture.TEXTURE_NAME_DIRT)
		block.Faces[1] = NewFace(rl.Vector3{}, *NewBlockTextureBase(dirtTexture, rl.Brown), faceDirections[1])

		// Side faces
		baseTexture := texture.GetTexture(texture.TEXTURE_NAME_DIRT)
		overlayTexture := texture.GetTexture(texture.TEXTURE_NAME_GRASS_SIDE_OVERLAY)
		sideFaceTexture := NewBlockTextureWithOverlay(baseTexture, rl.Brown, overlayTexture, rl.DarkGreen)

		for i := 2; i < 6; i++ {
			block.Faces[i] = NewFace(rl.Vector3{}, *sideFaceTexture, faceDirections[i])
		}
		return block

	case BlockAir:
		return block
	}

	panic(fmt.Sprintf("Unknown block type: %d", blockType))
}

func (b *Block) Draw(pos rl.Vector3) {
	if b.Visibility.IsEmpty() {
		return
	}

	rl.Begin(rl.Quads)
	for _, face := range b.Faces {
		face.Position = pos
		face.Draw()
	}
	rl.End()
	rl.SetTexture(0)
}

var faceDirections = [6]string{
	"up",
	"down",
	"front",
	"right",
	"back",
	"left",
}
