package world

import rl "github.com/gen2brain/raylib-go/raylib"

type FaceTexture struct {
	baseTexture    rl.Texture2D
	baseTint       rl.Color
	overlayTexture rl.Texture2D
	overlayTint    rl.Color
}

func NewBlockTextureBase(texture rl.Texture2D, tint rl.Color) *FaceTexture {
	return &FaceTexture{
		baseTexture: texture,
		baseTint:    tint,
	}
}

func NewBlockTextureWithOverlay(baseTexture rl.Texture2D, baseTint rl.Color, overlayTexture rl.Texture2D, overlayTint rl.Color) *FaceTexture {
	return &FaceTexture{
		baseTexture:    baseTexture,
		baseTint:       baseTint,
		overlayTexture: overlayTexture,
		overlayTint:    overlayTint,
	}
}
