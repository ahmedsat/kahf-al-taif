package main

import (
	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Kahf Al Taif"
)

func startClient(url string) (err error) {

	if err := noor.Init(noor.Options{
		Width:      windowWidth,
		Height:     windowHeight,
		Title:      windowTitle,
		Background: [3]float32{0.1, 0.1, 0.1},
	}); err != nil {
		return err
	}

	// Create shader program
	shader, err := noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
	if err != nil {
		return err
	}

	// Load textures
	textures, err := loadTextures()
	if err != nil {
		return err
	}

	// Create cube mesh
	mesh, err := createCubeMesh()
	if err != nil {
		return err
	}

	// Create material
	material := &noor.Material{
		Shader:   shader,
		Textures: textures,
	}

	// Create cube object
	cube, err := noor.NewObject(mesh, material, madar.Identity())
	if err != nil {
		return err
	}

	// Create camera
	camera := noor.NewCamera(
		madar.Vector3{-3, -3, -3},
		madar.Vector3{},
		madar.Vector3{0, 1, 0},
		windowWidth/windowHeight,
		noor.Perspective)

	// Create scene
	scene := noor.NewScene(camera, cube)

	// Main loop
	return noor.Run(func() {
		updateScene(scene, cube)
		scene.Draw()
	})
}

func loadTextures() ([]noor.Texture, error) {
	textureFiles := []string{"textures/stone.webp", "textures/wall.jpg", "textures/StudentNTP_Aurora-Tennant_x1140.jpg"}
	textures := make([]noor.Texture, len(textureFiles))

	for i, file := range textureFiles {
		texture, err := noor.NewTextureFromFile(file, file)
		if err != nil {
			return nil, err
		}
		textures[i] = texture
	}

	return textures, nil
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

func updateScene(scene *noor.Scene, cube *noor.Object) {
	// Update cube rotation
	cube.UpdateModelMatrix(func() {
		// cube.ModelMatrix = cube.ModelMatrix.
		// Multiply(madar.RotationX(0.01)).
		// Multiply(madar.RotationY(0.01)).
		// Multiply(madar.RotationZ(0.01))

	})

	scene.UpdateCamera()
}
