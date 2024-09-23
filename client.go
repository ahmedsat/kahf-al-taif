package main

import (
	"embed"

	"github.com/ahmedsat/noor"
)

var (
	vertexShaderSource   = ``
	fragmentShaderSource = ``
)

//go:embed shaders/*
var shaders embed.FS

func init() {

	vertexShaderBytes, err := shaders.ReadFile("shaders/base.vert")
	if err != nil {
		panic(err)
	}
	vertexShaderSource = string(vertexShaderBytes)

	fragmentShaderBytes, err := shaders.ReadFile("shaders/base.frag")
	if err != nil {
		panic(err)
	}
	fragmentShaderSource = string(fragmentShaderBytes)
}

func startClient(url string) (err error) {

	// c, err := client.NewClient(url)
	// if err != nil {
	// 	return err
	// }

	// go func() {

	// 	go c.ReceiveMessages()
	// 	ch := c.GetIncomingChannel()
	// 	for message := range ch {
	// 		fmt.Println(message)
	// 	}
	// }()

	noor.Init(noor.Options{
		Title:      "Kahf Al Taif",
		Background: [3]float32{0.2, 0.3, 0.3},
	})

	shader, err := noor.CreateShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		return err
	}

	scene := noor.NewScene()
	scene.AddObject(noor.NewObject(
		[]noor.Vertex{
			{Position: [3]float32{-.5, -.5, 0}, Color: [3]float32{1.0, 0.5, 0.2}},
			{Position: [3]float32{.5, -.5, 0}, Color: [3]float32{1.0, 0.5, 0.2}},
			{Position: [3]float32{0, .5, 0}, Color: [3]float32{1.0, 0.5, 0.2}},
		},
		[]uint32{},
		&noor.Material{Shader: shader, Textures: []uint32{}},
	))

	noor.Run(func() {

		scene.Draw()

	})

	// return c.Close()
	return nil
}
