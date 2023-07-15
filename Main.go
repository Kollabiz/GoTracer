package main

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"RayTracer/ObjParser"
	"RayTracer/RayTracing/Structures"
	"RayTracer/RayTracing/Structures/LightSources"
)

func main() {
	scene := Structures.NewScene(128, 128)
	scene.Camera.SetRotation(Maths.Vector3{
		X: 0,
		Y: 0,
		Z: 0,
	})
	scene.Camera.Position = Maths.Vector3{0, 0, 10}
	meshes, _ := ObjParser.ParseObjFile("Examples\\high_poly.obj")
	light := LightSources.DirectionalLight{
		Color:     Color.Color{1, 1, 1},
		Direction: Maths.Vector3{0.5, 0.5, 0.5}.Normalized(),
		Intensity: 1,
	}
	scene.AddLight(&light)
	scene.AddMeshes(meshes)
	scene.RenderScene()
}
