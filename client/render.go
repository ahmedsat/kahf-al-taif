package client

// todo:

import (
	"image/color"
	"time"

	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
	"github.com/ahmedsat/noor/input"
)

const (
	windowTitle = "Kahf Al Taif"
)

var (
	diffuseMap  *noor.Texture
	specularMap *noor.Texture
)

func Init() (err error) {
	err = noor.Init(noor.InitSettings{
		WindowTitle:         windowTitle,
		BackGround:          colorFromInt(0x0d3642ff),
		GLCoreProfile:       true,
		EnableMultiSampling: true,
		VSyncEnabled:        true,
	})
	if err != nil {
		return err
	}

	return
}

func LoadTextures() (err error) {
	// Load diffuse map
	diffuseMap = noor.DefaultDiffuseTextureMap()
	diffuseMap.Name = "uMaterial.diffuseMap"

	// Load specular map
	specularMap = noor.DefaultSpecularTextureMap()
	specularMap.Name = "uMaterial.specularMap"

	return nil
}

func cleanup() {
	diffuseMap.Delete()
	specularMap.Delete()
	noor.Terminate()
}

func render() (err error) {

	err = Init()
	if err != nil {
		return err
	}

	defer cleanup()

	mesh, err := noor.LoadMesh("objects/car.obj", &noor.LoadMeshOptions{
		FlipUVs:     false,
		CalcNormals: true,
		Scale:       0.05,
	})
	if err != nil {
		return err
	}

	// mesh := meshes.CreateCube()

	sh, err := noor.CreateShaderProgramFromFiles("shaders/test.vert", "shaders/test.frag")
	if err != nil {
		return err
	}
	defer sh.Delete()

	err = LoadTextures()
	if err != nil {
		return err
	}

	sh.Activate()

	obj := noor.CreateObject(noor.ObjectCreateInfo{
		Mesh:     mesh,
		Shader:   sh,
		Textures: []noor.Texture{ /* wallTexture */ },
	})

	projection := &noor.Perspective{
		Fov:    100,
		Aspect: 4 / 3,
		Near:   0.1,
		Far:    10,
	}

	camera := noor.NewCamera(
		madar.Vector3{Z: 2.0},
		madar.Vector3{X: -1, Y: -1, Z: -1},
		madar.Vector3{X: 0, Y: 1, Z: 0},
		projection,
	)
	defer camera.Cleanup()
	camera.LookAt(madar.Vector3{})

	sh.SetUniform3f("uMainLight.position", madar.Vector3{Z: 2.0})
	sh.SetUniform3f("uMainLight.color", madar.Vector3{X: 1.0, Y: 0.95, Z: 0.85})
	sh.SetUniform1f("uMainLight.ambient", 0.2)
	sh.SetUniform1f("uMainLight.diffuse", 0.8)
	sh.SetUniform1f("uMainLight.specular", 0.6)

	diffuseMap.Activate(sh, 0, diffuseMap.Name)
	specularMap.Activate(sh, 1, diffuseMap.Name)
	sh.SetUniform1f("uMaterial.shininess", 32)

	// sh.SetUniform1b("uEnableFog", true)
	// sh.SetUniform3f("uFogColor", madar.Vector3{X: 0.7, Y: 0.7, Z: 0.8})
	// sh.SetUniform1f("uFogDensity", 0.5)

	// Initial timestamp
	lastTime := time.Now()
	for !noor.IsWindowShouldClose() {
		// Calculate delta time
		currentTime := time.Now()
		deltaTime := currentTime.Sub(lastTime).Seconds() // Convert to seconds
		lastTime = currentTime

		camera.Update(float32(deltaTime))

		obj.Rotate(madar.Vector3{
			X: 90,
			Y: 180,
			Z: 270,
		}.Scale(float32(deltaTime / 10)))

		obj.Draw(*camera)

		if input.IsKeyPressed(input.KeyW) {
			camera.MoveForward(float32(deltaTime))
		}

		if input.IsKeyPressed(input.KeyS) {
			camera.MoveBackward(float32(deltaTime))
		}

		if input.IsKeyPressed(input.KeyA) {
			camera.MoveLeft(float32(deltaTime))
		}

		if input.IsKeyPressed(input.KeyD) {
			camera.MoveRight(float32(deltaTime))
		}

		if input.IsKeyPressed(input.KeyQ) {
			camera.MoveUp(float32(deltaTime))
		}

		if input.IsKeyPressed(input.KeyE) {
			camera.MoveDown(float32(deltaTime))
		}

	}

	return nil
}

func colorFromInt(hex uint32) color.Color {

	return color.RGBA{
		R: uint8(hex >> 0x18),
		G: uint8(hex >> 0x10),
		B: uint8(hex >> 0x08),
		A: uint8(hex >> 0x00),
	}
}
