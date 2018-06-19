package meshio

import "math"

type Vec3f struct {
	X, Y, Z float64
}

type Vec3i struct {
	A, B, C int
}

func (v1 *Vec3f) Add(v2 Vec3f) Vec3f {
	X := v1.X + v2.X
	Y := v1.Y + v2.Y
	Z := v1.Z + v2.Z

	v := Vec3f{X, Y, Z}

	return v
}

func (v1 *Vec3f) Subtract(v2 Vec3f) Vec3f {
	X := v1.X - v2.X
	Y := v1.Y - v2.Y
	Z := v1.Z - v2.Z

	v := Vec3f{X, Y, Z}

	return v
}

func (v1 *Vec3f) DotProduct(v2 Vec3f) float64 {
	X := v1.X * v2.X
	Y := v1.Y * v2.Y
	Z := v1.Z * v2.Z

	product := X + Y + Z

	return product
}

func (v1 *Vec3f) CrossProduct(v2 Vec3f) Vec3f {
	X := v1.Y*v2.Z - v1.Z*v2.Y
	Y := v1.Z*v2.X - v1.X*v2.Z
	Z := v1.X*v2.Y - v1.Y*v2.X

	v := Vec3f{X, Y, Z}

	return v
}

func (v1 *Vec3f) Norm() float64 {
	X := v1.X * v1.X
	Y := v1.Y * v1.Y
	Z := v1.Z * v1.Z

	sum := X + Y + Z
	root := math.Sqrt(sum)

	return root
}
