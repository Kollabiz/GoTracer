package Color

import (
	"RayTracer/Maths"
	"image/color"
	"math"
)

type Color struct {
	R float32
	G float32
	B float32
}

func MakeColor(r float32, g float32, b float32) Color {
	return Color{r, g, b}
}

// Mix blending between two colors
func (color Color) Mix(color2 Color, factor float32) Color {
	return Color{
		color.R + (color2.R-color.R)*factor,
		color.G + (color2.G-color.G)*factor,
		color.B + (color2.B-color.B)*factor,
	}
}

// Mul blending between two colors
func (color Color) Mul(color2 Color) Color {
	return Color{color.R * color2.R, color.G * color2.G, color.B * color2.B}
}

// Mul bledning between a color and float32
func (color Color) MulF(factor float32) Color {
	return Color{color.R * factor, color.G * factor, color.B * factor}
}

// In-place Mix blending
func (color *Color) IMix(color2 Color, factor float32) {
	color.R += (color2.R - color.R) * factor
	color.G += (color2.G - color.G) * factor
	color.B += (color2.B - color.B) * factor
}

// In-place Mul blending
func (color *Color) IMul(color2 Color) {
	color.R *= color2.R
	color.G *= color2.G
	color.B *= color2.B
}

// In-place MulF blending
func (color *Color) IMulF(factor float32) {
	color.R *= factor
	color.G *= factor
	color.B *= factor
}

// Addition
func (color Color) Add(color2 Color) Color {
	return Color{
		float32(math.Min(float64(color.R+color2.R), 1)),
		float32(math.Min(float64(color.G+color2.G), 1)),
		float32(math.Min(float64(color.B+color2.B), 1)),
	}
}

// Addition without clamping
func (color Color) UAdd(color2 Color) Color {
	return Color{
		color.R + color2.R,
		color.G + color2.G,
		color.B + color2.B,
	}
}

func (color *Color) IUAdd(color2 Color) {
	color.R += color2.R
	color.G += color2.G
	color.B += color2.B
}

func (color *Color) IDivF(divisor float32) {
	color.R /= divisor
	color.G /= divisor
	color.B /= divisor
}

func (color Color) DivF(divisor float32) Color {
	return Color{color.R / divisor, color.G / divisor, color.B / divisor}
}

func (color Color) Grayscale() float32 {
	return (color.R + color.G + color.B) / 3
}

func (color Color) VLength() float32 {
	return float32(math.Sqrt(float64(color.R*color.R + color.B*color.B + color.G*color.G)))
}

func (color Color) Normalize() {
	l := color.VLength()
	color.R /= l
	color.G /= l
	color.B /= l
}

func FromImageColor(clr color.Color) Color {
	r, g, b, _ := clr.RGBA()
	return Color{float32(r) / 255, float32(g) / 255, float32(b) / 255}
}

func (clr *Color) ToImageColor() color.Color {
	r, g, b := uint8(clr.R*255), uint8(clr.G*255), uint8(clr.B*255)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func (clr Color) Dot(vector Maths.Vector3) float32 {
	return clr.R*vector.X + clr.G*vector.Y + clr.B*vector.Z
}

func (clr Color) DotC(color Color) float32 {
	return clr.R*color.R + clr.G*color.G + clr.B*color.B
}

func MakeColor2DArray(dim1 int, dim2 int) [][]Color {
	arr := make([][]Color, dim1)
	for i := 0; i < dim1; i++ {
		arr[i] = make([]Color, dim2)
	}
	return arr
}
