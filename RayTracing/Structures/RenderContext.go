package Structures

import "RayTracer/Color"

type RenderContext struct {
	Scene           *Scene
	ImageWidth      int
	ImageHeight     int
	MaxLightBounces int
	BackgroundColor Color.Color

	// Render settings

	MinLightIntensity float32
	RenderSettings    *RenderSettings
}

func NewRenderContext(scene *Scene, imgWidth int, imgHeight int, maxLightBounces int, backgroundColor Color.Color,
	minLightIntensity float32, igSampleCount int, igDepth int) *RenderContext {
	ctx := new(RenderContext)
	ctx.Scene = scene
	ctx.ImageWidth = imgWidth
	ctx.ImageHeight = imgHeight
	ctx.MaxLightBounces = maxLightBounces
	ctx.BackgroundColor = backgroundColor
	ctx.MinLightIntensity = minLightIntensity
	ctx.RenderSettings = &RenderSettings{
		UseIndirectIllumination:              true,
		OptimizeIndirectIlluminationRayCount: true,
		IndirectIlluminationSampleCount:      igSampleCount,
		IndirectIlluminationDepth:            igDepth,
	}
	return ctx
}
