package world

import (
	"fmt"

	"github.com/brunobmello25/3d-game/src/chunk"
	"github.com/brunobmello25/3d-game/src/player"
)

const RENDER_DISTANCE = 4 // Distance in chunks to render

type World struct {
	ChunkColumns map[string]*chunk.ChunkColumn
}

func NewWorld() *World {
	return &World{
		ChunkColumns: make(map[string]*chunk.ChunkColumn),
	}
}

func (w *World) Update(player player.Player) {
	w.loadAroundPlayer(int(player.Position.X), int(player.Position.Z))

	for _, column := range w.ChunkColumns {
		column.Update()
	}
}

func (w *World) Render() {
	for _, column := range w.ChunkColumns {
		column.Render()
	}
}

func (w *World) loadAroundPlayer(px, pz int) {
	minx := (px / chunk.CHUNK_SIZE) - RENDER_DISTANCE
	minz := (pz / chunk.CHUNK_SIZE) - RENDER_DISTANCE
	maxx := (px / chunk.CHUNK_SIZE) + RENDER_DISTANCE
	maxz := (pz / chunk.CHUNK_SIZE) + RENDER_DISTANCE

	// unload unecessary chunks
	// TODO: this breaks the game
	// for key := range w.ChunkColumns {
	// 	x, z := keyToPos(key)
	// 	distance := (x-px)*(x-px) + (z-pz)*(z-pz)
	// 	if distance > RENDER_DISTANCE*RENDER_DISTANCE {
	// 		w.ChunkColumns[key].Unload()
	// 		w.ChunkColumns[key] = nil
	// 	}
	// }

	fmt.Println("min x:", minx, "minz:", minz, "maxx:", maxx, "maxz:", maxz)
	// load proper chunks
	for x := minx; x <= maxx; x++ {
		for z := minz; z <= maxz; z++ {
			key := posToKey(x, z)
			if _, exists := w.ChunkColumns[key]; !exists {
				w.ChunkColumns[key] = chunk.NewChunkColumn(x, z)
				w.ChunkColumns[key].Generate()
			}
		}
	}
}

func posToKey(x, z int) string {
	return fmt.Sprintf("%d_%d", x, z)
}

func keyToPos(key string) (int, int) {
	var x, z int
	_, err := fmt.Sscanf(key, "%d_%d", &x, &z)
	if err != nil {
		panic(fmt.Sprintf("Invalid key format: %s", key))
	}
	return x, z
}
