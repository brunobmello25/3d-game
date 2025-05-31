package world

const (
	CHUNK_SIZE   = 16
	CHUNK_HEIGHT = 256
)

type World struct {
	chunks []*Chunk
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Init() {
	w.chunks = w.buildChunks()
}

func (w *World) Render() {
	for _, chunk := range w.chunks {
		chunk.Render()
	}
}

func (w *World) buildChunks() []*Chunk {
	width, depth := 1, 1
	chunks := []*Chunk{}

	for x := 0; x < width; x++ {
		for z := 0; z < depth; z++ {
			chunk := NewChunk(x, 0, z)
			chunks = append(chunks, chunk)
		}
	}

	return chunks
}
