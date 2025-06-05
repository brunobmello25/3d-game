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

	c.rebuildMesh()

	return c
}

func (c *Chunk) SetBlock(x, y, z int, b block.Block) {
	if x < 0 || x >= CHUNK_SIZE || y < 0 || y >= CHUNK_SIZE || z < 0 || z >= CHUNK_SIZE {
		panic(fmt.Sprintf("Coordinates out of bounds: (%d, %d, %d)", x, y, z))
	}

	index := c.linearize(x, y, z)
	c.Blocks[index] = b
	c.dirty = true
}

func (c *Chunk) Update() {
	if c.dirty {
		c.rebuildMesh()
	}
}

func (c *Chunk) Render() {
	rl.DrawModel(c.model, rl.NewVector3(c.Position.X*CHUNK_SIZE, c.Position.Y*CHUNK_SIZE, c.Position.Z*CHUNK_SIZE), 1.0, rl.White)
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
		x, y, z := c.Delinearize(i)

		localCoord := rl.NewVector3(float32(x), float32(y), float32(z))

		b := c.Blocks[i]

		if !b.Visibility.IsEmpty() {
			for _, face := range b.Faces {
				neighbor := c.getNeighbor(localCoord, face)
				if neighbor == nil || neighbor.Visibility.IsEmpty() {
					meshBuilder.AddFace(face, localCoord)
				}
			}
		}
	}

	c.mesh = meshBuilder.Build()
	c.model = rl.LoadModelFromMesh(c.mesh)
	rl.SetMaterialTexture(c.model.Materials, rl.MapDiffuse, texture.GetAtlasTexture())

	c.dirty = false
}

func (c *Chunk) getNeighbor(currentCoord rl.Vector3, currentFace block.BlockFace) *block.Block {
	normal := currentFace.Direction.GetNormal()
	neighborCoord := rl.NewVector3(
		currentCoord.X+normal.X,
		currentCoord.Y+normal.Y,
		currentCoord.Z+normal.Z,
	)

	if neighborCoord.X < 0 || neighborCoord.X >= CHUNK_SIZE ||
		neighborCoord.Y < 0 || neighborCoord.Y >= CHUNK_SIZE ||
		neighborCoord.Z < 0 || neighborCoord.Z >= CHUNK_SIZE {
		return nil
	}

	neighborIndex := c.linearize(int(neighborCoord.X), int(neighborCoord.Y), int(neighborCoord.Z))
	return &c.Blocks[neighborIndex]
}

func (c *Chunk) Delinearize(index int) (int, int, int) {
	x := index % CHUNK_SIZE
	y := (index / CHUNK_SIZE) % CHUNK_SIZE
	z := index / (CHUNK_SIZE * CHUNK_SIZE)
	return x, y, z
}

func (c *Chunk) linearize(x, y, z int) int {
	if x < 0 || x >= CHUNK_SIZE || y < 0 || y >= CHUNK_SIZE || z < 0 || z >= CHUNK_SIZE {
		panic(fmt.Sprintf("Coordinates out of bounds: (%d, %d, %d)", x, y, z))
	}
	return x + y*CHUNK_SIZE + z*CHUNK_SIZE*CHUNK_SIZE
}
