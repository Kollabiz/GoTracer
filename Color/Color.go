package Color

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
func (color Color) IMix(color2 Color, factor float32) {
	color.R += (color2.R - color.R) * factor
	color.G += (color2.G - color.G) * factor
	color.B += (color2.B - color.B) * factor
}

// In-place Mul blending
func (color Color) IMul(color2 Color) {
	color.R *= color2.R
	color.G *= color2.G
	color.B *= color2.B
}

// In-place MulF blending
func (color Color) IMulF(factor float32) {
	color.R *= factor
	color.G *= factor
	color.B *= factor
}
