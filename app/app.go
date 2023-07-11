package app

import (
	"fmt"
	"unsafe"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	CAMERA_SPEED         = 10
	CAMERA_SCROLL_AMOUNT = 0.2
)

var (
	camera                     rl.Camera3D
	sideTex, bottomTex, topTex rl.Texture2D
	chunk                      *ChunkLoader
	mesh                       rl.Mesh
	mat                        rl.Material
	obj                        rl.Model
)

func Run() {
	load()
	for !rl.WindowShouldClose() {
		update()
	}
	quit()
}

func load() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.SetConfigFlags(rl.FlagFullscreenMode)
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	monitor := rl.GetCurrentMonitor()

	rl.InitWindow(int32(rl.GetMonitorWidth(monitor)), int32(rl.GetMonitorHeight(monitor)), "raylib [core] example - basic window")
	rl.DisableCursor()
	rl.SetTargetFPS(-1)

	bottomTex = rl.LoadTexture("res/bottom.png")
	topTex = rl.LoadTexture("res/top.png")
	sideTex = rl.LoadTexture("res/side.png")

	camera = rl.Camera3D{}
	camera.Position = rl.NewVector3(4.0, 10, 4.0)
	camera.Target = rl.NewVector3(0.0, 1.8, 0.0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective

	mesh = rl.GenMeshCube(1, 1, 1)

	vertices := unsafe.Pointer(mesh.Vertices)

	size := unsafe.Sizeof(vertices)

	for i := 0; i < int(mesh.VertexCount); i++ {
		fmt.Println(*(*float32)(unsafe.Pointer(uintptr(vertices) + size*uintptr(i))))

	}

	mat = rl.LoadMaterialDefault()
	obj = rl.LoadModelFromMesh(mesh)

	obj.Materials.Maps.Texture = bottomTex

	chunk = NewChunkLoader()
	chunk.RenderDistance = 1

}
func update() {

	updateInput()
	rl.BeginDrawing()

	rl.BeginMode3D(camera)
	rl.ClearBackground(rl.RayWhite)

	chunk.Update(camera)

	// rl.DrawModel(obj, rl.NewVector3(10, 20, 10), 1, rl.NewColor(0, 0, 0, 255))
	// rl.DrawModelWires(obj, rl.NewVector3(10, 20, 10), 1, rl.White)

	rl.EndMode3D()

	rl.DrawFPS(0, 0)
	rl.EndDrawing()

}

func updateInput() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		rl.TakeScreenshot("screenshot.png")

	}

	updateCamera()
}

func quit() {
	rl.UnloadTexture(bottomTex)
	rl.UnloadTexture(sideTex)
	rl.UnloadTexture(topTex)

	rl.CloseWindow()
}

func updateCamera() {
	var movement rl.Vector3 = rl.NewVector3(0, 0, 0)

	if rl.IsKeyDown(rl.KeyA) {
		movement.Y -= CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyD) {
		movement.Y += CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyW) {
		movement.X += CAMERA_SPEED * rl.GetFrameTime()
		if rl.IsKeyDown(rl.KeyLeftControl) {
			movement.X += CAMERA_SPEED * rl.GetFrameTime()

		}
	}
	if rl.IsKeyDown(rl.KeyS) {
		movement.X -= CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeySpace) {
		movement.Z += CAMERA_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		movement.Z -= CAMERA_SPEED * rl.GetFrameTime()
	}

	rotation := rl.Vector2Scale(rl.GetMouseDelta(), 0.1)

	rl.UpdateCameraPro(&camera, movement, rl.NewVector3(rotation.X, rotation.Y, 0), 0)
}
