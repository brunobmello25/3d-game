package block

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/texture"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	// TODO: remove this hardcoded texture
	dirtTexture := texture.GetTexture(texture.TEXTURE_NAME_DIRT)

	faces := [6]BlockFace{
		{direction: "up", texture: dirtTexture},
		{direction: "down", texture: dirtTexture},
		{direction: "front", texture: dirtTexture},
		{direction: "back", texture: dirtTexture},
		{direction: "left", texture: dirtTexture},
		{direction: "right", texture: dirtTexture},
	}

	return Block{
		Type:       blockType,
		Visibility: visiblity,
		Faces:      faces,
	}
}
