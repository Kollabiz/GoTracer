package Maths

import (
	"math"
)

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func MakeVector3(x float32, y float32, z float32) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

func (vector Vector3) Length() float32 {
	return float32(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)))
}

func (vector Vector3) LengthSqr() float32 {
	return vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z
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

func (vector Vector3) Normalized() Vector3 {
	return vector.MulF(FastInverseSqrt(vector.LengthSqr()))
}

func (vector Vector3) Reflect(normal Vector3) Vector3 {
	return vector.Sub(normal.MulF(vector.Dot(normal) * 2))
}

func (vector Vector3) Lerp(to Vector3, factor float32) Vector3 {
	return Vector3{
		Lerp(vector.X, to.X, factor),
		Lerp(vector.Y, to.Y, factor),
		Lerp(vector.Z, to.Z, factor),
	}
}

func (vector Vector3) Invert() Vector3 {
	return Vector3{
		-vector.X,
		-vector.Y,
		-vector.Z,
	}
}

func (vector Vector3) Abs() Vector3 {
	return Vector3{
		float32(math.Abs(float64(vector.X))),
		float32(math.Abs(float64(vector.Y))),
		float32(math.Abs(float64(vector.Z))),
	}
}

// Comparing

func AreEqualVectors(vec1 Vector3, vec2 Vector3) bool {
	return vec1.X == vec2.X && vec1.Y == vec2.Y && vec1.Z == vec2.Z
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

func InfiniteVector3(sign int) Vector3 {
	return Vector3{float32(math.Inf(sign)), float32(math.Inf(sign)), float32(math.Inf(sign))}
}
