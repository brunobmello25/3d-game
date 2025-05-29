package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Chunk struct {
	Width, Height, Depth int
	blocks               []Block
}

func NewChunk(width, height, depth int) *Chunk {
	c := &Chunk{
		Width:  width,
		Height: height,
		Depth:  depth,
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
	return x + c.Width*(y+c.Height*z)
}

func (c *Chunk) SetBlock(x, y, z int, block Block) int {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height || z < 0 || z >= c.Depth {
		panic("Index out of bounds")
	}

	idx := c.idx(x, y, z)
	c.blocks[idx] = block
	return idx
}

func (c *Chunk) Draw() {
	for x := range c.Width {
		for y := range c.Height {
			for z := range c.Depth {
				block := c.blocks[c.idx(x, y, z)]
				if block.Type != BlockAir {
					pos := rl.NewVector3(float32(x), float32(y), float32(z))
					block.Draw(pos)
				}
			}
		}
	}
}
