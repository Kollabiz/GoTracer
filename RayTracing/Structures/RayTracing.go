package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
)

const (
	Epsilon = 0.0000001
)

func TraceStartRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext) Color.Color {
	ray := CastRay(origin, direction, ctx, nil)
	if !ray.Hit {
		return ctx.BackgroundColor
	}
	nextRay := TraceRay(ray.HitPosition, direction.Reflect(ray.GetHitNormal()), ctx, Color.MakeColor(1, 1, 1), &ray, 0)
	return nextRay
}

func TraceRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext, energy Color.Color, previousRay *RayIntersection, depth int) Color.Color {
	var illumination = Color.MakeColor(0, 0, 0)
	prevUv := previousRay.GetHitUv()
	prevRoughness := 1 - previousRay.GetHitMaterial().Glossiness
	prevSpecular := previousRay.GetHitMaterial().GetSpecular(prevUv)
	prevDiffuse := previousRay.GetHitMaterial().GetDiffuse(prevUv)
	samples := 0
	for i := 0; i < ctx.RenderSettings.IndirectIlluminationSampleCount; i++ {
		rayDirection := Maths.RandomPointOnHemisphere(previousRay.GetHitNormal(), 1)
		ray := CastRay(origin, rayDirection, ctx, previousRay.HitTriangle)
		if !ray.Hit {
			continue
		}
		illumination.IUAdd(calcDirectIllumination(origin, ray.GetHitNormal(), ctx, ray.HitTriangle))
		rayDiffuse := ray.GetHitMaterial().GetDiffuse(ray.GetHitUv())
		rayEnergy := energy
		samples++
		// If this ray can be used as a reflection sample
		if rayDirection.Dot(direction) < prevRoughness {
			reflColor := prevSpecular.Mul(rayDiffuse).Mul(energy)
			illumination.IUAdd(reflColor)
			rayEnergy.IMul(prevSpecular)
		} else {
			illumination.IUAdd(rayDiffuse.Mul(energy))
			rayEnergy.IMul(prevDiffuse)
		}
		// Casting next ray
		if depth < ctx.RenderSettings.IndirectIlluminationDepth {
			nextRayColor := TraceRay(ray.HitPosition, rayDirection.Reflect(ray.GetHitNormal()), ctx, rayEnergy, &ray, depth+1)
			illumination.IUAdd(nextRayColor)
			illumination.IDivF(2)
		}
	}
	if samples > 0 {
		illumination.IDivF(float32(samples))
	} else {
		illumination = ctx.BackgroundColor
	}
	return illumination
}

func calcDirectIllumination(position Maths.Vector3, normal Maths.Vector3, ctx *RenderContext, excludeTriangle *Triangle) Color.Color {
	var directIllumination = Color.MakeColor(0, 0, 0)
	for i := 0; i < len(ctx.Scene.Lights); i++ {
		light := ctx.Scene.Lights[i]
		illum := light.GetLight(position, normal, excludeTriangle, ctx)
		directIllumination.IUAdd(illum)
	}
	return directIllumination
}
