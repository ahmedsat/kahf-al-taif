package client

// todo:

import (
	"fmt"
	"image/color"
	"time"

	"github.com/ahmedsat/kahf-al-taif/utils"
	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
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

	camera := noor.NewCamera(noor.CreateCameraInfo{
		Position:   madar.Vector3{X: 0, Y: 0, Z: 2},
		Projection: noor.Perspective,
		Mode:       noor.Free,
		Width:      800,
		Height:     600,
	})
	defer camera.Cleanup()

	camera.SetControls(noor.CameraControls{
		BaseMovementSpeed: 10.0,
		MouseSensitivity:  1,
		FOVSpeed:          25,
		// ... other settings
	})

	// Initial timestamp
	lastTime := time.Now()
	for !noor.IsWindowShouldClose() {
		// Calculate delta time
		currentTime := time.Now()
		deltaTime := currentTime.Sub(lastTime).Seconds() // Convert to seconds
		lastTime = currentTime

		camera.ProcessInput(float32(deltaTime))
		camera.Update()

		obj.Draw(*camera)

		fmt.Printf("FOV: %f\r", camera.GetFOV())
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
