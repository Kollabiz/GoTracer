package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
)

func TraceStartRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext) RayResult {
	ray := TraceRay(origin, direction, ctx, nil)
	if !ray.Hit {
		return MakeNonHitResult()
	}
	uv := ray.HitTriangle.GetUv(ray.BarycentricHit)
	rayDiffuse := ray.GetHitMaterial().GetDiffuse(uv)
	raySpecular := ray.GetHitMaterial().GetSpecular(uv)
	rayNormal := ray.HitTriangle.TriangleNormal
	rayGlossiness := ray.GetHitMaterial().Glossiness
	directLight := TraceDirectIlluminationRay(ray.HitPosition, rayNormal, ray.HitTriangle, ctx)
	globalIllumination := TraceIndirectIllumination(direction.Reflect(rayNormal), &ray, rayGlossiness, ctx, 0)
	result := RayResult{}
	result.Diffuse = rayDiffuse                                             // Diffuse pass
	result.Lighting = directLight                                           // Lighting pass
	result.ReflectedDiffuse = globalIllumination.ReflectedDiffuse           // Reflected diffuse pass
	result.ReflectedLighting = globalIllumination.ReflectedLighting         // Reflected lighting pass
	result.Side = ctx.Scene.Camera.GetForwardDirection().Dot(rayNormal) > 0 // Side pass
	result.Normal = rayNormal                                               // Normal pass
	result.Depth = ray.RayLength                                            // Depth pass
	result.Specular = raySpecular                                           // Specular pass
	result.Glossiness = rayGlossiness                                       // Glossiness pass
	result.Diffuse.IMul(globalIllumination.Diffuse)
	return result
}

// Direct Illumination

func TraceDirectIlluminationRay(origin Maths.Vector3, normal Maths.Vector3, hitTriangle *Triangle, ctx *RenderContext) Color.Color {
	lightColor := Color.MakeColor(0, 0, 0)
	for i := 0; i < len(ctx.Scene.Lights); i++ {
		// Direct illumination without shadow smoothing (penumbra)
		light := ctx.Scene.Lights[i]
		lColor := light.GetLight(origin, normal, hitTriangle, ctx)
		lightColor.IUAdd(lColor)
	}
	return lightColor
}

// Indirect illumination

func TraceIndirectIllumination(reflectedRayDirection Maths.Vector3, previousRay *RayIntersection, glossiness float32, ctx *RenderContext, recursionDepth int) RayResult {
	roughness := 1 - glossiness
	rayCount := ctx.RenderSettings.IndirectIlluminationSampleCount
	reflRayCount := 0
	result := MakeEmptyResult()
	if !ctx.RenderSettings.UseIndirectIllumination {
		return result
	}
	norm := previousRay.GetHitNormal()
	if ctx.RenderSettings.OptimizeIndirectIlluminationRayCount {
		rayCount = int(float32(rayCount) * roughness)
	}
	for i := 0; i < rayCount; i++ {
		rayDirection := Maths.RandomPointOnHemisphere(norm, 1)
		ray := TraceRay(previousRay.HitPosition, rayDirection, ctx, previousRay.HitTriangle)
		if ray.Hit {
			rayUv := ray.GetHitUv()
			rayDiffuse := ray.GetHitMaterial().GetDiffuse(rayUv)
			raySpecular := ray.GetHitMaterial().GetSpecular(rayUv)
			directLighting := TraceDirectIlluminationRay(ray.HitPosition, ray.GetHitNormal(), ray.HitTriangle, ctx)
			result.Diffuse.IUAdd(rayDiffuse)
			result.Specular.IUAdd(raySpecular)
			result.Lighting.IUAdd(directLighting)
			if -rayDirection.Dot(reflectedRayDirection) > glossiness {
				result.ReflectedDiffuse.IUAdd(rayDiffuse)
				result.ReflectedLighting.IUAdd(directLighting)
				reflRayCount++
			}
			if recursionDepth < ctx.RenderSettings.IndirectIlluminationDepth {
				reflDir := rayDirection.Reflect(ray.GetHitNormal())
				rayGlossiness := ray.GetHitMaterial().Glossiness
				nextRay := TraceIndirectIllumination(reflDir, &ray, rayGlossiness, ctx, recursionDepth+1)
				result.ReflectedDiffuse.IMix(nextRay.Diffuse, rayGlossiness)
				result.ReflectedLighting.IMix(nextRay.Lighting, rayGlossiness)
			}
		}
	}
	// Adding some additional reflection rays to reduce noise
	// If glossiness is high enough, fewer rays will be accepted for reflections (they will be outside reflection cone)
	// This can lead to complete lack of rays, so image will be very noisy
	if glossiness > 0.75 {
		for i := 0; i < int(float32(rayCount/10)*roughness); i++ { // We don't need many rays, 1/10 of total ray count is enough
			rayDirection := Maths.RandomPointOnHemisphere(reflectedRayDirection, glossiness)
			ray := TraceRay(previousRay.HitPosition, rayDirection, ctx, previousRay.HitTriangle)
			if ray.Hit {
				uv := ray.GetHitUv()
				rLight := TraceDirectIlluminationRay(previousRay.HitPosition, norm, previousRay.HitTriangle, ctx)
				result.ReflectedDiffuse.IUAdd(ray.GetHitMaterial().GetDiffuse(uv))
				result.ReflectedLighting.IUAdd(rLight)
				reflRayCount++
			}
		}
	}
	result.Diffuse.IDivF(float32(rayCount))
	result.ReflectedDiffuse.IDivF(float32(rayCount))
	result.Lighting.IDivF(float32(rayCount))
	result.ReflectedLighting.IDivF(float32(rayCount))
	result.Specular.IDivF(float32(rayCount))
	return result
}
