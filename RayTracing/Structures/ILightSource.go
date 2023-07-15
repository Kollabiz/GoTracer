package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
)

type ILightSource interface {
	GetLight(position Maths.Vector3, normal Maths.Vector3, excludeTri *Triangle, ctx *RenderContext) Color.Color
}
