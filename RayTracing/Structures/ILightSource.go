package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"math"
)

type ILightSource interface {
	GetLight(position Maths.Vector3, normal Maths.Vector3, excludeTri *Triangle, ctx *RenderContext) Color.Color
}

func TraceSoftShadow(position Maths.Vector3, lightPosition Maths.Vector3, lightRadius float32, ctx *RenderContext, excludeTriangle *Triangle) float32 {
	rayCount := int(math.Max(float64(lightRadius)*float64(ctx.RenderSettings.ShadowSamples), float64(ctx.RenderSettings.ShadowSamples)))
	rayNormal := lightPosition.Sub(position).Normalized()
	var energy float32 = 0
	dist := position.Sub(lightPosition).Length()
	sampleRadius := Maths.Vector3{Y: lightRadius, Z: dist}.Normalized().Dot(Maths.Vector3{Y: 1})
	for i := 0; i < rayCount; i++ {
		rayDirection := Maths.RandomPointOnHemisphere(rayNormal, sampleRadius)
		ray := TraceRay(position, rayDirection, ctx, excludeTriangle)
		if !ray.Hit || ray.RayLength > dist {
			energy++
		} else {
			energy += ctx.MinLightIntensity
		}
	}
	return energy / float32(rayCount)
}

func TraceSimpleShadow(position Maths.Vector3, lightPosition Maths.Vector3, ctx *RenderContext, excludeTriangle *Triangle) float32 {
	rayDir := lightPosition.Sub(position).Normalized()
	dist := lightPosition.Sub(position).Length()
	ray := TraceRay(position, rayDir, ctx, excludeTriangle)
	if !ray.Hit || ray.RayLength > dist {
		return 1
	} else {
		return ctx.MinLightIntensity
	}
}
