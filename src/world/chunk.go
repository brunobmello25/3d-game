package world

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	x, y, z int
	blocks  [][]Block
}

func NewChunk(x, y, z int) *Chunk {
	chunk := &Chunk{
		x:      x,
		y:      y,
		z:      z,
		blocks: make([][]Block, CHUNK_SIZE),
	}

	// Initialize blocks
	for i := range chunk.blocks {
		chunk.blocks[i] = make([]Block, CHUNK_SIZE)
		for j := range chunk.blocks[i] {
			chunk.blocks[i][j] = Block{
				Type: BlockTypeDirt,
				Position: rl.Vector3{
					X: float32(x*CHUNK_SIZE + i),
					Y: 0,
					Z: float32(z*CHUNK_SIZE + j),
				},
			}
		}
	}

	return chunk
}

func (c *Chunk) Render() {
	for i := range c.blocks {
		for j := range c.blocks[i] {
			c.blocks[i][j].Render()
		}
	}
}
