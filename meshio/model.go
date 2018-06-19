/*Package meshio  Package contains the data structures and functions
  required for reading a Wavefront obj to format which
  can be used by the renderer to properly render an image.
*/

package meshio

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Model  Data Structure fol holding a Wavefront obj Model
type Model struct {
	Verts []Vec3f
	Faces []Vec3i
}

//NewModel  Function for parsing a Wavefront obj and creating a new model
func NewModel(filename string) Model {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var m Model

	for scanner.Scan() {
		line := scanner.Text()
		tokenized := strings.Split(line, " ")

		if tokenized[0] == "v" {
			var vertex Vec3f
			vertex.X, _ = strconv.ParseFloat(tokenized[1], 64)
			vertex.Y, _ = strconv.ParseFloat(tokenized[2], 64)
			vertex.Z, _ = strconv.ParseFloat(tokenized[3], 64)

			m.Verts = append(m.Verts, vertex)

		} else if tokenized[0] == "f" {
			var f []int
			for i := 1; i <= 3; i++ {
				broken := strings.Split(tokenized[i], "/")
				num, _ := strconv.Atoi(broken[0])
				f = append(f, num)
			}

			var face Vec3i
			face.A = f[0] - 1
			face.B = f[1] - 1
			face.C = f[2] - 1

			m.Faces = append(m.Faces, face)
		}
	}

	return m

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
