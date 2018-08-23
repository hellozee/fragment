package meshio

import (
	"math"
)

//Vec3f  Data Structure for holding a 3d vector with float data
type Vec3f struct {
	X, Y, Z float64
}

//Vec3i  Data Structure for holding a 3d vector with integer data
type Vec3i struct {
	A, B, C int
}

//Vec2i  Data Structure for holding a 2d vector with integer data
type Vec2i struct {
	X, Y int
}

//Vec2f  Data Structure for holding a 2d vector with integer data
type Vec2f struct {
	X, Y float64
}

//Add  Function for adding 2 Vectors
func (v1 *Vec3f) Add(v2 Vec3f) Vec3f {
	X := v1.X + v2.X
	Y := v1.Y + v2.Y
	Z := v1.Z + v2.Z

	v := Vec3f{X, Y, Z}

	return v
}

//Subtract  Function for subtracting v2 from v1
func (v1 *Vec3f) Subtract(v2 Vec3f) Vec3f {
	X := v1.X - v2.X
	Y := v1.Y - v2.Y
	Z := v1.Z - v2.Z

	v := Vec3f{X, Y, Z}

	return v
}

//DotProduct  Function for calculating the dot product of 2 vectors
func (v1 *Vec3f) DotProduct(v2 Vec3f) float64 {
	X := v1.X * v2.X
	Y := v1.Y * v2.Y
	Z := v1.Z * v2.Z

	product := X + Y + Z

	return product
}

//CrossProduct  Function for calculating the cross product of 2 vectors
func (v1 *Vec3f) CrossProduct(v2 Vec3f) Vec3f {
	X := v1.Y*v2.Z - v1.Z*v2.Y
	Y := v1.Z*v2.X - v1.X*v2.Z
	Z := v1.X*v2.Y - v1.Y*v2.X

	v := Vec3f{X, Y, Z}

	return v
}

//Norm  Function for calculating the Norm of a vector
func (v1 *Vec3f) Norm() {
	X := v1.X * v1.X
	Y := v1.Y * v1.Y
	Z := v1.Z * v1.Z

	sum := X + Y + Z
	root := math.Sqrt(sum)

	v1.X /= root
	v1.Y /= root
	v1.Z /= root
}

//SortByX  Function for sorting vertices according to their x co-ordinate
func SortByX(a [3]Vec2i) [3]Vec2i {

	for i := 1; i < 3; i++ {
		j := i
		for j > 0 {
			if a[j-1].X > a[j].X {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}

	return a
}

//SortByY  Function for sorting vertices according to their y co-ordinate
func SortByY(a [3]Vec2i) [3]Vec2i {

	for i := 1; i < 3; i++ {
		j := i
		for j > 0 {
			if a[j-1].Y > a[j].Y {
				a[j-1], a[j] = a[j], a[j-1]
			}
			j = j - 1
		}
	}

	return a
}
