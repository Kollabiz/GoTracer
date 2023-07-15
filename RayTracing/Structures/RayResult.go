package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
)

type RayResult struct {
	Diffuse  Color.Color
	Lighting Color.Color
	// Reflection data
	ReflectedDiffuse  Color.Color
	ReflectedLighting Color.Color
	// Debug passes
	Side       bool
	Normal     Maths.Vector3
	Depth      float32
	Specular   Color.Color
	Glossiness float32
}

func MakeRayResult(diffuse Color.Color, lighting Color.Color, reflDiffuse Color.Color, reflLighting Color.Color, side bool, normal Maths.Vector3, depth float32, specular Color.Color, glossiness float32) RayResult {
	return RayResult{
		diffuse,
		lighting,
		reflDiffuse,
		reflLighting,
		side,
		normal,
		depth,
		specular,
		glossiness,
	}
}

func MakeEmptyResult() RayResult {
	return RayResult{
		Color.MakeColor(0, 0, 0),
		Color.MakeColor(0, 0, 0),
		Color.MakeColor(0, 0, 0),
		Color.MakeColor(0, 0, 0),
		false,
		Maths.ZeroVector3(),
		0,
		Color.MakeColor(0, 0, 0),
		0,
	}
}

func MakeNonHitResult() RayResult {
	return RayResult{}
}
