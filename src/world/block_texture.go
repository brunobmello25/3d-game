package world

import rl "github.com/gen2brain/raylib-go/raylib"

type BlockTexture struct {
	baseTexture    rl.Texture2D
	baseTint       rl.Color
	overlayTexture rl.Texture2D
	overlayTint    rl.Color
}

func NewBlockTextureBase(texture rl.Texture2D, tint rl.Color) *BlockTexture {
	return &BlockTexture{
		baseTexture: texture,
		baseTint:    tint,
	}
}

func NewBlockTextureWithOverlay(baseTexture rl.Texture2D, baseTint rl.Color, overlayTexture rl.Texture2D, overlayTint rl.Color) *BlockTexture {
	return &BlockTexture{
		baseTexture:    baseTexture,
		baseTint:       baseTint,
		overlayTexture: overlayTexture,
		overlayTint:    overlayTint,
	}
}
