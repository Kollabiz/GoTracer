package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"fmt"
	"math"
	"time"
)

type Scene struct {
	Meshes        []Mesh
	Lights        []ILightSource
	Camera        Camera
	RenderedTiles int
	renderContext *RenderContext
	passes        *RenderPasses
}

func NewScene(viewportWidth int, viewportHeight int) *Scene {
	sc := new(Scene)
	sc.renderContext = new(RenderContext)
	sc.renderContext.RenderSettings = &RenderSettings{
		UseIndirectIllumination:              false,
		OptimizeIndirectIlluminationRayCount: false,
		IndirectIlluminationSampleCount:      16,
		IndirectIlluminationDepth:            1,
		DumpDebugPasses:                      true,
		ProgressiveRenderingPassQuantity:     1,
		UseTiling:                            true,
		TileWidth:                            32,
		TileHeight:                           32,
	}
	sc.renderContext.Scene = sc
	sc.RenderedTiles = 0
	sc.renderContext.ShadowIntensity = 0.2
	sc.renderContext.MinLightIntensity = 0.1
	sc.renderContext.BackgroundColor = Color.Color{150, 150, 150}
	sc.renderContext.ImageWidth = viewportWidth
	sc.renderContext.ImageHeight = viewportHeight
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
	pixelAmount := scene.renderContext.ImageWidth * scene.renderContext.ImageHeight
	scene.bakeMeshTransforms()
	tileCountX := int(math.Ceil(float64(scene.renderContext.ImageWidth) / float64(scene.renderContext.RenderSettings.TileWidth)))
	tileCountY := int(math.Ceil(float64(scene.renderContext.ImageHeight) / float64(scene.renderContext.RenderSettings.TileHeight)))
	points := GenerateViewportGrid(scene.renderContext.ImageWidth, scene.renderContext.ImageHeight, scene.Camera.GetViewportPlaneCorners(), scene.Camera.LensSize)
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
		scene.renderContext.ImageWidth,
		scene.renderContext.ImageHeight,
		tileCountX,
		tileCountY,
		scene.renderContext.RenderSettings.TileWidth,
		scene.renderContext.RenderSettings.TileHeight,
		pixelAmount,
		scene.Camera.LensSize.X,
		scene.Camera.LensSize.Y,
		len(scene.Meshes),
	)
	for i := 0; i < scene.renderContext.RenderSettings.ProgressiveRenderingPassQuantity; i++ {
		for tileX := 0; tileX < tileCountX; tileX++ {
			for tileY := 0; tileY < tileCountY; tileY++ {
				tileStart := Maths.Vector2{
					float32(tileX * scene.renderContext.RenderSettings.TileWidth),
					float32(tileY * scene.renderContext.RenderSettings.TileHeight),
				}
				tileEnd := Maths.Vector2{
					float32((tileX + 1) * scene.renderContext.RenderSettings.TileWidth),
					float32((tileY + 1) * scene.renderContext.RenderSettings.TileHeight),
				}
				go RenderTile(points, focalPoint, tileStart, tileEnd, scene.passes, scene.renderContext)
			}
		}
		scene.passes.AccumulatedPassesCount++
		for scene.RenderedTiles < tileCountY*tileCountX {
			fmt.Printf(" Rendering progress: %d%% (%d/%d tiles)   				            \r",
				pixelAmount,
				scene.RenderedTiles,
				tileCountY*tileCountX,
			)
			time.Sleep(10 * time.Millisecond)
		}
		scene.RenderedTiles = 0
	}
	scene.passes.SavePassesToImages("passes", scene.renderContext)
	scene.passes.ClearPasses()
}
