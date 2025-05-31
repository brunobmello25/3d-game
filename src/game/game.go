package game

import (
	"github.com/brunobmello25/3d-game/src/player"
	"github.com/brunobmello25/3d-game/src/render"
	"github.com/brunobmello25/3d-game/src/world"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	player   *player.Player
	world    *world.World
	renderer *render.Renderer
}

func NewGame() *Game {
	return &Game{
		player:   player.NewPlayer(),
		world:    world.NewWorld(),
		renderer: render.NewRenderer(),
	}
}

func (g *Game) Run() {
	screenWidth, screenHeight := int32(1200), int32(675)
	rl.InitWindow(screenWidth, screenHeight, "MC Clone")

	g.renderer.Init()
	g.player.Init()
	g.world.Init()
	g.world.SetRenderer(g.renderer)

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Render(screenWidth, screenHeight)
	}

	g.renderer.Cleanup()
	rl.CloseWindow()
}

func (g *Game) Update() {
	g.player.Update()
}

func (g *Game) Render(screenWidth, screenHeight int32) {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.GetCamera())
	g.world.Render()
	rl.EndMode3D()

	g.renderer.DrawFPS(screenWidth)
	g.renderer.DrawPosition(g.player.GetPosition())

	rl.EndDrawing()
}
