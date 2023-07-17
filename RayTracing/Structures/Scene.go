package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"fmt"
	"math"
	"time"
)

type Scene struct {
	Meshes         []Mesh
	Lights         []ILightSource
	Camera         Camera
	RenderedTiles  int
	RenderedPixels int
	RenderContext  *RenderContext
	passes         *RenderPasses
}

func NewScene(viewportWidth int, viewportHeight int) *Scene {
	sc := new(Scene)
	sc.RenderContext = new(RenderContext)
	sc.RenderContext.RenderSettings = &RenderSettings{
		UseIndirectIllumination:              false,
		OptimizeIndirectIlluminationRayCount: false,
		IndirectIlluminationSampleCount:      16,
		IndirectIlluminationDepth:            1,
		DumpDebugPasses:                      true,
		ProgressiveRenderingPassQuantity:     32,
		UseTiling:                            true,
		TileWidth:                            16,
		TileHeight:                           16,
	}
	sc.RenderContext.Scene = sc
	sc.RenderedTiles = 0
	sc.RenderContext.RenderSettings.ShadowSamples = 1
	sc.RenderContext.MinLightIntensity = 0.3
	sc.RenderContext.BackgroundColor = Color.Color{150, 150, 150}
	sc.RenderContext.ImageWidth = viewportWidth
	sc.RenderContext.ImageHeight = viewportHeight
	sc.passes = NewRenderPasses(viewportWidth, viewportHeight)
	sc.Camera = MakeCamera(Maths.ZeroVector3(), Maths.ZeroVector3(), Maths.LensSizeFromAspectRatio(viewportWidth, viewportHeight), 8)
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
	pixelAmount := scene.RenderContext.ImageWidth * scene.RenderContext.ImageHeight * scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity
	points := GenerateViewportGrid(scene.RenderContext.ImageWidth, scene.RenderContext.ImageHeight, scene.Camera.GetViewportPlaneCorners(), scene.Camera.LensSize)
	focalPoint := Maths.Vector3{0, 0, scene.Camera.FocalLength}.Add(scene.Camera.Position)
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
	scene.passes.AccumulatedPassesCount = scene.RenderContext.RenderSettings.ProgressiveRenderingPassQuantity
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
			go RenderTile(points, focalPoint, tileStart, tileEnd, scene.passes, scene.RenderContext)
		}
	}
	for scene.RenderedTiles < tileCountY*tileCountX {
		//fmt.Printf(" Rendering progress: %d%% (%d/%d tiles | %d/%d pixels)   				            \r",
		//	int(float32(scene.RenderedPixels)/float32(pixelAmount)*100),
		//	scene.RenderedTiles,
		//	tileCountY*tileCountX,
		//	scene.RenderedPixels,
		//	pixelAmount,
		//)
		time.Sleep(10 * time.Millisecond)
	}
	scene.passes.SavePassesToImages("passes", scene.RenderContext)
	scene.passes.ClearPasses()
}
