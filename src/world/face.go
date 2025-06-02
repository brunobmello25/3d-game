package world

import rl "github.com/gen2brain/raylib-go/raylib"

type Face struct {
	Position rl.Vector3
	Texture  FaceTexture
	Normal   rl.Vector3
}

type VertexData struct {
	texCoords [4]rl.Vector2
	vertices  [4]rl.Vector3
}

var faceVertexData = map[rl.Vector3]VertexData{
	rl.NewVector3(0, 1, 0): { // Top
		texCoords: [4]rl.Vector2{
			{X: 0.0, Y: 1.0},
			{X: 0.0, Y: 0.0},
			{X: 1.0, Y: 0.0},
			{X: 1.0, Y: 1.0},
		},
		vertices: [4]rl.Vector3{
			{X: -0.5, Y: 0.5, Z: -0.5},
			{X: -0.5, Y: 0.5, Z: 0.5},
			{X: 0.5, Y: 0.5, Z: 0.5},
			{X: 0.5, Y: 0.5, Z: -0.5},
		},
	},
	rl.NewVector3(0, -1, 0): { // Bottom
		texCoords: [4]rl.Vector2{
			{X: 1.0, Y: 1.0},
			{X: 0.0, Y: 1.0},
			{X: 0.0, Y: 0.0},
			{X: 1.0, Y: 0.0},
		},
		vertices: [4]rl.Vector3{
			{X: -0.5, Y: -0.5, Z: -0.5},
			{X: 0.5, Y: -0.5, Z: -0.5},
			{X: 0.5, Y: -0.5, Z: 0.5},
			{X: -0.5, Y: -0.5, Z: 0.5},
		},
	},
	rl.NewVector3(0, 0, 1): { // Front
		texCoords: [4]rl.Vector2{
			{X: 0.0, Y: 1.0},
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 0.0},
			{X: 0.0, Y: 0.0},
		},
		vertices: [4]rl.Vector3{
			{X: -0.5, Y: -0.5, Z: 0.5},
			{X: 0.5, Y: -0.5, Z: 0.5},
			{X: 0.5, Y: 0.5, Z: 0.5},
			{X: -0.5, Y: 0.5, Z: 0.5},
		},
	},
	rl.NewVector3(1, 0, 0): { // Right
		texCoords: [4]rl.Vector2{
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 0.0},
			{X: 0.0, Y: 0.0},
			{X: 0.0, Y: 1.0},
		},
		vertices: [4]rl.Vector3{
			{X: 0.5, Y: -0.5, Z: -0.5},
			{X: 0.5, Y: 0.5, Z: -0.5},
			{X: 0.5, Y: 0.5, Z: 0.5},
			{X: 0.5, Y: -0.5, Z: 0.5},
		},
	},
	rl.NewVector3(0, 0, -1): { // Back
		texCoords: [4]rl.Vector2{
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 0.0},
			{X: 0.0, Y: 0.0},
			{X: 0.0, Y: 1.0},
		},
		vertices: [4]rl.Vector3{
			{X: -0.5, Y: -0.5, Z: -0.5},
			{X: -0.5, Y: 0.5, Z: -0.5},
			{X: 0.5, Y: 0.5, Z: -0.5},
			{X: 0.5, Y: -0.5, Z: -0.5},
		},
	},
	rl.NewVector3(-1, 0, 0): { // Left
		texCoords: [4]rl.Vector2{
			{X: 0.0, Y: 1.0},
			{X: 1.0, Y: 1.0},
			{X: 1.0, Y: 0.0},
			{X: 0.0, Y: 0.0},
		},
		vertices: [4]rl.Vector3{
			{X: -0.5, Y: -0.5, Z: -0.5},
			{X: -0.5, Y: -0.5, Z: 0.5},
			{X: -0.5, Y: 0.5, Z: 0.5},
			{X: -0.5, Y: 0.5, Z: -0.5},
		},
	},
}

func NewFace(blockPosition rl.Vector3, texture FaceTexture, facing string) *Face {
	return &Face{
		Position: blockPosition,
		Texture:  texture,
		Normal:   normalFromFacing(facing),
	}
}

func (f *Face) Draw() {
	vertexData, ok := faceVertexData[f.Normal]
	if !ok {
		panic("Invalid face normal")
	}

	rl.SetTexture(f.Texture.baseTexture.ID)
	color := f.Texture.baseTint
	rl.Color4ub(color.R, color.G, color.B, color.A)
	rl.Normal3f(f.Normal.X, f.Normal.Y, f.Normal.Z)

	for i := range 4 {
		rl.TexCoord2f(vertexData.texCoords[i].X, vertexData.texCoords[i].Y)
		vertex := vertexData.vertices[i]
		rl.Vertex3f(
			f.Position.X+vertex.X,
			f.Position.Y+vertex.Y,
			f.Position.Z+vertex.Z,
		)
	}

	if f.Texture.overlayTexture.ID == 0 {
		return
	}

	rl.SetTexture(f.Texture.overlayTexture.ID)
	color = f.Texture.overlayTint
	rl.Color4ub(color.R, color.G, color.B, color.A)
	rl.Normal3f(f.Normal.X, f.Normal.Y, f.Normal.Z)

	for i := range 4 {
		rl.TexCoord2f(vertexData.texCoords[i].X, vertexData.texCoords[i].Y)
		vertex := vertexData.vertices[i]
		rl.Vertex3f(
			f.Position.X+vertex.X,
			f.Position.Y+vertex.Y,
			f.Position.Z+vertex.Z,
		)
	}
}

func normalFromFacing(facing string) rl.Vector3 {
	switch facing {
	case "up":
		return rl.NewVector3(0, 1, 0)
	case "down":
		return rl.NewVector3(0, -1, 0)
	case "left":
		return rl.NewVector3(-1, 0, 0)
	case "right":
		return rl.NewVector3(1, 0, 0)
	case "front":
		return rl.NewVector3(0, 0, 1)
	case "back":
		return rl.NewVector3(0, 0, -1)
	}

	panic("Invalid facing direction: " + facing)
}
