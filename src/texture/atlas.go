package texture

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ATLAS_SIZE = 256 // Size of the texture atlas in pixels
	TILE_SIZE  = 16  // Size of each texture tile in pixels
)

type TextureAtlas struct {
	texture rl.Texture2D
	tiles   map[string]rl.Rectangle // Maps texture name to its UV coordinates in the atlas
}

var atlas *TextureAtlas

func InitAtlas() {
	// Create a blank texture for the atlas
	atlasImage := rl.GenImageColor(ATLAS_SIZE, ATLAS_SIZE, rl.Blank)
	atlas = &TextureAtlas{
		texture: rl.LoadTextureFromImage(atlasImage),
		tiles:   make(map[string]rl.Rectangle),
	}
	rl.UnloadImage(atlasImage)

	// Add textures to the atlas
	// For now, we'll just add the dirt texture as an example
	// You'll want to add all your block textures here
	addTextureToAtlas(TEXTURE_NAME_DIRT)
}

func addTextureToAtlas(textureName string) {
	// Get the individual texture
	texture := GetTexture(textureName)

	// Calculate position in atlas
	// This is a simple grid layout - you might want a more sophisticated packing algorithm
	index := len(atlas.tiles)
	row := index / (ATLAS_SIZE / TILE_SIZE)
	col := index % (ATLAS_SIZE / TILE_SIZE)

	// Calculate UV coordinates
	uv := rl.Rectangle{
		X:      float32(col*TILE_SIZE) / ATLAS_SIZE,
		Y:      float32(row*TILE_SIZE) / ATLAS_SIZE,
		Width:  float32(TILE_SIZE) / ATLAS_SIZE,
		Height: float32(TILE_SIZE) / ATLAS_SIZE,
	}

	// Store the UV coordinates
	atlas.tiles[textureName] = uv

	// Copy the texture into the atlas
	// Note: This is a simplified version. In a real implementation,
	// you'd want to handle this more efficiently, possibly by
	// pre-generating the atlas image with all textures
	textureImage := rl.LoadImageFromTexture(texture)
	atlasImage := rl.LoadImageFromTexture(atlas.texture)

	rl.ImageDraw(
		atlasImage,
		textureImage,
		rl.Rectangle{X: 0, Y: 0, Width: float32(texture.Width), Height: float32(texture.Height)},
		rl.Rectangle{X: float32(col * TILE_SIZE), Y: float32(row * TILE_SIZE), Width: TILE_SIZE, Height: TILE_SIZE},
		rl.White,
	)

	// Update the atlas texture with the new image
	rl.UnloadTexture(atlas.texture)
	atlas.texture = rl.LoadTextureFromImage(atlasImage)

	// Cleanup
	rl.UnloadImage(textureImage)
	rl.UnloadImage(atlasImage)
}

func GetAtlasTexture() rl.Texture2D {
	return atlas.texture
}

func GetTextureUV(textureName string) rl.Rectangle {
	return atlas.tiles[textureName]
}
