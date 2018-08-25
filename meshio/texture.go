package meshio

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

//Texture  Data Structure to hold texture
type Texture struct {
	tex    *image.RGBA
	width  int
	height int
}

//GetColor  Function to fetch the color of the model at the given point
func (t *Texture) GetColor(x, y float64) color.RGBA {
	a := int(x * float64(t.width))
	b := int(y * float64(t.height))
	return t.tex.RGBAAt(a, b)
}

//NewTexture  Function for creating a new texture
func NewTexture(filename string) *Texture {
	file, err := os.Open(filename)
	defer file.Close()
	check(err)

	src, err := png.Decode(file)
	check(err)

	b := src.Bounds()
	imgSet := image.NewRGBA(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			oldPixel := src.At(x, y)
			r, g, b, a := oldPixel.RGBA()
			r = (r + g + b) / 3
			pixel := color.RGBA{uint8(r), uint8(r), uint8(r), uint8(a)}
			imgSet.Set(x, y, pixel)
		}
	}

	tex := Texture{
		tex:    imgSet,
		width:  b.Max.X,
		height: b.Max.Y,
	}

	return &tex
}
