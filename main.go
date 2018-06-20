package main

import (
	"image"
	"image/color"

	"github.com/hellozee/fragment/light"
	"github.com/hellozee/fragment/meshio"
	"github.com/hellozee/fragment/renderer"
)

func main() {
	width, height := 800, 800
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	//red := color.RGBA{255, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}

	rend := renderer.NewRenderer(img, width, height)

	model := meshio.NewModel("models/african_head.obj")

	lht := light.Light{
		Position: meshio.Vec3f{X: 0, Y: 0, Z: -1},
		Col: light.Color{
			R: 1.0,
			G: 1.0,
			B: 1.0,
			A: 1.0,
		},
	}

	rend.DrawFaces(model, white, lht)

	rend.Save("output.png")
}
