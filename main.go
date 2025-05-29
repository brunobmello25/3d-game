package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Block byte

const (
	BlockAir Block = iota
	BlockStone
	BlockDirt
	BlockGrass
)

func ColorFromBlock(block Block) rl.Color {
	switch block {
	case BlockAir:
		return rl.Blank
	case BlockDirt:
		return rl.Brown
	case BlockStone:
		return rl.Gray
	case BlockGrass:
		return rl.Green
	default:
		return rl.White
	}
}

type Chunk struct {
	Width, Height, Depth int
	blocks               []Block
}

func NewChunk(width, height, depth int) *Chunk {
	c := &Chunk{
		Width:  width,
		Height: height,
		Depth:  depth,
		blocks: make([]Block, width*height*depth),
	}

	for x := range width {
		for y := range height {
			for z := range depth {
				block := BlockAir
				if y == 0 {
					block = BlockStone
				} else if y > 0 && y <= 20 {
					block = BlockDirt
				} else if y == 21 {
					block = BlockGrass
				} else {
					block = BlockAir
				}
				c.SetBlock(x, y, z, block)
			}
		}
	}

	return c
}

func (c *Chunk) idx(x, y, z int) int {
	return x + c.Width*(y+c.Height*z)
}

func (c *Chunk) SetBlock(x, y, z int, block Block) int {
	if x < 0 || x >= c.Width || y < 0 || y >= c.Height || z < 0 || z >= c.Depth {
		panic("Index out of bounds")
	}

	idx := c.idx(x, y, z)
	c.blocks[idx] = block
	return idx
}

func (c *Chunk) Draw() {
	for x := range c.Width {
		for y := range c.Height {
			for z := range c.Depth {
				block := c.blocks[c.idx(x, y, z)]
				if block != BlockAir {
					pos := rl.NewVector3(float32(x), float32(y), float32(z))
					size := rl.NewVector3(1, 1, 1)
					color := ColorFromBlock(block)

					rl.DrawCubeV(pos, size, color)
					rl.DrawCubeWiresV(pos, size, rl.Black)
				}
			}
		}
	}
}

func main() {
	screenWidth, screenHeight := int32(800), int32(600)
	rl.InitWindow(screenWidth, screenHeight, "Minecraft Clone")

	camera := rl.NewCamera3D(
		rl.NewVector3(32, 24, 32), // position
		rl.NewVector3(0, 0, 0),    // target
		rl.NewVector3(0, 1, 0),    // up
		60,                        // fov
		rl.CameraPerspective,
	)

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	chunk := NewChunk(32, 32, 32)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFirstPerson)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		chunk.Draw()

		rl.EndMode3D()

		rl.DrawText("Move with WASD", 10, 10, 20, rl.DarkGray)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
