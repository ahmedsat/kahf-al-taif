package main

import (
	"fmt"

	"github.com/ahmedsat/noor"
	"github.com/ahmedsat/silah/client"
)

var (
	vertexShaderSource = `
	#version 330
	in vec3 vp;
	void main() {
		gl_Position = vec4(vp, 1.0);
	}
	` + "\x00"

	fragmentShaderSource = `
	#version 330
	out vec4 frag_colour;
	void main() {
		frag_colour = vec4(1, 1, 1, 1);
	}
	` + "\x00"
)

func startClient(url string) (err error) {

	c, err := client.NewClient(url)
	if err != nil {
		return err
	}

	go func() {

		go c.ReceiveMessages()
		ch := c.GetIncomingChannel()
		for message := range ch {
			fmt.Println(message)
		}
	}()

	noor.Init(noor.Options{
		Title: "Kahf Al Taif",
	})

	shader, err := noor.CreateShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		return err
	}

	scene := noor.NewScene()
	scene.AddObject(noor.NewObject(
		[]noor.Vertex{
			{Position: [3]float32{-.5, -.5, 0}},
			{Position: [3]float32{.5, -.5, 0}},
			{Position: [3]float32{0, .5, 0}},
		},
		[]uint32{},
		&noor.Material{Shader: shader, Textures: []uint32{}},
	))

	noor.Run(func() {

		scene.Draw()

	})

	return c.Close()
}
