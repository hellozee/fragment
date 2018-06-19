/*Package renderer  contains the data structure and functions required for the
renderer.
  The renderer is the heart of fragment, it is the part which draws the
  lines and points according to the data provided by the other packages.
*/
package renderer

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hellozee/fragment/flipper"
	"github.com/hellozee/fragment/meshio"
)

//Renderer  Data Structure for holding the Renderer
type Renderer struct {
	img           *image.RGBA
	width, height int
}

//DrawLine  Function for drawing straight Lines
func (r *Renderer) DrawLine(x1, y1, x2, y2 int, c color.Color) {
	var xInc, yInc, dx, dy int

	if x2 > x1 {
		dx = x2 - x1
		xInc = 1
	} else {
		dx = x1 - x2
		xInc = -1
	}

	if y2 > y1 {
		dy = y2 - y1
		yInc = 1
	} else {
		dy = y1 - y2
		yInc = -1
	}

	x, y := x1, y1

	r.img.Set(x, y, c)

	if dx > dy {
		err := 2 * (dy - dx)

		for i := 0; i < dx; i++ {
			if err >= 0 {
				y += yInc
				err += 2 * (dy - dx)
			} else {
				err += 2 * dy
			}
			x += xInc
			r.img.Set(x, y, c)
		}
	} else {
		err := 2 * (dx - dy)

		for i := 0; i < dy; i++ {
			if err >= 0 {
				x += xInc
				err += 2 * (dx - dy)
			} else {
				err += 2 * dx
			}
			y += yInc
			r.img.Set(x, y, c)
		}
	}
}

//DrawFaces  Function for Drawing Triangular Faces
func (r *Renderer) DrawFaces(m meshio.Model, col color.Color) {
	for _, face := range m.Faces {
		a, b, c := face.A, face.B, face.C
		var verts = []meshio.Vec3f{m.Verts[a], m.Verts[b], m.Verts[c]}
		r.DrawTriangle(verts, col)
	}
}

//DrawTriangle  Function for drawing bare Triangles
func (r *Renderer) DrawTriangle(verts []meshio.Vec3f, c color.Color) {
	for i := 0; i < 3; i++ {
		v1 := verts[i]
		v2 := verts[(i+1)%3]

		x1 := (v1.X + 1.0) * float64(r.width/2)
		y1 := (v1.Y + 1.0) * float64(r.height/2)
		x2 := (v2.X + 1.0) * float64(r.width/2)
		y2 := (v2.Y + 1.0) * float64(r.height/2)

		r.DrawLine(int(x1), int(y1), int(x2), int(y2), c)
	}
}

//Save  Function for saving the Image to a png File
func (r *Renderer) Save(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, flipper.FlipV(r.img))
}

//NewRenderer  Function for creating a new Renderer
func NewRenderer(i *image.RGBA, w int, h int) *Renderer {
	r := Renderer{
		img:    i,
		width:  w,
		height: h,
	}

	return &r
}
