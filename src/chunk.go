package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	X, Y, Z              int
	width, height, depth int
	blocks               []Block
}

func NewChunk(x, y, z int) *Chunk {
	width, height, depth := CHUNK_SIZE, CHUNK_SIZE, CHUNK_SIZE

	c := &Chunk{
		width:  width,
		height: height,
		depth:  depth,
		X:      x,
		Y:      y,
		Z:      z,
		blocks: make([]Block, width*height*depth),
	}

	for x := range width {
		for y := range height {
			for z := range depth {
				blockType := BlockAir
				if y == 0 {
					blockType = BlockStone
				} else if y > 0 && y <= 10 {
					blockType = BlockDirt
				} else if y > 10 && y < 20 {
					blockType = BlockGrass
				}
				c.SetBlock(x, y, z, *NewBlock(blockType))
			}
		}
	}

	return c
}

func (c *Chunk) idx(x, y, z int) int {
	return x + c.width*(y+c.height*z)
}

func (c *Chunk) SetBlock(x, y, z int, block Block) int {
	if x < 0 || x >= c.width || y < 0 || y >= c.height || z < 0 || z >= c.depth {
		panic("Index out of bounds")
	}

	idx := c.idx(x, y, z)
	c.blocks[idx] = block
	return idx
}

func (c *Chunk) Draw() {
	for x := range c.width {
		for y := range c.height {
			for z := range c.depth {
				block := c.blocks[c.idx(x, y, z)]
				// TODO: shouldn't store air blocks
				if block.Type != BlockAir {
					blockX, blockY, blockZ := float32(c.X*c.width+x), float32(c.Y*c.height+y), float32(c.Z*c.depth+z)
					pos := rl.NewVector3(blockX, blockY, blockZ)
					block.Draw(pos)
				}
			}
		}
	}
}
