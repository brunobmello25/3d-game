package mesh

import (
	"unsafe"

	"github.com/brunobmello25/3d-game/src/block"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MeshBuilder struct {
	vertices    []float32
	normals     []float32
	texcoords   []float32
	indices     []uint16
	vertexCount int
}

func NewMeshBuilder() *MeshBuilder {
	return &MeshBuilder{
		vertices:    make([]float32, 0),
		normals:     make([]float32, 0),
		texcoords:   make([]float32, 0),
		indices:     make([]uint16, 0),
		vertexCount: 0,
	}
}

func (mb *MeshBuilder) AddFace(face block.BlockFace, faceCenter rl.Vector3) {
	// Add vertices
	mb.vertices = append(mb.vertices, face.GetVertexCoords(faceCenter)...)

	// Add normals
	mb.normals = append(mb.normals, face.GetVertexNormals(faceCenter)...)

	// Add texture coordinates
	mb.texcoords = append(mb.texcoords, face.GetTextureCoords()...)

	// Add indices (assuming each face is a quad with 2 triangles)
	baseIndex := uint16(mb.vertexCount)
	mb.indices = append(mb.indices,
		baseIndex, baseIndex+1, baseIndex+2, // First triangle
		baseIndex, baseIndex+2, baseIndex+3, // Second triangle
	)

	mb.vertexCount += 4 // Each face adds 4 vertices
}

func (mb *MeshBuilder) Build() rl.Mesh {
	mesh := rl.Mesh{
		VertexCount:   int32(mb.vertexCount),
		TriangleCount: int32(len(mb.indices) / 3),
	}

	// Convert slices to unsafe pointers
	mesh.Vertices = unsafe.SliceData(mb.vertices)
	mesh.Normals = unsafe.SliceData(mb.normals)
	mesh.Texcoords = unsafe.SliceData(mb.texcoords)
	mesh.Indices = unsafe.SliceData(mb.indices)

	// Upload mesh data to GPU
	rl.UploadMesh(&mesh, false)

	return mesh
}

func (mb *MeshBuilder) Clear() {
	mb.vertices = make([]float32, 0)
	mb.normals = make([]float32, 0)
	mb.texcoords = make([]float32, 0)
	mb.indices = make([]uint16, 0)
	mb.vertexCount = 0
}
