package light

import (
	"errors"
	"image/color"

	"github.com/hellozee/fragment/meshio"
)

//Color  Data Structure for holding float64 color values
type Color struct {
	R, G, B, A float64
}

//Light  Data Structure for holding information of a scene light
type Light struct {
	Col      Color
	Position meshio.Vec3f
}

//SurfaceColor  Function for determining color of the given surface
func (l *Light) SurfaceColor(diffuse color.RGBA, normal meshio.Vec3f) color.Color {
	intensity := normal.DotProduct(l.Position)
	i := uint8(intensity)
	c := color.RGBA{
		R: diffuse.R * i,
		G: diffuse.G * i,
		B: diffuse.B * i,
		A: diffuse.A,
	}

	return c
}

//Product  Function for multiplying 2 color values
func (c *Color) Product(c1 Color) Color {
	col := Color{}
	col.A = c1.A * c.A
	col.R = c1.R * c.R
	col.G = c1.G * c.G
	col.B = c1.B * c.B

	return col
}

//ScalarProduct  Function for multipying a color with a scalar value
func (c *Color) ScalarProduct(v float64) {
	c.R *= v
	c.G *= v
	c.B *= v
}

//NewLight  Function for generating a new Light
func NewLight(pos meshio.Vec3f, r, g, b, a float64) (*Light, error) {
	if checkValue(r) || checkValue(g) || checkValue(b) || checkValue(a) {
		return &Light{}, errors.New("Color Value must be between 0 and 1")
	}
	l := Light{
		Col: Color{
			R: r,
			G: g,
			B: b,
			A: a,
		},
		Position: pos,
	}

	return &l, nil
}

//NewColor  Function for generating a new Color
func NewColor(r, g, b, a float64) (*Color, error) {
	if checkValue(r) || checkValue(g) || checkValue(b) || checkValue(a) {
		return &Color{}, errors.New("Color Value must be between 0 and 1")
	}
	c := Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	return &c, nil
}

func checkValue(a float64) bool {
	if a < 0 || a > 1 {
		return true
	}
	return false
}
