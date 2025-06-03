package block

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/texture"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockType int

const (
	BlockTypeAir BlockType = iota
	BlockTypeStone
	BlockTypeDirt
	BlockTypeGrass
)

type BlockVisibility int

const (
	BlockVisibilityEmpty BlockVisibility = iota
	BlockVisibilityTransparent
	BlockVisibilityOpaque
)

func VisibilityFromType(blockType BlockType) BlockVisibility {
	switch blockType {
	case BlockTypeAir:
		return BlockVisibilityEmpty
	case BlockTypeStone, BlockTypeDirt, BlockTypeGrass:
		return BlockVisibilityOpaque
	}

	panic(fmt.Sprintf("unknown block type: %d", blockType))
}

type Block struct {
	Type       BlockType
	Visibility BlockVisibility
	Faces      [6]BlockFace // top, bottom, front, right, back, left
}

func NewBlock(blockType BlockType) Block {
	visiblity := VisibilityFromType(blockType)

	// TODO: remove this hardcoded texture
	dirtTexture := texture.GetTexture(texture.TEXTURE_NAME_DIRT)

	faces := [6]BlockFace{
		{direction: "up", texture: dirtTexture},
		{direction: "down", texture: dirtTexture},
		{direction: "front", texture: dirtTexture},
		{direction: "back", texture: dirtTexture},
		{direction: "left", texture: dirtTexture},
		{direction: "right", texture: dirtTexture},
	}

	return Block{
		Type:       blockType,
		Visibility: visiblity,
		Faces:      faces,
	}
}

// BuildBlockMesh takes the block center and the six faces,
// and returns a raylib.Mesh that, when uploaded, lives on the GPU.
func (b Block) BuildMesh(
	blockCenter rl.Vector3,
) rl.Mesh {
	// these three slices will hold all of our vertex data:
	var positions []float32 // x,y,z triples
	var texcoords []float32 // u,v pairs
	var normals []float32   // nx,ny,nz triples

	// for each face we know:
	//   - face.getVertexData() → [6]Vector3 tri‐list of corner positions
	//   - face.direction.getNormal() → Vector3 normal
	// and for UVs we’ll just use the full [0..1]×[0..1] square
	// (replace this with atlas‐sub‐region if you have one)
	uvBL := []float32{0, 0}
	uvBR := []float32{1, 0}
	uvTR := []float32{1, 1}
	uvTL := []float32{0, 1}

	for _, face := range b.Faces {
		// 1) fetch the 6 points (two triangles)
		faceVertices := face.getVertexData(blockCenter)
		// 2) fetch the face‐normal (same for all 6)
		n := face.direction.getNormal()
		// 3) assign UV coordinates in the same tri‐order
		textureUVs := [][]float32{
			uvBL, uvBR, uvTR, // tri1: BL→BR→TR
			uvBL, uvTR, uvTL, // tri2: BL→TR→TL
		}

		// 4) append to our big arrays
		for i := range 6 {
			faceVertex := faceVertices[i]
			positions = append(positions, faceVertex.X, faceVertex.Y, faceVertex.Z)
			normals = append(normals, float32(n.X), float32(n.Y), float32(n.Z))
			texcoords = append(texcoords, textureUVs[i][0], textureUVs[i][1])
		}
	}

	// 5) build the raylib.Mesh
	mesh := rl.Mesh{}
	mesh.VertexCount = int32(len(positions) / 3)
	mesh.Vertices = &positions[0]
	mesh.Texcoords = &texcoords[0]
	mesh.Normals = &normals[0]
	// (you can also set mesh.Indices if you want indexed drawing)

	// 6) push it up to the GPU
	rl.UploadMesh(&mesh, false) // false = keep CPU data around

	return mesh
}
