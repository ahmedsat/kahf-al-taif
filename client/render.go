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

var vertices = []float32{
	// Front face
	-1.0, -1.0, 1.0, 0.0, 0.0, // Bottom-left
	1.0, -1.0, 1.0, 1.0, 0.0, // Bottom-right
	1.0, 1.0, 1.0, 1.0, 1.0, // Top-right
	-1.0, 1.0, 1.0, 0.0, 1.0, // Top-left
	// Back face
	-1.0, -1.0, -1.0, 1.0, 0.0, // Bottom-left
	1.0, -1.0, -1.0, 0.0, 0.0, // Bottom-right
	1.0, 1.0, -1.0, 0.0, 1.0, // Top-right
	-1.0, 1.0, -1.0, 1.0, 1.0, // Top-left

	// Left face
	-1.0, 1.0, 1.0, 1.0, 0.0, // Top-right
	-1.0, 1.0, -1.0, 1.0, 1.0, // Top-left
	-1.0, -1.0, -1.0, 0.0, 1.0, // Bottom-left
	-1.0, -1.0, 1.0, 0.0, 0.0, // Bottom-right
	// Right face
	1.0, 1.0, 1.0, 1.0, 0.0, // Top-left
	1.0, 1.0, -1.0, 1.0, 1.0, // Top-right
	1.0, -1.0, -1.0, 0.0, 1.0, // Bottom-right
	1.0, -1.0, 1.0, 0.0, 0.0, // Bottom-left
	// Top face
	-1.0, 1.0, -1.0, 0.0, 1.0, // Top-left
	1.0, 1.0, -1.0, 1.0, 1.0, // Top-right
	1.0, 1.0, 1.0, 1.0, 0.0, // Bottom-right
	-1.0, 1.0, 1.0, 0.0, 0.0, // Bottom-left
	// Bottom face
	-1.0, -1.0, -1.0, 0.0, 1.0, // Top-left
	1.0, -1.0, -1.0, 1.0, 1.0, // Top-right
	1.0, -1.0, 1.0, 1.0, 0.0, // Bottom-right
	-1.0, -1.0, 1.0, 0.0, 0.0, // Bottom-left
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

	wall, err = noor.NewTexture(wallImage, "wall")
	if err != nil {
		return errors.Join(err, errors.New("failed to create wall texture"))
	}

	scene := noor.Scene{
		Objects: createRandomCubes(50),
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
