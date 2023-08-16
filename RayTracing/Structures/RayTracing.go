package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"sync"
	"sync/atomic"
)

const (
	Epsilon = 0.0000001
)

func TraceStartRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext, imgBuff *ImageBuffer, fragmentPosX int, fragmentPosY int, waitGroup *sync.WaitGroup) {
	ray := CastRay(origin, direction, ctx, nil)
	if !ray.Hit {
		imgBuff.Add(fragmentPosX, fragmentPosY, ctx.BackgroundColor)
		waitGroup.Done()
		atomic.AddInt32(&ctx.Scene.RenderedPixels, 1)
		return
	}
	dirIll := calcDirectIllumination(origin, ray.GetHitNormal(), ctx).Mul(ray.GetHitMaterial().GetDiffuse(ray.GetHitUv()))
	nextRay := TraceRay(ray.HitPosition, direction.Reflect(ray.GetHitNormal()), ctx, ray.GetHitMaterial().GetDiffuse(ray.GetHitUv()), dirIll, &ray, 0)
	imgBuff.Add(fragmentPosX, fragmentPosY, nextRay)
	waitGroup.Done()
	atomic.AddInt32(&ctx.Scene.RenderedPixels, 1)
}

func TraceRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext, energy Color.Color, prevIllumination Color.Color, previousRay *RayIntersection, depth int) Color.Color {
	prevUv := previousRay.GetHitUv()
	prevRoughness := 1 - previousRay.GetHitMaterial().Glossiness
	prevSpecular := previousRay.GetHitMaterial().GetSpecular(prevUv)
	prevDiffuse := previousRay.GetHitMaterial().GetDiffuse(prevUv)
	var illumination = prevIllumination
	samples := 0
	for i := 0; i < ctx.RenderSettings.IndirectIlluminationSampleCount; i++ {
		rayDirection := Maths.RandomPointOnHemisphere(previousRay.GetHitNormal(), 1)
		ray := CastRay(origin.Add(rayDirection.MulF(Epsilon)), rayDirection, ctx, previousRay.HitTriangle)
		if !ray.Hit {
			continue
		}
		rayDiffuse := ray.GetHitMaterial().GetDiffuse(ray.GetHitUv())
		illumination.IUAdd(calcDirectIllumination(origin.Add(ray.GetHitNormal().MulF(Epsilon)), ray.GetHitNormal(), ctx).Mul(rayDiffuse))
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
			nextRayColor := TraceRay(ray.HitPosition, rayDirection.Reflect(ray.GetHitNormal()), ctx, rayEnergy, Color.MakeColor(0, 0, 0), &ray, depth+1)
			illumination.IUAdd(nextRayColor)
			illumination.IDivF(2)
		}
	}
	if samples > 0 {
		illumination.IDivF(float32(samples))
	}
	return illumination
}

func calcDirectIllumination(position Maths.Vector3, normal Maths.Vector3, ctx *RenderContext) Color.Color {
	var directIllumination = Color.MakeColor(0, 0, 0)
	for i := 0; i < len(ctx.Scene.Lights); i++ {
		light := ctx.Scene.Lights[i]
		illum := light.GetLight(position.Add(normal.MulF(Epsilon)), normal, nil, ctx)
		directIllumination.IUAdd(illum)
	}
	return directIllumination
}
