package Maths

type Vector2 struct {
	X float32
	Y float32
}

func MakeVector2(x float32, y float32) Vector2 {
	return Vector2{x, y}
}
