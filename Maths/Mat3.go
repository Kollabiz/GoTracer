package Maths

import (
	"math"
)

type Mat3 struct {
	data [3][3]float32
}

func (mat *Mat3) MatMul(mat2 *Mat3) *Mat3 {
	product := new(Mat3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				product.data[i][j] += mat.data[i][k] * mat2.data[k][j]
			}
		}
	}

	return product
}

func (mat *Mat3) GetElement(x int, y int) float32 {
	return mat.data[x][y]
}

func (mat *Mat3) SetElement(x int, y int, data float32) {
	mat.data[x][y] = data
}

// Matrix constructors

func Mat3FromArray(arr [3][3]float32) *Mat3 {
	mat := new(Mat3)
	mat.data = arr
	return mat
}

func Mat3Identity() *Mat3 {
	arr := [3][3]float32{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	return Mat3FromArray(arr)
}

func Mat3RotationX(angle float64) *Mat3 {
	angle = DegToRad(angle)
	s := float32(math.Sin(angle))
	c := float32(math.Cos(angle))
	arr := [3][3]float32{
		{1, 0, 0},
		{0, c, -s},
		{0, s, c},
	}
	return Mat3FromArray(arr)
}

func Mat3RotationY(angle float64) *Mat3 {
	angle = DegToRad(angle)
	s := float32(math.Sin(angle))
	c := float32(math.Cos(angle))
	arr := [3][3]float32{
		{c, 0, s},
		{0, 1, 0},
		{-s, 0, c},
	}
	return Mat3FromArray(arr)
}

func Mat3RotationZ(angle float64) *Mat3 {
	angle = DegToRad(angle)
	s := float32(math.Sin(angle))
	c := float32(math.Cos(angle))
	arr := [3][3]float32{
		{c, -s, 0},
		{s, c, 0},
		{0, 0, 1},
	}
	return Mat3FromArray(arr)
}

func Mat3FromEulerAngles(euler Vector3) *Mat3 {
	// RotationX(euler.X) * RotationY(euler.Y) * RotationZ(euler.Z)
	return Mat3RotationX(float64(euler.X)).MatMul(Mat3RotationY(float64(euler.Y)).MatMul(Mat3RotationZ(float64(euler.Z))))
}

func Mat3Scale(scale Vector3) *Mat3 {
	arr := [3][3]float32{
		{scale.X, 0, 0},
		{0, scale.Y, 0},
		{0, 0, scale.Z},
	}
	return Mat3FromArray(arr)
}

func Mat3FromVectors(v1 Vector3, v2 Vector3, v3 Vector3) *Mat3 {
	arr := [3][3]float32{
		{v1.X, v1.Y, v1.Z},
		{v2.X, v2.Y, v2.Z},
		{v3.X, v3.Y, v3.Z},
	}
	return Mat3FromArray(arr)
}
