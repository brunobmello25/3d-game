package chunk

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/block"
	"github.com/brunobmello25/3d-game/src/mesh"
	"github.com/brunobmello25/3d-game/src/texture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CHUNK_SIZE = 16 // 16x16x16 blocks

type Chunk struct {
	Blocks   [CHUNK_SIZE * CHUNK_SIZE * CHUNK_SIZE]block.Block
	Position rl.Vector3
	dirty    bool
	mesh     rl.Mesh
	model    rl.Model
}

func NewChunk(position rl.Vector3) *Chunk {
	c := &Chunk{
		Position: position,
	}

	for i := range c.Blocks {
		x, y, z := c.delinearize(i)
		fmt.Println("Block position:", x, y, z)

		btype := block.BlockTypeAir // Default to air block

		if y == 0 {
			btype = block.BlockTypeStone // Set stone for the bottom layer
		} else if y >= 1 && y <= 14 {
			btype = block.BlockTypeDirt // Set dirt for the middle layers
		} else if y == 15 {
			btype = block.BlockTypeGrass // Set grass for the top layer
		}
		c.Blocks[i] = block.NewBlock(btype)
	}

	c.rebuildMesh()
	c.dirty = false

	return c
}

func (c *Chunk) Update() {
	if c.dirty {
		fmt.Println("===================")
		fmt.Println("Rebuilding chunk mesh...")
		fmt.Println("===================")
		c.rebuildMesh()
	}
}

func (c *Chunk) Render() {
	rl.DrawModel(c.model, c.Position, 1.0, rl.White)
}

func (c *Chunk) Cleanup() {
	if c.model.Meshes != nil {
		rl.UnloadModel(c.model)
	}
	if c.mesh.VertexCount > 0 {
		rl.UnloadMesh(&c.mesh)
	}
}

func (c *Chunk) rebuildMesh() {
	meshBuilder := mesh.NewMeshBuilder()

	for i := range len(c.Blocks) {
		x, y, z := c.delinearize(i)

		blockCenter := rl.NewVector3(
			c.Position.X*CHUNK_SIZE+float32(x),
			c.Position.Y*CHUNK_SIZE+float32(y),
			c.Position.Z*CHUNK_SIZE+float32(z),
		)

		b := c.Blocks[i]
		// TODO: only add face if it should actually be rendered
		if !b.Visibility.IsEmpty() {
			meshBuilder.AddFaces(b.Faces[:], blockCenter)
		}
	}

	c.mesh = meshBuilder.Build()
	c.model = rl.LoadModelFromMesh(c.mesh)
	rl.SetMaterialTexture(c.model.Materials, rl.MapDiffuse, texture.GetAtlasTexture())

	c.dirty = false
}

func (c *Chunk) delinearize(index int) (int, int, int) {
	x := index % CHUNK_SIZE
	y := (index / CHUNK_SIZE) % CHUNK_SIZE
	z := index / (CHUNK_SIZE * CHUNK_SIZE)
	return x, y, z
}
