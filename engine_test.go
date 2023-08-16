package main

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"RayTracer/ObjParser"
	"RayTracer/RayTracing/Structures"
	"RayTracer/RayTracing/Structures/LightSources"
	"testing"
)

func TestRender(t *testing.T) {
	rc := &Structures.RenderContext{
		ImageWidth:        256,
		ImageHeight:       256,
		BackgroundColor:   Color.Color{R: 200, G: 200, B: 200},
		MinLightIntensity: 0.2,
		RenderSettings: &Structures.RenderSettings{
			UseIndirectIllumination:              true,
			OptimizeIndirectIlluminationRayCount: false,
			IndirectIlluminationSampleCount:      16,
			IndirectIlluminationDepth:            1,
			ProgressiveRenderingPassQuantity:     1,
			UseTiling:                            true,
			TileWidth:                            256,
			TileHeight:                           256,
			ShadowSamples:                        1,
		},
	}
	scene := Structures.NewScene(rc)
	scene.Camera.SetRotation(Maths.Vector3{
		X: 0,
		Y: 0,
		Z: 0,
	})
	scene.Camera.Position = Maths.Vector3{0, 0, 10}
	meshes, _ := ObjParser.ParseObjFile("Examples\\high_poly.obj")
	pointLight := LightSources.DirectionalLight{
		Color:     Color.Color{1, 1, 1},
		Direction: Maths.Vector3{0, 0.2, 1},
		Intensity: 1,
	}
	scene.AddLight(&pointLight)
	scene.AddMeshes(meshes)
	scene.RenderScene()
}
