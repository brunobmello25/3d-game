package render

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TextureManager handles loading and accessing textures
type TextureManager struct {
	textures map[string]rl.Texture2D
}

// NewTextureManager creates a new texture manager
func NewTextureManager() *TextureManager {
	return &TextureManager{
		textures: make(map[string]rl.Texture2D),
	}
}

// LoadTexture loads a texture and stores it with the given name
func (tm *TextureManager) LoadTexture(name, path string) {
	texture := rl.LoadTexture(path)
	if texture.ID == 0 {
		panic(fmt.Sprintf("failed to load texture: %s", path))
	}
	tm.textures[name] = texture
}

// GetTexture returns a texture by name
func (tm *TextureManager) GetTexture(name string) rl.Texture2D {
	texture, ok := tm.textures[name]
	if !ok {
		panic(fmt.Sprintf("texture not found: %s", name))
	}
	return texture
}

// Cleanup unloads all textures
func (tm *TextureManager) Cleanup() {
	for _, texture := range tm.textures {
		rl.UnloadTexture(texture)
	}
	tm.textures = make(map[string]rl.Texture2D)
}
