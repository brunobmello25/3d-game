package texture

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	TEXTURE_NAME_STONE      = "stone"
	TEXTURE_NAME_DIRT       = "dirt"
	TEXTURE_NAME_GRASS_TOP  = "grass_top"
	TEXTURE_NAME_GRASS_SIDE = "grass_side"
)

// Manager handles loading and accessing textures
type Manager struct {
	textures    map[string]rl.Texture2D
	initialized bool
}

var instance *Manager

// Init initializes the texture manager
func Init() {
	tm := getInstance()

	if tm.initialized {
		return
	}

	tm.textures = make(map[string]rl.Texture2D)
	if tm.textures == nil {
		panic("failed to initialize texture manager")
	}

	loadTexture(TEXTURE_NAME_STONE)
	loadTexture(TEXTURE_NAME_DIRT)
	loadTexture(TEXTURE_NAME_GRASS_TOP)
	loadTexture(TEXTURE_NAME_GRASS_SIDE)

	tm.initialized = true
}

// GetTexture returns a texture by name
func GetTexture(name string) rl.Texture2D {
	tm := getInstance()

	texture, ok := tm.textures[name]
	if !ok {
		panic(fmt.Sprintf("texture not found: %s", name))
	}
	return texture
}

// Cleanup unloads all textures
func Cleanup() {
	tm := getInstance()

	for _, texture := range tm.textures {
		rl.UnloadTexture(texture)
	}
	tm.textures = make(map[string]rl.Texture2D)
}

// getInstance returns the singleton instance of Manager
func getInstance() *Manager {
	if instance == nil {
		instance = &Manager{
			textures: make(map[string]rl.Texture2D),
		}
	}
	return instance
}

// loadTexture loads a texture and stores it with the given name
func loadTexture(file_name string) {
	tm := getInstance()

	path := fmt.Sprintf("assets/blocks/%s.png", file_name)
	texture := rl.LoadTexture(path)
	if texture.ID == 0 {
		panic(fmt.Sprintf("failed to load texture: %s", file_name))
	}
	tm.textures[file_name] = texture
}
