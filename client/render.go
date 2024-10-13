package client

import (
	"errors"
	"math"
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

var (
	sh   noor.Shader
	wall noor.Texture
	cam  *noor.Camera
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

	sphereVertices, sphereIndices := generateSphereVertices(30, 30, 1)

	cam = noor.NewCamera(windowWidth, windowHeight)
	cam.SetPosition(0, 0, 5)
	cam.SetTarget(madar.Vector3{X: 0, Y: 0, Z: 0})
	cam.SetMode(noor.Orbit)
	cam.SetOrbitSpeed(0.1, 0.1, 0.1)
	cam.SetMovementSpeed(5.0)
	cam.SetMouseSensitivity(0.002)

	scene := noor.Scene{
		Objects: []*noor.Object{noor.NewObject(sphereVertices, sphereIndices, sh, wall)},
		Camera:  *cam,
	}

	lastTime := time.Now()

	noor.Run(
		func() {
			scene.Draw()
		},
		func(dt float32) {
			currentTime := time.Now()
			deltaTime := float32(currentTime.Sub(lastTime).Seconds())
			lastTime = currentTime

			cam.HandleInput(deltaTime)

			// Update camera in the scene
			scene.Camera = *cam

			// Check for mode change
			if input.IsKeyPressed(input.Key1) {
				cam.SetMode(noor.Free)
			} else if input.IsKeyPressed(input.Key2) {
				cam.SetMode(noor.Orbit)
			}

			// Handle zooming
			scroll := input.GetMouseScroll()
			cam.ZoomIn(scroll.Y * 0.1)
		},
		time.Second/60,
	)

	return
}

func generateSphereVertices(latitudeBands, longitudeBands int, radius float32) ([]noor.Vertex, []uint32) {
	var vertices []noor.Vertex
	var indices []uint32

	// Generate vertices
	for lat := 0; lat <= latitudeBands; lat++ {
		theta := float64(lat) * math.Pi / float64(latitudeBands)
		sinTheta := float32(math.Sin(theta))
		cosTheta := float32(math.Cos(theta))

		for lon := 0; lon <= longitudeBands; lon++ {
			phi := float64(lon) * 2.0 * math.Pi / float64(longitudeBands)
			sinPhi := float32(math.Sin(phi))
			cosPhi := float32(math.Cos(phi))

			x := cosPhi * sinTheta
			y := cosTheta
			z := sinPhi * sinTheta
			u := float32(lon) / float32(longitudeBands)
			v := float32(lat) / float32(latitudeBands)

			vertices = append(vertices, noor.Vertex{
				X: radius * x, Y: radius * y, Z: radius * z, W: 1.0,
				R: u, G: v, B: 1.0 - u, A: 1.0, // Color
				U: u, V: v, // Texture coordinates
			})
		}
	}

	// Generate indices
	for lat := 0; lat < latitudeBands; lat++ {
		for lon := 0; lon < longitudeBands; lon++ {
			first := uint32(lat*(longitudeBands+1) + lon)
			second := uint32((lat+1)*(longitudeBands+1) + lon)

			indices = append(indices, first, second, first+1)
			indices = append(indices, second, second+1, first+1)
		}
	}

	return vertices, indices
}
