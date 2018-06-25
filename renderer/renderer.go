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
	"math"
	"os"

	"github.com/hellozee/fragment/flipper"
	"github.com/hellozee/fragment/light"
	"github.com/hellozee/fragment/meshio"
)

//Renderer  Data Structure for holding the Renderer
type Renderer struct {
	img           *image.RGBA
	width, height int
	zBuffer       []float64
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
		var verts = [3]meshio.Vec3f{m.Verts[a], m.Verts[b], m.Verts[c]}
		temp := verts[2].Subtract(verts[0])
		normal := temp.CrossProduct(verts[1].Subtract(verts[0]))
		normal.Norm()
		shade := l.SurfaceColor(col, normal)
		r.DrawTriangle(verts, shade)
	}
}

//FillTriangle  Function for filling the triangle with the given color
func (r *Renderer) FillTriangle(verts [3]meshio.Vec2i, c color.Color, original [3]meshio.Vec3f) {

	temp := meshio.SortByX(verts)
	x1 := temp[0].X
	x2 := temp[2].X
	temp = meshio.SortByY(verts)
	y1 := temp[0].Y
	y2 := temp[2].Y

	for i := x1; i < x2; i++ {
		for j := y1; j < y2; j++ {
			screen := barycentricCoords(verts, meshio.Vec2i{X: i, Y: j})
			if screen.X < 0 || screen.Y < 0 || screen.Z < 0 {
				continue
			}
			z := original[0].Z*screen.X + original[1].Z*screen.Y +
				original[2].Z*screen.Z
			if r.zBuffer[i+j*r.width] < z {
				r.zBuffer[i+j*r.width] = z
				r.img.Set(i, j, c)
			}
		}
	}

}

//DrawTriangle  Function for drawing bare Triangles
func (r *Renderer) DrawTriangle(verts [3]meshio.Vec3f, c color.Color) {
	var points [3]meshio.Vec2i

	for i := 0; i < 3; i++ {
		v := verts[i]

		x1 := (v.X + 1.0) * float64(r.width/2)
		y1 := (v.Y + 1.0) * float64(r.height/2)

		p := meshio.Vec2i{X: int(x1), Y: int(y1)}
		points[i] = p
	}

	r.FillTriangle(points, c, verts)
}

func barycentricCoords(pts [3]meshio.Vec2i, P meshio.Vec2i) meshio.Vec3f {
	v1 := meshio.Vec3f{
		X: float64(pts[2].X - pts[0].X),
		Y: float64(pts[1].X - pts[0].X),
		Z: float64(pts[0].X - P.X),
	}

	v2 := meshio.Vec3f{
		X: float64(pts[2].Y - pts[0].Y),
		Y: float64(pts[1].Y - pts[0].Y),
		Z: float64(pts[0].Y - P.Y),
	}

	u := v1.CrossProduct(v2)

	if math.Abs(u.Y) < 1 {
		return meshio.Vec3f{
			X: -1,
			Y: 1,
			Z: 1,
		}
	}

	return meshio.Vec3f{
		X: 1.0 - (u.X+u.Y)/u.Z,
		Y: u.Y / u.Z,
		Z: u.X / u.Z,
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
	buffer := make([]float64, w*h)

	for num := range buffer {
		buffer[num] = -1 * math.MaxFloat64
	}

	r := Renderer{
		img:     i,
		width:   w,
		height:  h,
		zBuffer: buffer,
	}

	return &r
}
