package world

type BlockType int

const (
	BlockAir BlockType = iota
	BlockStone
	BlockDirt
	BlockGrass
)

type Block struct {
	Type     BlockType
	Textures [6]*BlockTexture // top, bottom, front, right, back, left
}

func NewBlock(blockType BlockType) *Block {
	block := &Block{
		Type: blockType,
	}

	switch blockType {
	case BlockStone:
		for i := range 6 {
			color := rl.NewColor(143, 143, 143, 255) // Stone color
			block.Textures[i] = NewBlockTextureBase(stoneTexture, color)
		}
		return block
	case BlockDirt:
		for i := range 6 {
			// TODO: proper color
			block.Textures[i] = NewBlockTextureBase(dirtTexture, rl.Brown)
		}
		return block
	case BlockGrass:
		// TODO: proper colors
		block.Textures[0] = NewBlockTextureBase(grassTopTexture, rl.DarkGreen)
		block.Textures[1] = NewBlockTextureBase(dirtTexture, rl.Brown)
		for i := 2; i < 6; i++ {
			block.Textures[i] = NewBlockTextureWithOverlay(dirtTexture, rl.Brown, grassSideOverlayTexture, rl.DarkGreen)
		}
		return block
	case BlockAir:
		return block
	}

	panic(fmt.Sprintf("Unknown block type: %d", blockType))
}
