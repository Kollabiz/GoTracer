package Structures

import (
	"RayTracer/Maths"
	"fmt"
	"math"
	"sync"
	"time"
)

type Scene struct {
	Meshes         []Mesh
	Lights         []ILightSource
	Camera         Camera
	RenderedTiles  int
	RenderedPixels int
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
	}
}

func (scene *Scene) RenderScene() {
	scene.bakeMeshTransforms()
	tileCountY := int(math.Ceil(float64(scene.RenderContext.ImageHeight) / float64(scene.RenderContext.RenderSettings.TileHeight)))
	tileCountX := int(math.Ceil(float64(scene.RenderContext.ImageWidth) / float64(scene.RenderContext.RenderSettings.TileWidth)))
	pixelAmount := scene.RenderContext.ImageWidth * scene.RenderContext.ImageHeight
	points := GenerateViewportGrid(scene.RenderContext.ImageWidth, scene.RenderContext.ImageHeight, scene.Camera.GetViewportPlaneCorners(), scene.Camera.LensSize)
	focalPoint := Maths.Vector3{0, 0, scene.Camera.FocalLength}.Add(scene.Camera.Position)
	var renderedSamples int
	var mutex sync.Mutex
	fmt.Printf(
		`------------DEBUG--------------
Image width: %d
Image height: %d
Tiles X: %d
Tiles Y: %d
Tile width: %d
Tile height: %d
Pixel amount: %d
Camera lens size: (%f; %f)
Mesh count: %d
`,
		scene.RenderContext.ImageWidth,
		scene.RenderContext.ImageHeight,
		tileCountX,
		tileCountY,
		scene.RenderContext.RenderSettings.TileWidth,
		scene.RenderContext.RenderSettings.TileHeight,
		pixelAmount,
		scene.Camera.LensSize.X,
		scene.Camera.LensSize.Y,
		len(scene.Meshes),
	)
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
			go RenderTile(points, focalPoint, tileStart, tileEnd, scene.imageBuffer, scene.RenderContext, &mutex, &renderedSamples)
		}
	}
	for scene.RenderedTiles < tileCountY*tileCountX {
		fmt.Printf(" Rendering progress: %d%% (%d/%d tiles | %d/%d pixels)   				            \r",
			int(float32(scene.RenderedPixels)/float32(pixelAmount*scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity)*100),
			scene.RenderedTiles,
			tileCountY*tileCountX,
			scene.RenderedPixels,
			pixelAmount*scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity,
		)
		time.Sleep(10 * time.Millisecond)
	}
	scene.imageBuffer.DivAll(float32(scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity))
	scene.imageBuffer.SaveToFile("result.png")
}
