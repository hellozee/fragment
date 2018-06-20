package meshio

import (
	"math"
	"math/rand"
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
func SortByX(a []Vec2i) []Vec2i {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i].X < a[right].X {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	SortByX(a[:left])
	SortByX(a[left+1:])

	return a
}

//SortByY  Function for sorting vertices according to their y co-ordinate
func SortByY(a []Vec2i) []Vec2i {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i].Y < a[right].Y {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	SortByY(a[:left])
	SortByY(a[left+1:])

	return a
}
