package client

import (
	"image/color"
	"math"

	"github.com/ahmedsat/noor"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Kahf Al Taif"
)

func render() (err error) {

	err = noor.Init(noor.InitSettings{
		WindowWidth:     windowWidth,
		WindowHeight:    windowHeight,
		WindowTitle:     windowTitle,
		WindowResizable: false,
		GLMajorVersion:  4,
		GLMinorVersion:  5,
		GLCoreProfile:   true,
		DebugLines:      false,
		BackGround:      colorFromInt(0x0d3642ff),
	})
	if err != nil {
		return err
	}
	defer noor.Terminate()

	mesh := noor.CreateMesh(noor.CreateMeshInfo{
		Vertices: []float32{
			-0.5, -0.5, 0.18, 0.55, 0.80,
			0.5, -0.5, 0.18, 0.55, 0.80,
			0.0, 0.5, 0.18, 0.55, 0.80,
		},
		Sizes: []int32{
			2,
			3,
		},
	})

	sh, err := noor.CreateShaderProgramFromFiles("shaders/test.vert", "shaders/test.frag")
	if err != nil {
		return err
	}
	defer sh.Delete()

	for !noor.IsWindowShouldClose() {
		sh.Activate()

		sh.SetUniform1f("uScaler", float32(math.Sin(glfw.GetTime())+1))
		mesh.Draw()
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
