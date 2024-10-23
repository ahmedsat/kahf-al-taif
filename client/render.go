package client

import (
	"math"
	"time"

	"github.com/ahmedsat/bayaan"
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
	earth  noor.Texture
	camera *noor.Camera
	scene  noor.Scene
	sunPos madar.Vector3
)

func render() (err error) {

	if err := noor.Init(windowWidth, windowHeight, windowTitle, false); err != nil {
		return bayaan.Error("Failed to initialize Noor: %s", err)
	}
	defer noor.Terminate()

	if err := loadResources(); err != nil {
		return bayaan.Error("Failed to load resources: %s", err)
	}

	setupCamera()

	sunPos = madar.Vector3{X: -2, Y: 0, Z: 2}

	earthObj := createObj()
	sunObj := createSun(sunPos, madar.Vector3{X: 0.1, Y: 0.1, Z: 0.1})
	scene = noor.Scene{
		Objects:         []*noor.Object{earthObj, sunObj},
		Camera:          camera,
		LightPos:        sunPos,
		AmbientColor:    madar.Vector3{X: 1.0, Y: 1.0, Z: 1.0},
		AmbientStrength: 0.1,
		CameraPos:       madar.Vector3{X: 0, Y: 0, Z: 9},
		LightColor:      madar.Vector3{X: 1, Y: 1, Z: 1},
	}

	earthObj.Rotation.Z = -23.4 * math.Pi / 180

	noor.Run(
		func() {
			scene.Draw()
		},
		func(dt float32) {
			handleInput(dt)
			earthObj.Rotation.Y += dt
		},
		time.Second/60,
	)

	return nil
}

func loadResources() error {

	earthImage, err := utils.LoadImages("textures/stone.webp")
	if err != nil {
		return bayaan.Error("Failed to load Earth texture: %s", err)
	}

	earth, err = noor.NewTexture(earthImage, "uTexture")
	if err != nil {
		return bayaan.Error("Failed to create Earth texture: %s", err)
	}

	return nil
}

func setupCamera() {

	camera = &noor.Camera{
		Position:         madar.Vector3{X: 0, Y: 0, Z: 4},
		Direction:        madar.Vector3{X: 0, Y: 0, Z: -1},
		Up:               madar.Vector3{X: 0, Y: 1, Z: 0},
		Projection:       noor.Perspective,
		Mode:             noor.Fixed,
		Zoom:             1,
		Width:            windowWidth,
		Height:           windowHeight,
		FOV:              45,
		Near:             0.1,
		Far:              100,
		MovementSpeed:    1,
		MouseSensitivity: 0.002,
		Damping:          0.9,
		Target:           madar.Vector3{X: 0, Y: 0, Z: 0},
		OrbitDistance:    5,
		OrbitSpeed:       madar.Vector3{X: 0.5, Y: 0.5, Z: 0.5},
	}
	camera.MoveForward(-1)

	camera.Update()
	input.LockMouse()
}

func handleInput(dt float32) {
	camera.HandleInput(dt)

	scroll := input.GetMouseScroll()
	camera.ZoomIn(scroll.Y * 0.1)

	if input.IsKeyPressed(input.Key0) {
		camera.SetMode(noor.Fixed)
	}

	if input.IsKeyPressed(input.Key1) {
		camera.SetMode(noor.Free)
	}

	if input.IsKeyPressed(input.Key2) {
		camera.SetMode(noor.ThirdPerson)
	}

	if input.IsKeyPressed(input.Key3) {
		camera.SetMode(noor.FirstPerson)
	}

	if input.IsKeyPressed(input.Key4) {
		camera.SetMode(noor.Orbit)
	}

}

func generateSphereVertices(latitudeBands, longitudeBands int, radius float32) ([]noor.Vertex, []uint32) {
	var vertices []noor.Vertex
	var indices []uint32

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
			u := 1 - float32(lon)/float32(longitudeBands)
			v := 1 - float32(lat)/float32(latitudeBands)

			vertices = append(vertices, noor.Vertex{
				Vx: radius * x, Vy: radius * y, Vz: radius * z,
				Nx: x, Ny: y, Nz: z,
				U: u, V: v,
			})
		}
	}

	for lat := 0; lat < latitudeBands; lat++ {
		for lon := 0; lon < longitudeBands; lon++ {
			first := uint32(lat*(longitudeBands+1) + lon)
			second := first + uint32(longitudeBands) + 1

			indices = append(indices, first, second, first+1, second, second+1, first+1)
		}
	}

	return vertices, indices
}

func createSun(pos, size madar.Vector3) *noor.Object {

	sh, err := noor.CreateShaderProgramFromFiles("shaders/sun.vert", "shaders/sun.frag")
	if err != nil {
		bayaan.Error("Failed to create sun shader program: %s", err)
		return nil
	}

	vertices, indices := generateSphereVertices(30, 30, 1.0)

	sun := noor.NewObject(vertices, indices, sh)
	sun.Position = pos
	sun.Scale = size
	return sun
}

func createObj() *noor.Object {

	sh, err := noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
	if err != nil {
		bayaan.Error("Failed to create base shader program: %s", err)
		return nil
	}

	vertices, indices := generateSphereVertices(30, 30, 1.0)

	obj := noor.NewObject(vertices, indices, sh, earth)
	return obj
}
