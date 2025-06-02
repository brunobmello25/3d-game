package world

const (
	CHUNK_SIZE = 16
)

type World struct {
	chunks []*Chunk
}

func NewWorld() *World {
	world := &World{}
	world.buildChunks()
	return world
}

func (w *World) Draw() {
	for _, chunk := range w.chunks {
		chunk.Draw()
	}
}

func (w *World) buildChunks() {
	width, depth := 1, 1
	w.chunks = make([]*Chunk, 0, width*depth)

	for x := range width {
		for z := range depth {
			chunk := NewChunk(x, 0, z)
			w.chunks = append(w.chunks, chunk)
		}
	}
}
