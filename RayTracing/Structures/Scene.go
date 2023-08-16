package Structures

import (
	"RayTracer/Maths"
	"fmt"
	"math"
	"sync"
)

type Scene struct {
	Meshes         []Mesh
	Lights         []ILightSource
	Camera         Camera
	RenderedTiles  int
	RenderedPixels int32
	RenderContext  *RenderContext
	imageBuffer    *ImageBuffer
}

func NewScene(ctx *RenderContext) *Scene {
	sc := new(Scene)
	sc.RenderContext = ctx
	sc.RenderContext.Scene = sc
	sc.RenderedTiles = 0
	sc.imageBuffer = NewImageBuffer(ctx.ImageWidth, ctx.ImageHeight)
	sc.Camera = MakeCamera(Maths.ZeroVector3(), Maths.ZeroVector3(), Maths.LensSizeFromAspectRatio(ctx.ImageWidth, ctx.ImageHeight), 8)
	return sc
}

func (scene *Scene) AddMesh(mesh Mesh) {
	scene.Meshes = append(scene.Meshes, mesh)
}

func (scene *Scene) AddMeshes(meshes []Mesh) {
	scene.Meshes = append(scene.Meshes, meshes...)
}

func (scene *Scene) RemoveMeshAt(i int) {
	scene.Meshes[i] = scene.Meshes[len(scene.Meshes)-1]
	scene.Meshes = scene.Meshes[:len(scene.Meshes)-1]
}

func (scene *Scene) RemoveMeshByName(meshName string) {
	for i := 0; i < len(scene.Meshes); i++ {
		if scene.Meshes[i].MeshName == meshName {
			scene.RemoveMeshAt(i)
			return
		}
	}
}

func (scene *Scene) RemoveMesh(mesh Mesh) {
	scene.RemoveMeshByName(mesh.MeshName)
}

func (scene *Scene) AddLight(light ILightSource) {
	scene.Lights = append(scene.Lights, light)
}

func (scene *Scene) AddLights(lights []ILightSource) {
	scene.Lights = append(scene.Lights, lights...)
}

func (scene *Scene) bakeMeshTransforms() {
	for i := 0; i < len(scene.Meshes); i++ {
		scene.Meshes[i].BakeTransform()
		scene.Meshes[i].BuildMeshTree()
		mMin, mMax := scene.Meshes[i].GetBoundaries()
		scene.Meshes[i].Volume = &BoxVolume{
			Min: mMin,
			Max: mMax,
		}
	}
}

func (scene *Scene) RenderScene() {
	scene.bakeMeshTransforms()
	fmt.Println("Building BVH tree")
	fmt.Printf("Done (%d meshes)\n", len(scene.Meshes))
	tileCountY := int(math.Ceil(float64(scene.RenderContext.ImageHeight) / float64(scene.RenderContext.RenderSettings.TileHeight)))
	tileCountX := int(math.Ceil(float64(scene.RenderContext.ImageWidth) / float64(scene.RenderContext.RenderSettings.TileWidth)))
	points := GenerateViewportGrid(scene.RenderContext.ImageWidth, scene.RenderContext.ImageHeight, scene.Camera.GetViewportPlaneCorners(), scene.Camera.LensSize)
	focalPoint := Maths.Vector3{0, 0, scene.Camera.FocalLength}.Add(scene.Camera.Position)
	fmt.Printf("Starting rendering (%d threads)\n", tileCountX*tileCountY)
	var wGroup sync.WaitGroup
	totalPixels := scene.RenderContext.ImageWidth * scene.RenderContext.ImageHeight * scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity
	go scene.displayProgress(int32(totalPixels), int32(tileCountX*tileCountY), &wGroup)
	wGroup.Add(1)
	for tileX := 0; tileX < tileCountX; tileX++ {
		for tileY := 0; tileY < tileCountY; tileY++ {
			tileStart := Maths.Vector2{
				float32(tileX * scene.RenderContext.RenderSettings.TileWidth),
				float32(tileY * scene.RenderContext.RenderSettings.TileHeight),
			}
			tileEnd := Maths.Vector2{
				float32((tileX + 1) * scene.RenderContext.RenderSettings.TileWidth),
				float32((tileY + 1) * scene.RenderContext.RenderSettings.TileHeight),
			}
			RenderTile(points, focalPoint, tileStart, tileEnd, scene.imageBuffer, scene.RenderContext)
		}
	}
	wGroup.Wait()
	fmt.Println("Done")
	fmt.Println("Dumping render result")
	scene.imageBuffer.DivAll(float32(scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity))
	scene.imageBuffer.SaveToFile("result.png")
	fmt.Println("Rendering done")
}

func (scene *Scene) displayProgress(totalPixels int32, totalTiles int32, waitGroup *sync.WaitGroup) {
	for scene.RenderedPixels < totalPixels {
		fmt.Printf(" Rendering progress: %d%% (%d/%d tiles, %d/%d rays)\r",
			int(float32(scene.RenderedPixels)/float32(totalPixels)*100),
			scene.RenderedTiles,
			totalTiles,
			scene.RenderedPixels,
			totalPixels,
		)
	}
	waitGroup.Done()
}
