package render

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Renderer struct {
	textureManager *TextureManager
}

func NewRenderer() *Renderer {
	return &Renderer{
		textureManager: NewTextureManager(),
	}
}

func (r *Renderer) InitTextures() {
	// Load all textures
	textures := map[string]string{
		"stone":      "assets/blocks/stone.png",
		"dirt":       "assets/blocks/dirt.png",
		"grass_top":  "assets/blocks/grass_top.png",
		"grass_side": "assets/blocks/grass_side_overlay.png",
	}

	for name, path := range textures {
		r.textureManager.LoadTexture(name, path)
	}
}

func (r *Renderer) Cleanup() {
	r.textureManager.Cleanup()
}

func (r *Renderer) GetTexture(name string) rl.Texture2D {
	return r.textureManager.GetTexture(name)
}

func (r *Renderer) DrawFPS(screenW int32) {
	fps := rl.GetFPS()
	text := fmt.Sprintf("%d FPS", fps)
	fontSize := int32(20)

	// Measure the width of the text so we can right-align it
	textWidth := rl.MeasureText(text, fontSize)

	// draw it with a little padding from the edges:
	x := screenW - textWidth - 10
	y := int32(10)

	rl.DrawText(text, x, y, fontSize, rl.Black)
}

func (r *Renderer) DrawPosition(position rl.Vector3) {
	text := fmt.Sprintf("Position: %.1f, %.1f, %.1f", position.X, position.Y, position.Z)
	fontSize := int32(20)

	// Draw position in the top-left corner with padding
	x := int32(10)
	y := int32(10)

	rl.DrawText(text, x, y, fontSize, rl.Black)
}
