package game

import (
	"github.com/brunobmello25/3d-game/src/chunk"
	"github.com/brunobmello25/3d-game/src/player"
	texture "github.com/brunobmello25/3d-game/src/texture"
	"github.com/brunobmello25/3d-game/src/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const GRID_SIZE = 5 // 5x5 grid of chunk columns

type Game struct {
	ui               *ui.UI
	player           *player.Player
	screenDimensions rl.Vector2
	chunkGrid        [GRID_SIZE][GRID_SIZE]*chunk.ChunkColumn
}

func NewGame() *Game {
	screenDimensions := rl.NewVector2(1200, 675)

	rl.InitWindow(int32(screenDimensions.X), int32(screenDimensions.Y), "MC Clone")

	texture.Init()
	texture.InitAtlas() // Initialize the texture atlas

	// Create and initialize the 5x5 grid of chunk columns
	var chunkGrid [GRID_SIZE][GRID_SIZE]*chunk.ChunkColumn

	// Generate chunks centered around origin (offset by half the grid size)
	offset := GRID_SIZE / 2
	for x := 0; x < GRID_SIZE; x++ {
		for z := 0; z < GRID_SIZE; z++ {
			// Position chunks with offset so they're centered around 0,0
			chunkX := x - offset
			chunkZ := z - offset
			chunkGrid[x][z] = chunk.NewChunkColumn(chunkX, chunkZ)
			chunkGrid[x][z].Generate()
		}
	}

	return &Game{
		player:           player.NewPlayer(),
		ui:               ui.NewUI(),
		screenDimensions: screenDimensions,
		chunkGrid:        chunkGrid,
	}
}

func (g *Game) Run() error {
	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Render()
	}

	rl.CloseWindow()
	return nil
}

func (g *Game) Update() {
	g.player.Update()

	// Update all chunk columns in the grid
	for x := 0; x < GRID_SIZE; x++ {
		for z := 0; z < GRID_SIZE; z++ {
			if g.chunkGrid[x][z] != nil {
				g.chunkGrid[x][z].Update()
			}
		}
	}
}

func (g *Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.Camera)

	// Render all generated chunk columns in the grid
	for x := 0; x < GRID_SIZE; x++ {
		for z := 0; z < GRID_SIZE; z++ {
			column := g.chunkGrid[x][z]
			if column != nil && column.Generated {
				column.Render()
			}
		}
	}

	rl.EndMode3D()

	g.ui.DrawFPS(int32(g.screenDimensions.X))
	g.ui.DrawPosition(g.player.Position)

	rl.EndDrawing()
}
