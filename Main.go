package main

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"RayTracer/ObjParser"
	"RayTracer/RayTracing/Structures"
	"RayTracer/RayTracing/Structures/LightSources"
)

func main() {
	scene := Structures.NewScene(512, 512)
	scene.Camera.SetRotation(Maths.Vector3{
		X: 0,
		Y: 0,
		Z: 0,
	})
	scene.Camera.Position = Maths.Vector3{0, 0, 10}
	meshes, _ := ObjParser.ParseObjFile("Examples\\high_poly.obj")
	pointLight := LightSources.PointLight{
		Color:     Color.Color{1, 1, 1},
		Position:  Maths.MakeVector3(0.2, 0, 2),
		Radius:    0.1,
		Intensity: 1.7,
	}
	for i := 0; i < len(meshes); i++ {
		meshes[i].SetRotation(Maths.Vector3{X: -24})
	}
	scene.AddLight(&pointLight)
	scene.AddMeshes(meshes)
	scene.RenderScene()
}
