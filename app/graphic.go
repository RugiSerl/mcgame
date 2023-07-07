package app

import rl "github.com/gen2brain/raylib-go/raylib"

type faceDrawn struct {
	Top, Bottom, Left, Right, Front, Back bool
}

func DrawCubeTexture(bottomTex rl.Texture2D, sideTex rl.Texture2D, topTex rl.Texture2D, faces faceDrawn, position rl.Vector3, width float32, height float32, length float32, color rl.Color) {
	x := position.X
	y := position.Y
	z := position.Z

	rl.Begin(rl.RL_QUADS)
	rl.Color4ub(color.R, color.G, color.B, color.A)

	// Set desired texture to be enabled while drawing following vertex data
	rl.SetTexture(sideTex.ID)
	if faces.Front {
		// Front Face
		rl.Normal3f(0.0, 0.0, 1.0) // Normal Pointing Towards Viewer
		rl.TexCoord2f(0.0, 1)
		rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1.0, 1)
		rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1.0, 0)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Left Of The Texture and Quad
	}

	if faces.Back {
		// Back Face
		rl.Normal3f(0, 0, -1) // Normal Pointing Away From Viewer
		rl.TexCoord2f(1, 1)
		rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1, 0)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0, 1)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Left Of The Texture and Quad
	}
	if faces.Right {
		// Right face
		rl.Normal3f(1, 0, 0) // Normal Pointing Right
		rl.TexCoord2f(1, 1)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1, 0)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0, 1)
		rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Left Of The Texture and Quad
	}

	if faces.Left {
		// Left Face
		rl.Normal3f(-1, 0, 0) // Normal Pointing Left
		rl.TexCoord2f(0, 1)
		rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1, 1)
		rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1, 0)
		rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left Of The Texture and Quad
		rl.End()
	}

	if faces.Top {
		rl.SetTexture(topTex.ID)
		// Top Face
		rl.Normal3f(0, 1, 0) // Normal Pointing Up
		rl.TexCoord2f(0, 1)
		rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1, 0)
		rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Bottom Right Of The Texture and Quad
		rl.TexCoord2f(1, 1)
		rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Right Of The Texture and Quad
	}

	if faces.Bottom {
		rl.SetTexture(bottomTex.ID)
		// Bottom Face
		rl.Normal3f(0, -1, 0) // Normal Pointing Down
		rl.TexCoord2f(1, 1)
		rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Top Right Of The Texture and Quad
		rl.TexCoord2f(0, 1)
		rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Top Left Of The Texture and Quad
		rl.TexCoord2f(0, 0)
		rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Left Of The Texture and Quad
		rl.TexCoord2f(1, 0)
		rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Right Of The Texture and Quad
	}

	rl.SetTexture(0)

}
