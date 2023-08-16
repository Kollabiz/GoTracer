package Maths

import (
	"math"
	"math/rand"
)

func Frac(f float64) float64 {
	return math.Mod(f, 1)
}

func Clamp(f float32, min float32, max float32) float32 {
	return float32(math.Max(math.Min(float64(f), float64(max)), float64(min)))
}

func Rand(co Vector2) float64 {
	return Frac(math.Sin(float64(co.Dot(Vector2{12.9898, 78.233}))) * 43758.5453)
}

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
	azimuthal := 2 * math.Pi * (math.Mod(rand.NormFloat64(), 1))
	sin2Zenith := math.Mod(rand.NormFloat64(), 1)
	sinZenith := math.Sqrt(sin2Zenith)
	x := sinZenith * math.Cos(azimuthal) * float64(radius)
	y := sinZenith * math.Sin(azimuthal) * float64(radius)
	z := math.Sqrt(1 - sin2Zenith)
	return Vector3{
		float32(x),
		float32(y),
		float32(z),
	}.MatMul(GetTangentSpace(normal))
}

func FastInverseSqrt(x float32) float32 {
	i := math.Float32bits(x)
	j := 0x5f3759df - (i >> 1)
	y := math.Float32frombits(j)
	return y * (1.5 - 0.5*x*y*y)
}

func GetTangentSpace(normal Vector3) *Mat3 {
	var helper Vector3
	if normal.X > 0.99 {
		helper = Vector3{0, 0, 1}
	} else {
		helper = Vector3{1, 0, 0}
	}

	tangent := normal.Cross(helper).Normalized()
	binormal := normal.Cross(tangent).Normalized()
	return Mat3FromVectors(
		tangent,
		binormal,
		normal,
	)
}

func Lerp(from float32, to float32, factor float32) float32 {
	return from*(1-factor) + to*factor
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
