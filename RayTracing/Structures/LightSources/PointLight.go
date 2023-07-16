package LightSources

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"RayTracer/RayTracing/Structures"
	"math"
)

const (
	Epsilon = 0.0000001
)

type PointLight struct {
	Color     Color.Color
	Radius    float32
	Position  Maths.Vector3
	Intensity float32
}

func (light *PointLight) GetLight(position Maths.Vector3, normal Maths.Vector3, excludeTri *Structures.Triangle, ctx *Structures.RenderContext) Color.Color {
	distance := light.Position.Sub(position).Length()
	dot := float32(math.Max(float64(light.Position.Sub(position).Normalized().Dot(normal)), 0))
	if distance < Epsilon {
		return light.Color.MulF(light.Intensity)
	} // Zero division
	mul := float32(math.Min(float64(light.Intensity/distance*dot), 1))
	shadow := Structures.TraceSoftShadow(position, light.Position, light.Radius, ctx, excludeTri)
	mul *= shadow
	return light.Color.MulF(mul)
}
