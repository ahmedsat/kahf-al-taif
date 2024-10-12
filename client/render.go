package client

import (
	"errors"
	"time"

	"github.com/ahmedsat/kahf-al-taif/utils"
	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
	"github.com/ahmedsat/noor/input"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Kahf Al Taif"
)

var vertices = []noor.Vertex{
	// Front face
	{X: -1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 1.0, G: 0.0, B: 0.0, A: 1.0, U: 0.0, V: 0.0}, // Bottom-left - Red
	{X: 1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 0.0, G: 1.0, B: 0.0, A: 1.0, U: 1.0, V: 0.0},  // Bottom-right - Green
	{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 0.0, G: 0.0, B: 1.0, A: 1.0, U: 1.0, V: 1.0},   // Top-right - Blue
	{X: -1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 1.0, G: 1.0, B: 0.0, A: 1.0, U: 0.0, V: 1.0},  // Top-left - Yellow
	// Back face
	{X: -1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 1.0, G: 0.0, B: 1.0, A: 1.0, U: 1.0, V: 0.0}, // Bottom-left - Magenta
	{X: 1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 0.0, G: 1.0, B: 1.0, A: 1.0, U: 0.0, V: 0.0},  // Bottom-right - Cyan
	{X: 1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 1.0, G: 0.5, B: 0.0, A: 1.0, U: 0.0, V: 1.0},   // Top-right - Orange
	{X: -1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 0.5, G: 0.5, B: 0.5, A: 1.0, U: 1.0, V: 1.0},  // Top-left - Gray

	// Left face
	{X: -1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 0.0, G: 0.0, B: 1.0, A: 1.0, U: 1.0, V: 0.0},   // Top-right - Blue
	{X: -1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 1.0, G: 0.5, B: 0.0, A: 1.0, U: 1.0, V: 1.0},  // Top-left - Orange
	{X: -1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 0.5, G: 0.5, B: 0.5, A: 1.0, U: 0.0, V: 1.0}, // Bottom-left - Gray
	{X: -1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 0.0, G: 1.0, B: 0.0, A: 1.0, U: 0.0, V: 0.0},  // Bottom-right - Green
	// Right face
	{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 1.0, G: 0.0, B: 0.0, A: 1.0, U: 1.0, V: 0.0},   // Top-left - Red
	{X: 1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 0.0, G: 1.0, B: 0.0, A: 1.0, U: 1.0, V: 1.0},  // Top-right - Green
	{X: 1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 0.0, G: 0.0, B: 1.0, A: 1.0, U: 0.0, V: 1.0}, // Bottom-right - Blue
	{X: 1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 1.0, G: 1.0, B: 0.0, A: 1.0, U: 0.0, V: 0.0},  // Bottom-left - Yellow

	// Top face
	{X: -1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 1.0, G: 0.0, B: 1.0, A: 1.0, U: 0.0, V: 1.0}, // Top-left - Magenta
	{X: 1.0, Y: 1.0, Z: -1.0, W: 1.0, R: 0.0, G: 1.0, B: 1.0, A: 1.0, U: 1.0, V: 1.0},  // Top-right - Cyan
	{X: 1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 1.0, G: 0.5, B: 0.0, A: 1.0, U: 1.0, V: 0.0},   // Bottom-right - Orange
	{X: -1.0, Y: 1.0, Z: 1.0, W: 1.0, R: 0.5, G: 0.5, B: 0.5, A: 1.0, U: 0.0, V: 0.0},  // Bottom-left - Gray
	// Bottom face
	{X: -1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 0.5, G: 0.5, B: 0.5, A: 1.0, U: 0.0, V: 1.0}, // Top-left - Gray
	{X: 1.0, Y: -1.0, Z: -1.0, W: 1.0, R: 1.0, G: 0.0, B: 1.0, A: 1.0, U: 1.0, V: 1.0},  // Top-right - Magenta
	{X: 1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 0.0, G: 1.0, B: 0.0, A: 1.0, U: 1.0, V: 0.0},   // Bottom-right - Green
	{X: -1.0, Y: -1.0, Z: 1.0, W: 1.0, R: 1.0, G: 1.0, B: 0.0, A: 1.0, U: 0.0, V: 0.0},  // Bottom-left - Yellow
}

var indices = []uint32{
	// Front face
	0, 1, 2, 2, 3, 0,
	// Back face
	4, 5, 6, 6, 7, 4,
	// Left face
	8, 9, 10, 10, 11, 8,
	// Right face
	12, 13, 14, 14, 15, 12,
	// Top face
	16, 17, 18, 18, 19, 16,
	// Bottom face
	20, 21, 22, 22, 23, 20,
}

var (
	sh   noor.Shader
	wall noor.Texture
)

func render() (err error) {
	noor.Init(windowWidth, windowHeight, windowTitle, false)
	defer noor.Terminate()

	input.LockMouse()

	sh, err = noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
	if err != nil {
		return errors.Join(err, errors.New("failed to create shader program"))
	}

	wallImage, err := utils.LoadImages("textures/wall.jpg")
	if err != nil {
		return errors.Join(err, errors.New("failed to load wall image"))
	}

	wall, err = noor.NewTexture(wallImage, "uTexture")
	if err != nil {
		return errors.Join(err, errors.New("failed to create wall texture"))
	}

	scene := noor.Scene{
		Objects: []*noor.Object{createCenterCube()},
		Camera:  *noor.NewCamera(windowWidth, windowHeight),
	}

	noor.Run(
		func() {
			scene.Draw()
		},
		func(dt float32) {

			if input.IsKeyHeld(input.KeyW) {
				scene.Camera.MoveForward(dt * 5)
			}
			if input.IsKeyHeld(input.KeyS) {
				scene.Camera.MoveForward(-dt * 5)
			}
			if input.IsKeyHeld(input.KeyA) {
				scene.Camera.MoveRight(-dt * 5)
			}
			if input.IsKeyHeld(input.KeyD) {
				scene.Camera.MoveRight(dt * 5)
			}

			if input.IsKeyHeld(input.KeyE) {
				scene.Camera.MoveUp(dt * 5)
			}
			if input.IsKeyHeld(input.KeyQ) {
				scene.Camera.MoveUp(-dt * 5)
			}

			// make camera look around with mouse
			mouseDelta := input.GetMouseDelta()
			scene.Camera.Rotate(mouseDelta.X*dt, mouseDelta.Y*dt, 0)
			scene.Camera.Rotate(0, 0, 0)

		},
		time.Second/60,
	)

	return
}

func createRandomCubes(n int) []*noor.Object {
	var cubes []*noor.Object

	rand := madar.NewRand(0x0000000000000000, 0xffffffffffffffff)

	for i := 0; i < n; i++ {
		cube := noor.NewObject(vertices, indices, sh, wall)

		cube.SetPosition(
			rand.RandFloatRange(-50, 50),
			rand.RandFloatRange(-50, 50),
			rand.RandFloatRange(-50, 50),
		)

		cube.SetScale(
			rand.RandFloatRange(1, 10),
			rand.RandFloatRange(1, 10),
			rand.RandFloatRange(1, 10),
		)

		cube.SetRotation(
			rand.RandFloatRange(0, 360),
			rand.RandFloatRange(0, 360),
			rand.RandFloatRange(0, 360),
		)

		cubes = append(cubes, cube)

	}
	return cubes
}

func createCenterCube() *noor.Object {
	cube := noor.NewObject(vertices, indices, sh, wall)

	cube.SetPosition(0, 0, 0)

	cube.SetScale(1, 1, 1)

	cube.SetRotation(0, 0, 0)

	return cube
}
