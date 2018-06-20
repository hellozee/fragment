/*Package renderer contains the data structure and functions required for the
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
	"github.com/hellozee/fragment/light"
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
func (r *Renderer) DrawFaces(m meshio.Model, col light.Color, l light.Light) {
	for _, face := range m.Faces {
		a, b, c := face.A, face.B, face.C
		var verts = []meshio.Vec3f{m.Verts[a], m.Verts[b], m.Verts[c]}
		temp := verts[2].Subtract(verts[0])
		normal := temp.CrossProduct(verts[1].Subtract(verts[0]))
		normal.Norm()
		shade := l.SurfaceColor(col, normal)
		r.DrawTriangle(verts, shade)
	}
}

//FillTriangle  Function for filling the triangle with the given color
func (r *Renderer) FillTriangle(verts []meshio.Vec2i, c color.Color) {

	temp := meshio.SortByX(verts)
	x1 := temp[0].X
	x2 := temp[2].X
	temp = meshio.SortByY(verts)
	y1 := temp[0].Y
	y2 := temp[2].Y
	verts[1].X, verts[1].Y = verts[1].X-verts[0].X, verts[1].Y-verts[0].Y
	verts[2].X, verts[2].Y = verts[2].X-verts[0].X, verts[2].Y-verts[0].Y
	scalar := verts[1].X*verts[2].Y - verts[2].X*verts[1].Y
	for i := x1; i < x2; i++ {
		p1 := i - verts[0].X
		for j := y1; j < y2; j++ {
			p2 := j - verts[0].Y
			t := p1*(verts[1].Y-verts[2].Y) + p2*(verts[2].X-verts[1].X) +
				scalar
			wa := float64(t) / float64(scalar)
			if wa < 0 || wa > 1 {
				continue
			}
			wb := float64(p1*verts[2].Y-p2*verts[2].X) / float64(scalar)
			if wb < 0 || wb > 1 {
				continue
			}
			wc := float64(p2*verts[1].X-p1*verts[1].Y) / float64(scalar)
			if wc < 0 || wc > 1 {
				continue
			}
			r.img.Set(i, j, c)
		}
	}

}

//DrawTriangle  Function for drawing bare Triangles
func (r *Renderer) DrawTriangle(verts []meshio.Vec3f, c color.Color) {
	var points []meshio.Vec2i

	for i := 0; i < 3; i++ {
		v := verts[i]

		x1 := (v.X + 1.0) * float64(r.width/2)
		y1 := (v.Y + 1.0) * float64(r.height/2)

		p := meshio.Vec2i{X: int(x1), Y: int(y1)}
		points = append(points, p)
	}

	r.FillTriangle(points, c)
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
