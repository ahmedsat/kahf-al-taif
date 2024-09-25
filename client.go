package main

import (
	"errors"

	"github.com/ahmedsat/madar"
	"github.com/ahmedsat/noor"
)

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

	shader, err := noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
	if err != nil {
		err = errors.Join(err, errors.New("failed to create shader program: "+"shaders/base.vert"+" and "+"shaders/base.frag"))
		return
	}

	stoneTexture, err := noor.NewTextureFromFile("textures/stone.webp", "stone")
	if err != nil {
		err = errors.Join(err, errors.New("failed to create stone texture: "))
		return
	}

	wallTexture, err := noor.NewTextureFromFile("textures/wall.jpg", "wall")
	if err != nil {
		err = errors.Join(err, errors.New("failed to create wall texture: "))
		return
	}

	tennantTexture, err := noor.NewTextureFromFile("textures/StudentNTP_Aurora-Tennant_x1140.jpg", "tennant")
	if err != nil {
		err = errors.Join(err, errors.New("failed to create tennant texture: "))
		return
	}

	scene := noor.NewScene()

	obj, err := noor.NewObject(
		[]noor.Vertex{
			noor.NewVertex(madar.NewVector(0.5, 0.5, 0.0), madar.NewVector(1.0, 0.0, 0.0), madar.NewVector(1.0, 1.0)),
			noor.NewVertex(madar.NewVector(0.5, -0.5, 0.0), madar.NewVector(0.0, 1.0, 0.0), madar.NewVector(1.0, 0.0)),
			noor.NewVertex(madar.NewVector(-0.5, -0.5, 0.0), madar.NewVector(0.0, 0.0, 1.0), madar.NewVector(0.0, 0.0)),
			noor.NewVertex(madar.NewVector(-0.5, 0.5, 0.0), madar.NewVector(1.0, 1.0, 0.0), madar.NewVector(0.0, 1.0)),
		},
		[]uint32{0, 1, 3, 1, 2, 3},
		&noor.Material{Shader: shader, Textures: []noor.Texture{stoneTexture, wallTexture, tennantTexture}},
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
