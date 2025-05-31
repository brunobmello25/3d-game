package world

const (
	CHUNK_SIZE   = 16
	CHUNK_HEIGHT = 256
)

type World struct {
	chunks []*Chunk
}
