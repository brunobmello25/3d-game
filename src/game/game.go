package game

import (
	"github.com/brunobmello25/3d-game/src/player"
	"github.com/brunobmello25/3d-game/src/render"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	renderer *render.Renderer
	player   *player.Player
}

func NewGame() *Game {
	return &Game{
		player: player.NewPlayer(),
		// world:    world.NewWorld(),
		renderer: render.NewRenderer(),
	}
}

func (g *Game) Run() error {
	screenWidth, screenHeight := int32(1200), int32(675)
	rl.InitWindow(screenWidth, screenHeight, "MC Clone")

	g.renderer.InitTextures()
	// TODO: init player
	// TODO: init world
	// TODO: set world renderer

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Render(screenWidth, screenHeight)
	}

	g.renderer.Cleanup()
	rl.CloseWindow()
	return nil
}

func (g *Game) Update() {
	g.player.Update()
}

func (g *Game) Render(screenWidth, screenHeight int32) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.Camera)
	// TODO: world render
	rl.EndMode3D()

	g.renderer.DrawFPS(screenWidth)
	g.renderer.DrawPosition(g.player.Position)

	rl.EndDrawing()
}
