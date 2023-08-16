package LightSources

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"RayTracer/RayTracing/Structures"
	"math"
)

type DirectionalLight struct {
	Color     Color.Color
	Direction Maths.Vector3
	Intensity float32
}

func (light *DirectionalLight) GetLight(position Maths.Vector3, normal Maths.Vector3, excludeTri *Structures.Triangle, ctx *Structures.RenderContext) Color.Color {
	d := float32(math.Max(float64(normal.Dot(light.Direction)), float64(ctx.MinLightIntensity))) * light.Intensity
	ray := Structures.CastRay(position.Add(light.Direction.MulF(Epsilon)), light.Direction, ctx, excludeTri)
	if ray.Hit {
		return light.Color.MulF(ctx.MinLightIntensity * d)
	} else {
		return light.Color.MulF(d)
	}
}
