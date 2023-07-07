package app

import (
	"github.com/aquilax/go-perlin"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Chunk struct {
	blocks [CHUNK_SIZE][200][CHUNK_SIZE]int
}

const (
	SQUARE_SIZE = 20
	CHUNK_SIZE  = 16
)

const (
	AIR = iota
	GRASS_BLOCK
	DIRT_BLOCK
)

const (
	alpha       = 2
	beta        = 4
	n           = 5
	seed  int64 = 100
)

func generateChunk() *Chunk {
	//rand.Seed(0)
	c := new(Chunk)
	c.blocks = [CHUNK_SIZE][200][CHUNK_SIZE]int{}
	perlinGen := perlin.NewPerlin(alpha, beta, n, seed)

	for x := 0; x < CHUNK_SIZE; x++ {
		for z := 0; z < CHUNK_SIZE; z++ {
			height := int(0.5*(perlinGen.Noise2D(float64(x)/CHUNK_SIZE, float64(z)/CHUNK_SIZE)+1)*15) + 5
			c.blocks[x][height][z] = GRASS_BLOCK
			for height > 0 {
				height--
				c.blocks[x][height][z] = DIRT_BLOCK

			}
		}
	}

	return c
}

func (c *Chunk) displayChunk() {

	for x := 0; x < len(c.blocks); x++ {
		for y := 0; y < len(c.blocks[x]); y++ {
			for z := 0; z < len(c.blocks[x][y]); z++ {
				if c.blocks[x][y][z] != AIR {

					faces := faceDrawn{true, true, true, true, true, true}

					// left
					if x == 0 {
						faces.Left = true
					} else if c.blocks[x-1][y][z] != AIR {
						faces.Left = false
					}
					// right
					if x == len(c.blocks)-1 {
						faces.Right = true
					} else if c.blocks[x+1][y][z] != AIR {
						faces.Right = false
					}
					// Bottom
					if y == 0 {
						faces.Bottom = true
					} else if c.blocks[x][y-1][z] != AIR {
						faces.Bottom = false
					}
					// Top
					if y == len(c.blocks[x])-1 {
						faces.Top = true
					} else if c.blocks[x][y+1][z] != AIR {
						faces.Top = false
					}
					// Front
					if z == 0 {
						faces.Back = true
					} else if c.blocks[x][y][z-1] != AIR {
						faces.Back = false
					}
					// Back
					if z == len(c.blocks[x][y])-1 {
						faces.Front = true
					} else if c.blocks[x][y][z+1] != AIR {
						faces.Front = false
					}
					switch c.blocks[x][y][z] {
					case GRASS_BLOCK:
						DrawCubeTexture(bottomTex, sideTex, topTex, faces, rl.NewVector3(float32(x), float32(y), float32(z)), 1, 1, 1, rl.White)
					case DIRT_BLOCK:
						DrawCubeTexture(bottomTex, bottomTex, bottomTex, faces, rl.NewVector3(float32(x), float32(y), float32(z)), 1, 1, 1, rl.White)
					}

				}

			}
		}
	}
}
