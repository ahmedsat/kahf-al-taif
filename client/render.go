package client

import (
	"errors"

	"github.com/ahmedsat/bayaan"
	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Kahf Al Taif"

	cubesCount = 500
)

var (
	textures []*noor.Texture
	cubesPos []madar.Vector3
)

func render() (err error) {

	noor.Init(noor.Options{
		Width:             windowWidth,
		Height:            windowHeight,
		Title:             windowTitle,
		IsResizable:       false,
		DefaultExtButtons: 0,
		Background:        [3]float32{},
	})

	textures, err = loadTextures()
	if err != nil {
		err = errors.Join(err, errors.New("failed to load textures"))
		return
	}

	scene := noor.NewScene(&noor.Camera{
		Position:       [3]float32{0, 0, 10},
		LookAt:         [3]float32{},
		Up:             [3]float32{0, 1, 0},
		ProjectionType: noor.Perspective,
		Near:           0.1,
		Far:            100,
		Fov:            90,
		Aspect:         float32(windowWidth) / float32(windowHeight),
	}, createObjects(cubesCount)...)

	for range cubesCount {
		cubesPos = append(cubesPos, madar.Vector3{
			madar.RandFloatRange(-5, 5),
			madar.RandFloatRange(-5, 5),
			madar.RandFloatRange(-5, 5),
		})
	}

	noor.Run(func() {
		updateScene(scene)
		scene.Draw()
	})
	return
}

func loadTextures() ([]*noor.Texture, error) {
	textureFiles := []string{
		"textures/stone.webp",
		"textures/wall.jpg",
		"textures/StudentNTP_Aurora-Tennant_x1140.jpg"}

	textures := make([]*noor.Texture, len(textureFiles))

	for i, file := range textureFiles {
		texture, err := noor.NewTextureFromFile(file, file)
		if err != nil {
			return nil, err
		}
		textures[i] = &texture
	}

	return textures, nil
}

func createObjects(count int) []*noor.Object {

	var cubes []*noor.Object
	for range count {

		mesh, err := createCubeMesh()
		if err != nil {
			bayaan.Fatal("%s", err)
		}

		shader, err := noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
		if err != nil {
			bayaan.Fatal("%s", err)
		}

		material := noor.Material{
			Shader:   &shader,
			Textures: textures,
		}

		cube, err := noor.NewObject(mesh, &material, madar.Identity())
		if err != nil {
			panic(err)
		}
		cubes = append(cubes, cube)
	}
	return cubes
}

func createCubeMesh() (*noor.Mesh, error) {
	vertices := []noor.Vertex{
		// Front face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}), // Bottom left
		noor.NewVertex(madar.Vector3{0.5, -0.5, 0.5}, madar.Vector3{0.0, 1.0, 0.0}, madar.Vector2{1.0, 0.0}),  // Bottom right
		noor.NewVertex(madar.Vector3{0.5, 0.5, 0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{1.0, 1.0}),   // Top right
		noor.NewVertex(madar.Vector3{0.5, 0.5, 0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{1.0, 1.0}),   // Top right (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{-0.5, 0.5, 0.5}, madar.Vector3{1.0, 1.0, 0.0}, madar.Vector2{0.0, 1.0}),  // Top left
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}), // Bottom left (duplicate for second triangle)

		// Back face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, -0.5}, madar.Vector3{1.0, 0.0, 1.0}, madar.Vector2{1.0, 0.0}), // Bottom left
		noor.NewVertex(madar.Vector3{0.5, -0.5, -0.5}, madar.Vector3{0.0, 1.0, 1.0}, madar.Vector2{0.0, 0.0}),  // Bottom right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{0.0, 1.0}),   // Top right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{0.0, 1.0}),   // Top right (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{-0.5, 0.5, -0.5}, madar.Vector3{0.5, 0.5, 0.5}, madar.Vector2{1.0, 1.0}),  // Top left
		noor.NewVertex(madar.Vector3{-0.5, -0.5, -0.5}, madar.Vector3{1.0, 0.0, 1.0}, madar.Vector2{1.0, 0.0}), // Bottom left (duplicate for second triangle)

		// Left face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom left
		noor.NewVertex(madar.Vector3{-0.5, 0.5, 0.5}, madar.Vector3{1.0, 1.0, 0.0}, madar.Vector2{0.0, 1.0}),   // Front top left
		noor.NewVertex(madar.Vector3{-0.5, 0.5, -0.5}, madar.Vector3{0.5, 0.5, 0.5}, madar.Vector2{1.0, 1.0}),  // Back top left
		noor.NewVertex(madar.Vector3{-0.5, 0.5, -0.5}, madar.Vector3{0.5, 0.5, 0.5}, madar.Vector2{1.0, 1.0}),  // Back top left (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, -0.5}, madar.Vector3{1.0, 0.0, 1.0}, madar.Vector2{1.0, 0.0}), // Back bottom left
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom left (duplicate for second triangle)

		// Right face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{0.5, -0.5, 0.5}, madar.Vector3{0.0, 1.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom right
		noor.NewVertex(madar.Vector3{0.5, 0.5, 0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{0.0, 1.0}),   // Front top right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{1.0, 1.0}),  // Back top right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{1.0, 1.0}),  // Back top right (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{0.5, -0.5, -0.5}, madar.Vector3{0.0, 1.0, 1.0}, madar.Vector2{1.0, 0.0}), // Back bottom right
		noor.NewVertex(madar.Vector3{0.5, -0.5, 0.5}, madar.Vector3{0.0, 1.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom right (duplicate for second triangle)

		// Top face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{-0.5, 0.5, 0.5}, madar.Vector3{1.0, 1.0, 0.0}, madar.Vector2{0.0, 1.0}),  // Front bottom left
		noor.NewVertex(madar.Vector3{0.5, 0.5, 0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{1.0, 1.0}),   // Front bottom right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{1.0, 0.0}),  // Back bottom right
		noor.NewVertex(madar.Vector3{0.5, 0.5, -0.5}, madar.Vector3{1.0, 1.0, 1.0}, madar.Vector2{1.0, 0.0}),  // Back bottom right (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{-0.5, 0.5, -0.5}, madar.Vector3{0.5, 0.5, 0.5}, madar.Vector2{0.0, 0.0}), // Back bottom left
		noor.NewVertex(madar.Vector3{-0.5, 0.5, 0.5}, madar.Vector3{1.0, 1.0, 0.0}, madar.Vector2{0.0, 1.0}),  // Front bottom left (duplicate for second triangle)

		// Bottom face (2 triangles, 6 vertices)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom left
		noor.NewVertex(madar.Vector3{0.5, -0.5, 0.5}, madar.Vector3{0.0, 1.0, 0.0}, madar.Vector2{1.0, 0.0}),   // Front bottom right
		noor.NewVertex(madar.Vector3{0.5, -0.5, -0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{1.0, 1.0}),  // Back bottom right
		noor.NewVertex(madar.Vector3{0.5, -0.5, -0.5}, madar.Vector3{0.0, 0.0, 1.0}, madar.Vector2{1.0, 1.0}),  // Back bottom right (duplicate for second triangle)
		noor.NewVertex(madar.Vector3{-0.5, -0.5, -0.5}, madar.Vector3{1.0, 1.0, 0.0}, madar.Vector2{0.0, 1.0}), // Back bottom left
		noor.NewVertex(madar.Vector3{-0.5, -0.5, 0.5}, madar.Vector3{1.0, 0.0, 0.0}, madar.Vector2{0.0, 0.0}),  // Front bottom left (duplicate for second triangle)
	}

	return noor.NewMesh(vertices, []uint32{})
}

func updateScene(scene *noor.Scene) {

	t := glfw.GetTime()

	for i := range scene.Objects {
		scene.Objects[i].Material.Shader.Activate()
		scene.Objects[i].UpdateModelMatrix(func() {
			scene.Objects[i].ModelMatrix = madar.Identity().
				Multiply(madar.RotationX(cubesPos[i][0] * float32(t))).
				Multiply(madar.RotationY(cubesPos[i][1] * float32(t))).
				Multiply(madar.RotationZ(cubesPos[i][2] * float32(t))).
				Multiply(madar.Translation(cubesPos[i][0], cubesPos[i][1], cubesPos[i][2]))
		})
	}

	// scene.UpdateCamera()
}
