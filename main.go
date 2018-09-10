package main

import (
	"fmt"
	"image"

	"./light"
	"./meshio"
	"./renderer"
)

func main() {
	width, height := 800, 800
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//white := light.Color{R: 1.0, G: 1.0, B: 1.0, A: 1.0}
	rend := renderer.NewRenderer(img, width, height)

	model := meshio.NewModel("models/african_head.obj")
	fmt.Println(model.Faces)
	texture := meshio.NewTexture("models/african_head_diffuse.png")

	lht := light.Light{
		Position: meshio.Vec3f{X: 0, Y: 0, Z: -1},
		Col: light.Color{
			R: 1.0,
			G: 1.0,
			B: 1.0,
			A: 1.0,
		},
	}

	rend.DrawFaces(model, texture, lht)

	rend.Save("output.png")
}
