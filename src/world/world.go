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
	// Convert player position to chunk coordinates
	chunkX := px / chunk.CHUNK_SIZE
	chunkZ := pz / chunk.CHUNK_SIZE

	// Track which chunks should be kept
	keepChunks := make(map[string]bool)

	// Load chunks within render distance
	for x := chunkX - RENDER_DISTANCE; x <= chunkX+RENDER_DISTANCE; x++ {
		for z := chunkZ - RENDER_DISTANCE; z <= chunkZ+RENDER_DISTANCE; z++ {
			key := posToKey(x, z)
			keepChunks[key] = true

			// If chunk doesn't exist, create it
			if _, exists := w.ChunkColumns[key]; !exists {
				w.ChunkColumns[key] = chunk.NewChunkColumn(x, z)
				w.ChunkColumns[key].Generate()
			}
		}
	}

	// Unload chunks that are too far
	for key, column := range w.ChunkColumns {
		if !keepChunks[key] {
			column.Unload()
			delete(w.ChunkColumns, key)
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
