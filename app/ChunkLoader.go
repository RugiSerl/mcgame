package app

import (
	gr "github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChunkLoader struct {
	chunks         map[gr.Vector2i]*Chunk
	RenderDistance float32
}

func NewChunkLoader() *ChunkLoader {
	c := new(ChunkLoader)

	c.RenderDistance = 2 // default value

	c.chunks = make(map[gr.Vector2i]*Chunk)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			c.chunks[gr.NewVector2i(x, y)] = generateChunk(gr.NewVector2i(x, y))

		}
	}

	return c

}

func (c *ChunkLoader) updateChunck(camera rl.Camera3D) {
	// add chunks
	var pos gr.Vector2i = gr.Vector2i{}
	for x := camera.Position.X/CHUNK_SIZE - c.RenderDistance; x <= camera.Position.X/CHUNK_SIZE+c.RenderDistance; x++ {
		for y := camera.Position.Z/CHUNK_SIZE - c.RenderDistance; y <= camera.Position.Z/CHUNK_SIZE+c.RenderDistance; y++ {
			pos.X, pos.Y = int(x), int(y)
			if _, isPresent := c.chunks[pos]; !isPresent {
				c.chunks[pos] = generateChunk(pos)
			}

		}
	}

	// remove chunks
	var distance float32
	for position, chunk := range c.chunks {

		distance = gr.NewVector2(camera.Position.X, camera.Position.Z).Substract(position.ToVector2().Scale(CHUNK_SIZE)).GetNorm()
		if distance > CHUNK_SIZE*c.RenderDistance {
			delete(c.chunks, position)
		}

		chunk.displayChunk(position)
	}

	c.renderChunks()
}

func (c *ChunkLoader) renderChunks() {
	for position, chunk := range c.chunks {
		chunk.displayChunk(position)
	}
}
