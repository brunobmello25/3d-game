package block

import "fmt"

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

	return Block{
		Type:       blockType,
		Visibility: visiblity,
	}
}
