package game

import (
	"github.com/brunobmello25/3d-game/src/player"
	texture "github.com/brunobmello25/3d-game/src/texture"
	"github.com/brunobmello25/3d-game/src/ui"
	"github.com/brunobmello25/3d-game/src/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ui               *ui.UI
	player           *player.Player
	screenDimensions rl.Vector2
	world            *world.World
}

func NewGame() *Game {
	screenDimensions := rl.NewVector2(1200, 675)

	rl.InitWindow(int32(screenDimensions.X), int32(screenDimensions.Y), "MC Clone")

	texture.Init()
	texture.InitAtlas()

	w := world.NewWorld()

	return &Game{
		player:           player.NewPlayer(),
		ui:               ui.NewUI(),
		screenDimensions: screenDimensions,
		world:            w,
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
	g.world.Update(*g.player)
}

func (g *Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.Camera)

	g.world.Render()

	rl.EndMode3D()

	g.ui.DrawFPS(int32(g.screenDimensions.X))
	g.ui.DrawPosition(g.player.Position)

	rl.EndDrawing()
}
