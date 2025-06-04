package game

import (
	"github.com/brunobmello25/3d-game/src/block"
	"github.com/brunobmello25/3d-game/src/mesh"
	"github.com/brunobmello25/3d-game/src/player"
	texture "github.com/brunobmello25/3d-game/src/texture"
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

	texture.Init()
	texture.InitAtlas() // Initialize the texture atlas

	// Create test mesh
	meshBuilder := mesh.NewMeshBuilder()

	blockCenter := rl.NewVector3(0, 0, 0)

	frontFace := block.NewFace(block.FacingDirectionFront, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(frontFace, blockCenter)

	rightFace := block.NewFace(block.FacingDirectionRight, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(rightFace, blockCenter)

	leftFace := block.NewFace(block.FacingDirectionLeft, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(leftFace, blockCenter)

	topFace := block.NewFace(block.FacingDirectionUp, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(topFace, blockCenter)

	bottomFace := block.NewFace(block.FacingDirectionDown, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(bottomFace, blockCenter)

	backFace := block.NewFace(block.FacingDirectionBack, texture.TEXTURE_NAME_DIRT)
	meshBuilder.AddFace(backFace, blockCenter)

	testMesh := meshBuilder.Build()

	testModel := rl.LoadModelFromMesh(testMesh)
	rl.SetMaterialTexture(testModel.Materials, rl.MapDiffuse, texture.GetAtlasTexture())

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
