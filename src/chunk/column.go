package chunk

import (
	"math"

	"github.com/brunobmello25/3d-game/src/block"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CHUNK_COLUMN_HEIGHT = 16

type ChunkColumn struct {
	Chunks [CHUNK_COLUMN_HEIGHT]Chunk
	X, Z   int
}

func NewChunkColumn(X, Z int) *ChunkColumn {
	column := &ChunkColumn{
		X: X,
		Z: Z,
	}
	for i := range CHUNK_COLUMN_HEIGHT {
		position := rl.NewVector3(float32(X), float32(i), float32(Z))
		column.Chunks[i] = *NewChunk(position)
		column.Chunks[i].Position = position
	}
	return column
}

func (cc *ChunkColumn) Generate() {
	for i := range cc.Chunks {
		for j := range cc.Chunks[i].Blocks {
			x, y, z := cc.Chunks[i].Delinearize(j)
			height := simpleNoise(x, z) // Use a noise function to determine height
			if y < int(height) {
				cc.Chunks[i].SetBlock(x, y, z, block.NewBlock(block.BlockTypeStone))
			} else if y == int(height) {
				cc.Chunks[i].SetBlock(x, y, z, block.NewBlock(block.BlockTypeGrass))
			} else {
				cc.Chunks[i].SetBlock(x, y, z, block.NewBlock(block.BlockTypeAir))
			}
		}
	}
}

func (cc *ChunkColumn) Update() {
	for i := range cc.Chunks {
		cc.Chunks[i].Update()
	}
}

func (cc *ChunkColumn) Render() {
	for i := range cc.Chunks {
		cc.Chunks[i].Render()
	}
}

func simpleNoise(x, z int) float64 {
	// noise function using sin and cos
	return (math.Sin(float64(x)*0.1) + math.Cos(float64(z)*0.1)) * 10
}
