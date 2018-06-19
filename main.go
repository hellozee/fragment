package main

import (
	"image"
	"image/png"
	"os"
)

func main() {
	width, height := 800, 800
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
