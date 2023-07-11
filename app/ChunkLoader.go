package app

import (
	gr "github.com/RugiSerl/rayutils/graphic"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ChunkLoader struct {
	chunks         map[gr.Vector2i]*Chunk
	RenderDistance float32
	RenderInterval float32

	lastChunkUpdate float32
}

func NewChunkLoader() *ChunkLoader {
	c := new(ChunkLoader)

	c.RenderDistance = 2 // default value
	c.RenderInterval = 0 // default value

	c.chunks = make(map[gr.Vector2i]*Chunk)
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			c.chunks[gr.NewVector2i(x, y)] = generateChunk(gr.NewVector2i(x, y))

		}
	}
	c.lastChunkUpdate = 0

	return c

}

func (c *ChunkLoader) Update(camera rl.Camera3D) {
	c.lastChunkUpdate += rl.GetFrameTime()

	if c.lastChunkUpdate > c.RenderInterval {
		c.updateChunck(camera)
		c.lastChunkUpdate = 0

	}
	c.renderChunks()

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
	for position := range c.chunks {

		distance = gr.NewVector2(camera.Position.X, camera.Position.Z).Substract(position.ToVector2().Scale(CHUNK_SIZE)).GetNorm()
		if distance > CHUNK_SIZE*c.RenderDistance {
			delete(c.chunks, position)
		}

	}

}

func (c *ChunkLoader) renderChunks() {
	for position, chunk := range c.chunks {
		chunk.getMesh(position)
	}
}
