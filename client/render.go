package client

// todo:

import (
	"image/color"
	"time"

	"github.com/ahmedsat/kahf-al-taif/utils"
	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
	"github.com/ahmedsat/noor/input"
	"github.com/ahmedsat/noor/meshes"
)

const (
	windowTitle = "Kahf Al Taif"
)

var (
	wallTexture noor.Texture
)

func Init() (err error) {
	err = noor.Init(noor.InitSettings{
		WindowTitle:   windowTitle,
		GLCoreProfile: true,
		DebugLines:    false,
		BackGround:    colorFromInt(0x0d3642ff),
	})
	if err != nil {
		return err
	}

	return
}

func LoadTextures() (err error) {
	wallImage, err := utils.LoadImages("textures/wall.jpg")
	if err != nil {
		return
	}

	wallTexture, err = noor.NewTexture(wallImage, "uWallTexture", noor.TextureParameters{
		WrappingS:    noor.Repeat,
		WrappingT:    noor.Repeat,
		BorderColor:  color.RGBA{R: 255, A: 255},
		FilteringMin: noor.LinearMipmapLinear,
		FilteringMag: noor.Linear,
		UseMipmaps:   true,
	})

	if err != nil {
		return
	}
	return
}

func cleanup() {
	wallTexture.Delete()
	noor.Terminate()
}

func render() (err error) {

	err = Init()
	if err != nil {
		return err
	}

	defer cleanup()

	mesh := meshes.CreateCube()

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
		madar.Vector3{Z: 2},
		madar.Vector3{X: -1, Y: -1, Z: -1},
		madar.Vector3{X: 0, Y: 1, Z: 0},
		projection,
	)
	defer camera.Cleanup()
	camera.LookAt(madar.Vector3{})

	sh.SetUniform1f("uAmbientStrength", 0.1)
	sh.SetUniform3f("uAmbientColor", 1, 1, 1)
	sh.SetUniform3f("uDiffuseLightPosition", 0, 0, 2)
	sh.SetUniform3f("uDiffuseLightColor", 1, 1, 1)

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
