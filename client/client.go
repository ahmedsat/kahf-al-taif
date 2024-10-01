package client

import (
	"errors"
)

func StartClient(url string) (err error) {

	err = render()
	if err != nil {
		err = errors.Join(err, errors.New("failed to render"))
	}

	// if err := noor.Init(noor.Options{
	// 	Width:      windowWidth,
	// 	Height:     windowHeight,
	// 	Title:      windowTitle,
	// 	Background: [3]float32{0.1, 0.1, 0.1},
	// }); err != nil {
	// 	return err
	// }

	// // Create shader program
	// shader, err = noor.CreateShaderProgramFromFiles("shaders/base.vert", "shaders/base.frag")
	// if err != nil {
	// 	return err
	// }

	// // Load textures
	// textures, err = loadTextures()
	// if err != nil {
	// 	return err
	// }

	// // Create cube mesh
	// mesh, err = createCubeMesh()
	// if err != nil {
	// 	return err
	// }

	// // Create material
	// material = &noor.Material{
	// 	Shader:   shader,
	// 	Textures: textures,
	// }

	// // Create camera
	// camera := noor.NewCamera(
	// 	madar.Vector3{-3, -3, -3},
	// 	madar.Vector3{},
	// 	madar.Vector3{0, 1, 0},
	// 	windowWidth/windowHeight,
	// 	noor.Perspective)

	// // obj, err := noor.NewObject(mesh, material, madar.Identity())
	// // if err != nil {
	// // 	return err
	// // }

	// // Create scene
	// scene := noor.NewScene(camera, createRandomCubes(10)...)

	// // Main loop
	// return noor.Run(func() {
	// 	updateScene(scene)
	// 	scene.Draw()
	// })
	return
}
