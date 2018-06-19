package meshio

import "math"

//Vec3f  Data Structure for holding a 3d vector with float data
type Vec3f struct {
	X, Y, Z float64
}

//Vec3i  Data Structure for holding a 3d vector with integer data
type Vec3i struct {
	A, B, C int
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
func (v1 *Vec3f) Norm() float64 {
	X := v1.X * v1.X
	Y := v1.Y * v1.Y
	Z := v1.Z * v1.Z

	sum := X + Y + Z
	root := math.Sqrt(sum)

	return root
}
