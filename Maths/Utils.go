package Maths

import (
	"RayTracer/Color"
	"math/rand"
)

func DegToRad(deg float64) float64 {
	return deg * 0.01745329
}

func TriInterpolate(v1 Vector3, v2 Vector3, v3 Vector3, factor Vector3) Vector3 {
	interpolated := v1.MulF(factor.Z).Add(v2.MulF(factor.X)).Add(v3.MulF(factor.Y))
	return interpolated.Normalized()
}

func TriInterpolate2(v1 Vector2, v2 Vector2, v3 Vector2, factor Vector3) Vector2 {
	interpolated := v1.MulF(factor.Z).Add(v2.MulF(factor.X)).Add(v3.MulF(factor.Y))
	return interpolated.Normalized()
}

func RandomPointOnHemisphere(normal Vector3, radius float32) Vector3 {
	x := float32(rand.NormFloat64() + float64(normal.X))
	y := float32(rand.NormFloat64()+float64(normal.Y)) * radius
	z := float32(rand.NormFloat64()+float64(normal.Z)) * radius
	vec := MakeVector3(x, y, z)
	vec = vec.Div(vec.Normalized())
	return vec
}

func Lerp(from float32, to float32, factor float32) float32 {
	return from*(1-factor) + to*factor
}

func MakeColor2DArray(dim1 int, dim2 int) [][]Color.Color {
	arr := make([][]Color.Color, dim1)
	for i := 0; i < dim1; i++ {
		arr[i] = make([]Color.Color, dim2)
	}
	return arr
}

func MakeFloat2DArray(dim1 int, dim2 int) [][]float32 {
	arr := make([][]float32, dim1)
	for i := 0; i < dim1; i++ {
		arr[i] = make([]float32, dim2)
	}
	return arr
}

func MakeBool2DArray(dim1 int, dim2 int) [][]bool {
	arr := make([][]bool, dim1)
	for i := 0; i < dim1; i++ {
		arr[i] = make([]bool, dim2)
	}
	return arr
}

func MakeVector2DArray(dim1 int, dim2 int) [][]Vector3 {
	arr := make([][]Vector3, dim1)
	for i := 0; i < dim1; i++ {
		arr[i] = make([]Vector3, dim2)
	}
	return arr
}

func LensSizeFromAspectRatio(width int, height int) Vector2 {
	aspect := float32(width) / float32(height)
	return Vector2{1, aspect}
}
