package block

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/texture"
)

type BlockType int

const (
	BlockTypeAir BlockType = iota
	BlockTypeStone
	BlockTypeDirt
	BlockTypeGrass
)

type BlockVisibility int

const (
	BlockVisibilityEmpty BlockVisibility = iota
	BlockVisibilityTransparent
	BlockVisibilityOpaque
)

func VisibilityFromType(blockType BlockType) BlockVisibility {
	switch blockType {
	case BlockTypeAir:
		return BlockVisibilityEmpty
	case BlockTypeStone, BlockTypeDirt, BlockTypeGrass:
		return BlockVisibilityOpaque
	}

	panic(fmt.Sprintf("unknown block type: %d", blockType))
}

type Block struct {
	Type       BlockType
	Visibility BlockVisibility
	Faces      [6]BlockFace // top, bottom, front, right, back, left
}

func NewBlock(blockType BlockType) Block {
	visiblity := VisibilityFromType(blockType)

	var faces [6]BlockFace
	switch blockType {
	case BlockTypeAir:
		faces = [6]BlockFace{} // Air block has no faces
	case BlockTypeStone:
		faces = stoneBlockFaces()
	case BlockTypeDirt:
		faces = dirtBlockFaces()
	case BlockTypeGrass:
		faces = grassBlockFaces()
	default:
		panic(fmt.Sprintf("unknown block type: %d", blockType))
	}

	return Block{
		Type:       blockType,
		Visibility: visiblity,
		Faces:      faces,
	}
}

func dirtBlockFaces() [6]BlockFace {
	faces := [6]BlockFace{
		NewFace(FacingDirectionUp, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionDown, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionFront, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionBack, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionLeft, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionRight, texture.TEXTURE_NAME_DIRT),
	}
	return faces
}

func stoneBlockFaces() [6]BlockFace {
	faces := [6]BlockFace{
		NewFace(FacingDirectionUp, texture.TEXTURE_NAME_STONE),
		NewFace(FacingDirectionDown, texture.TEXTURE_NAME_STONE),
		NewFace(FacingDirectionFront, texture.TEXTURE_NAME_STONE),
		NewFace(FacingDirectionBack, texture.TEXTURE_NAME_STONE),
		NewFace(FacingDirectionLeft, texture.TEXTURE_NAME_STONE),
		NewFace(FacingDirectionRight, texture.TEXTURE_NAME_STONE),
	}
	return faces
}

func grassBlockFaces() [6]BlockFace {
	// TODO: texture overlay
	faces := [6]BlockFace{
		NewFace(FacingDirectionUp, texture.TEXTURE_NAME_GRASS_TOP),
		NewFace(FacingDirectionDown, texture.TEXTURE_NAME_DIRT),
		NewFace(FacingDirectionFront, texture.TEXTURE_NAME_GRASS_SIDE),
		NewFace(FacingDirectionBack, texture.TEXTURE_NAME_GRASS_SIDE),
		NewFace(FacingDirectionLeft, texture.TEXTURE_NAME_GRASS_SIDE),
		NewFace(FacingDirectionRight, texture.TEXTURE_NAME_GRASS_SIDE),
	}
	return faces
}
