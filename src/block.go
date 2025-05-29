package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockType int

const (
	BlockAir BlockType = iota
	BlockStone
	BlockDirt
	BlockGrass
)

type Block struct {
	Type BlockType
}

func NewBlock(t BlockType) *Block {
	return &Block{
		Type: t,
	}
}

func (b *Block) GetColor() rl.Color {
	switch b.Type {
	case BlockAir:
		return rl.Blank
	case BlockDirt:
		return rl.Brown
	case BlockStone:
		return rl.Gray
	case BlockGrass:
		return rl.Green
	}

	panic(fmt.Sprintf("Unknown block type: %d", b.Type))
}

func (b *Block) Draw(pos rl.Vector3) {
	rl.DrawCube(pos, 1, 1, 1, b.GetColor())
	rl.DrawCubeWires(pos, 1, 1, 1, rl.Black)
}
