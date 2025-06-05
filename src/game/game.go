package game

import (
	"github.com/brunobmello25/3d-game/src/chunk"
	"github.com/brunobmello25/3d-game/src/player"
	texture "github.com/brunobmello25/3d-game/src/texture"
	"github.com/brunobmello25/3d-game/src/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ui               *ui.UI
	player           *player.Player
	screenDimensions rl.Vector2
	testChunks       []*chunk.Chunk
}

func NewGame() *Game {
	screenDimensions := rl.NewVector2(1200, 675)

	rl.InitWindow(int32(screenDimensions.X), int32(screenDimensions.Y), "MC Clone")

	texture.Init()
	texture.InitAtlas() // Initialize the texture atlas

	return &Game{
		player:           player.NewPlayer(),
		ui:               ui.NewUI(),
		screenDimensions: screenDimensions,
		testChunks: []*chunk.Chunk{
			chunk.NewChunk(rl.NewVector3(0, 0, 0)),
			chunk.NewChunk(rl.NewVector3(1, 0, 0)),
			chunk.NewChunk(rl.NewVector3(0, 0, 1)),
			chunk.NewChunk(rl.NewVector3(1, 0, 1)),
		},
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
	for _, c := range g.testChunks {
		c.Update()
	}
}

func (g *Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.Camera)

	for _, c := range g.testChunks {
		c.Render()
	}

	rl.EndMode3D()

	g.ui.DrawFPS(int32(g.screenDimensions.X))
	g.ui.DrawPosition(g.player.Position)

	rl.EndDrawing()
}
