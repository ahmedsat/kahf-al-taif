package utils

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	_ "golang.org/x/image/webp"
)

func LoadImages(path string) (img image.Image, err error) {

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	if err != nil {
		return
	}

	return

}
