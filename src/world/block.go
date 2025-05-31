package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockType int

const (
	BlockTypeAir BlockType = iota
	BlockTypeDirt
	BlockTypeGrass
	BlockTypeStone
)

type Block struct {
	Type     BlockType
	Position rl.Vector3
}

func (b *Block) Render() {
	if b.Type == BlockTypeAir {
		return
	}

	var color rl.Color
	switch b.Type {
	case BlockTypeDirt:
		color = rl.Brown
	case BlockTypeGrass:
		color = rl.Green
	case BlockTypeStone:
		color = rl.Gray
	}

	// Draw the block
	rl.DrawCube(
		b.Position,
		1.0, 1.0, 1.0,
		color,
	)
}
