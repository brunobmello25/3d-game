package game

import (
	"github.com/brunobmello25/3d-game/src/block"
	"github.com/brunobmello25/3d-game/src/mesh"
	"github.com/brunobmello25/3d-game/src/player"
	texture_manager "github.com/brunobmello25/3d-game/src/texture"
	"github.com/brunobmello25/3d-game/src/ui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ui               *ui.UI
	player           *player.Player
	screenDimensions rl.Vector2
	testMesh         rl.Mesh
	testModel        rl.Model
}

func NewGame() *Game {
	screenDimensions := rl.NewVector2(1200, 675)

	rl.InitWindow(int32(screenDimensions.X), int32(screenDimensions.Y), "MC Clone")

	texture_manager.Init()

	// Create test mesh
	dirtTexture := texture_manager.GetTexture(texture_manager.TEXTURE_NAME_DIRT) // Assuming you have a dirt texture
	meshBuilder := mesh.NewMeshBuilder()

	dirtBlockCenter := rl.NewVector3(0, 0, 0)

	dirtFace := block.NewFace(block.FacingDirectionFront, dirtTexture)

	meshBuilder.AddFace(dirtFace, dirtBlockCenter)
	testMesh := meshBuilder.Build()

	testModel := rl.LoadModelFromMesh(testMesh)
	rl.SetMaterialTexture(testModel.Materials, rl.MapDiffuse, dirtTexture)

	return &Game{
		player:           player.NewPlayer(),
		ui:               ui.NewUI(),
		screenDimensions: screenDimensions,
		testMesh:         testMesh,
		testModel:        testModel,
	}
}

func (g *Game) Run() error {
	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Render()
	}

	// Cleanup
	rl.UnloadModel(g.testModel)
	rl.UnloadMesh(&g.testMesh)

	rl.CloseWindow()
	return nil
}

func (g *Game) Update() {
	g.player.Update()
}

func (g *Game) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	rl.BeginMode3D(g.player.Camera)

	// Draw our test model
	rl.DrawModel(g.testModel, rl.NewVector3(0, 0, -5), 1, rl.White)

	rl.EndMode3D()

	g.ui.DrawFPS(int32(g.screenDimensions.X))
	g.ui.DrawPosition(g.player.Position)

	rl.EndDrawing()
}
