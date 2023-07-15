package Maths

import "math"

type Vector2 struct {
	X float32
	Y float32
}

func MakeVector2(x float32, y float32) Vector2 {
	return Vector2{x, y}
}

func (vec Vector2) Mul(vec2 Vector2) Vector2 {
	return Vector2{vec.X * vec2.X, vec.Y * vec2.Y}
}

func (vec Vector2) MulF(factor float32) Vector2 {
	return Vector2{vec.X * factor, vec.Y * factor}
}

func (vec Vector2) Add(vec2 Vector2) Vector2 {
	return Vector2{vec.X + vec2.X, vec.Y + vec2.Y}
}

func (vec Vector2) Length() float32 {
	return float32(math.Sqrt(float64(vec.X*vec.X + vec.Y + vec.Y)))
}

func (vec Vector2) Normalized() Vector2 {
	l := vec.Length()
	return Vector2{vec.X / l, vec.Y / l}
}
