package Maths

import "math"

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func MakeVector3(x float32, y float32, z float32) Vector3 {
	return Vector3{x, y, z}
}

func (vector Vector3) Length() float64 {
	return math.Sqrt(math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.Y), 2) + math.Pow(float64(vector.Z), 2))
}

func (vector Vector3) Dot(v2 Vector3) float32 {
	return vector.X*v2.X + vector.Y*v2.Y + vector.Z*v2.Z
}

func (vector Vector3) Cross(v2 Vector3) Vector3 {
	cX := vector.Y*v2.Z - vector.Z*v2.Y
	cY := vector.Z*v2.X - vector.X*v2.Z
	cZ := vector.X*v2.Y - vector.Y*v2.X
	return Vector3{cX, cY, cZ}
}

func (vector Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{vector.X - v2.X, vector.Y - v2.Y, vector.Z - v2.Z}
}

func (vector Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{vector.X + v2.X, vector.Y + v2.Y, vector.Z + v2.Z}
}

func (vector Vector3) Mul(v2 Vector3) Vector3 {
	return Vector3{vector.X * v2.X, vector.Y * v2.Y, vector.Z * v2.Z}
}

func (vector Vector3) MulF(f float32) Vector3 {
	return Vector3{vector.X * f, vector.Y * f, vector.Z * f}
}

func (vector Vector3) Div(v2 Vector3) Vector3 {
	return Vector3{vector.X / v2.X, vector.Y / v2.Y, vector.Z / v2.Z}
}

func (vector Vector3) DivF(f float32) Vector3 {
	return Vector3{vector.X / f, vector.Y / f, vector.Z / f}
}

// Matrix multiplication
func (vector Vector3) MatMul(mat *Mat3) Vector3 {
	multiplied := new(Vector3)
	multiplied.X = vector.X*mat.GetElement(0, 0) + vector.Y*mat.GetElement(1, 0) + vector.Z*mat.GetElement(2, 0)
	multiplied.Y = vector.X*mat.GetElement(0, 1) + vector.Y*mat.GetElement(1, 1) + vector.Z*mat.GetElement(2, 1)
	multiplied.Z = vector.X*mat.GetElement(0, 2) + vector.Y*mat.GetElement(1, 2) + vector.Z*mat.GetElement(2, 2)
	return *multiplied
}

func ZeroVector3() Vector3 {
	return Vector3{0, 0, 0}
}
