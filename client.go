package main

import (
	"embed"
	"errors"

	"github.com/ahmedsat/noor"
)

var (
	vertexShaderSource   = ``
	fragmentShaderSource = ``
	wallTexture          noor.Texture
	stoneTexture         noor.Texture
)

//go:embed shaders/* textures/*
var embedFiles embed.FS

func loadResources() (err error) {

	vertexShaderBytes, err := embedFiles.ReadFile("shaders/base.vert")
	if err != nil {
		return errors.Join(err, errors.New("failed to load vertex shader: "))
	}
	vertexShaderSource = string(vertexShaderBytes)

	fragmentShaderBytes, err := embedFiles.ReadFile("shaders/base.frag")
	if err != nil {
		return errors.Join(err, errors.New("failed to load fragment shader: "))
	}
	fragmentShaderSource = string(fragmentShaderBytes)

	wallTextureFile, err := embedFiles.Open("textures/wall.jpg")
	if err != nil {
		return errors.Join(err, errors.New("failed to load wall texture: "))
	}

	wallTexture, err = noor.LoadTexture(wallTextureFile, "wall")
	if err != nil {
		return errors.Join(err, errors.New("failed to load wall texture: "))
	}

	stoneTextureFile, err := embedFiles.Open("textures/stone.webp")
	if err != nil {
		return errors.Join(err, errors.New("failed to load stone texture: "))
	}

	stoneTexture, err = noor.LoadTexture(stoneTextureFile, "stone")
	if err != nil {
		return errors.Join(err, errors.New("failed to load stone texture: "))
	}

	return
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

	loadResources()

	shader, err := noor.CreateShaderProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		err = errors.Join(err, errors.New("failed to create shader program: "))
		return
	}

	scene := noor.NewScene()

	obj, err := noor.NewObject(
		[]noor.Vertex{
			{Position: [3]float32{0.5, 0.5, 0.0}, Color: [3]float32{1.0, 0.0, 0.0}, TexCoord: [2]float32{1.0, 1.0}},
			{Position: [3]float32{0.5, -0.5, 0.0}, Color: [3]float32{0.0, 1.0, 0.0}, TexCoord: [2]float32{1.0, 0.0}},
			{Position: [3]float32{-0.5, -0.5, 0.0}, Color: [3]float32{0.0, 0.0, 1.0}, TexCoord: [2]float32{0.0, 0.0}},
			{Position: [3]float32{-0.5, 0.5, 0.0}, Color: [3]float32{1.0, 1.0, 0.0}, TexCoord: [2]float32{0.0, 1.0}},
		},
		[]uint32{0, 1, 3, 1, 2, 3},
		&noor.Material{Shader: shader, Textures: []noor.Texture{stoneTexture, wallTexture}},
	)
	if err != nil {
		err = errors.Join(err, errors.New("failed to create object: "))
		return
	}

	scene.AddObject(obj)

	err = noor.Run(func() {
		scene.Draw()
	})
	if err != nil {
		err = errors.Join(err, errors.New("failed to run: "))
		return
	}

	// return c.Close()
	return nil
}
