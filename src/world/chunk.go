package world

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
				} else if y > 0 && y < 15 {
					blockType = BlockDirt
				} else if y == 15 {
					blockType = BlockGrass
				}
				c.SetBlock(x, y, z, NewBlock(blockType))
			}
		}
	}

	return c
}

func (c *Chunk) SetBlock(x, y, z int, block Block) int {
	if x < 0 || x >= c.width || y < 0 || y >= c.height || z < 0 || z >= c.depth {
		panic("Index out of bounds")
	}

	idx := c.linearize(x, y, z)
	c.blocks[idx] = block
	return idx
}

func (c *Chunk) Draw() {
	for x := range c.width {
		for y := range c.height {
			for z := range c.depth {
				block := c.blocks[c.linearize(x, y, z)]

				if block.Visibility.IsEmpty() {
					continue
				}

				blockX, blockY, blockZ := float32(c.X*c.width+x), float32(c.Y*c.height+y), float32(c.Z*c.depth+z)
				pos := rl.NewVector3(blockX, blockY, blockZ)
				block.Draw(pos)
			}
		}
	}
}

func (c *Chunk) Get(x, y, z int) Block {
	idx := c.linearize(x, y, z)
	if idx < 0 || idx >= len(c.blocks) {
		panic("Index out of bounds")
	}
	return c.blocks[idx]
}

func (c *Chunk) BuildMesh() {
	// We’ll iterate through all of the voxels in the mesh. If the voxel is empty,
	// we won’t produce any geometry. Otherwise, for each face of the voxel we’ll
	// compare it with the corresponding neighbor. If both are opaque, we won’t generate
	// geometry. If the neighbor is transparent, we’ll compare the voxels and
	// generate a face if they are different.

	for i := range c.blocks {
		x, y, z := c.delinearize(i)
		block := c.blocks[i]

		if block.Visibility == BlockVisibilityEmpty {
			continue
		}

		neighbors := []Block{
			c.Get(x, y-1, z), //top
			c.Get(x, y+1, z), //bottom
			c.Get(x, y, z+1), //front
			c.Get(x+1, y, z), //right
			c.Get(x, y, z-1), //back
			c.Get(x-1, y, z), //back
		}

		for _, neighbor := range neighbors {
			shouldGenerate :=
				block.Visibility.IsOpaque() && neighbor.Visibility.IsEmpty() ||
					block.Visibility.IsOpaque() && neighbor.Visibility.IsTransparent() ||
					block.Visibility.IsTransparent() && neighbor.Visibility.IsEmpty() ||
					block.Visibility.IsTransparent() && neighbor.Visibility.IsTransparent() && block.Type != neighbor.Type

			if !shouldGenerate {
				continue
			}
		}

		panic("TODO")
	}
}

func (c *Chunk) linearize(x, y, z int) int {
	return x + (y * c.width) + (z * c.width * c.height)
}

func (c *Chunk) delinearize(idx int) (int, int, int) {
	if idx < 0 || idx >= len(c.blocks) {
		panic("Index out of bounds")
	}

	x := idx % c.width
	y := (idx / c.width) % c.height
	z := idx / (c.width * c.height)
	return x, y, z
}
