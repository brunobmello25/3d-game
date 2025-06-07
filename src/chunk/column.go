package chunk

import (
	"github.com/brunobmello25/3d-game/src/block"
	"github.com/brunobmello25/3d-game/src/noise"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CHUNK_COLUMN_HEIGHT = 16

// TODO: this is a random value I picked just to test basic terrain generation.
// I'll probably need to make this more advanced sometime
const HEIGHT_OFFSET = 50

type ChunkColumn struct {
	Chunks [CHUNK_COLUMN_HEIGHT]*Chunk
	X, Z   int
}

func NewChunkColumn(X, Z int) *ChunkColumn {
	column := &ChunkColumn{
		X: X,
		Z: Z,
	}
	for i := range CHUNK_COLUMN_HEIGHT {
		position := rl.NewVector3(float32(X), float32(i), float32(Z))
		column.Chunks[i] = NewChunk(position)
		column.Chunks[i].Position = position
	}
	return column
}

func (cc *ChunkColumn) Generate() {
	for i, c := range cc.Chunks {
		for j := range c.Blocks {
			chunkX, chunkY, chunkZ := c.Delinearize(j)
			globalX, globalY, globalZ := cc.localToGlobal(chunkX, chunkY, chunkZ, i)
			block := cc.getBlockForPosition(globalX, globalY, globalZ)
			c.SetBlock(chunkX, chunkY, chunkZ, block)
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

func (cc *ChunkColumn) Unload() {
	for i := range cc.Chunks {
		cc.Chunks[i].Unload()
	}
}

func (cc *ChunkColumn) getBlockForPosition(x, y, z int) block.Block {
	btype := block.BlockTypeAir

	maxHeight := int(noise.Noise2D(x, z) + HEIGHT_OFFSET)

	// from bottom to 20 blocks up should be stone, then dirt up to the second to last block, and the last should be grass
	if y < 20 {
		btype = block.BlockTypeStone
	} else if y < maxHeight-1 {
		btype = block.BlockTypeDirt
	} else if y == maxHeight-1 {
		btype = block.BlockTypeGrass
	}

	return block.NewBlock(btype)
}

func (cc *ChunkColumn) localToGlobal(x, y, z, i int) (int, int, int) {
	return cc.X*CHUNK_SIZE + x, i*CHUNK_SIZE + y, cc.Z*CHUNK_SIZE + z
}
