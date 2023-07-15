package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"fmt"
	"log"
	"os"
)

type RenderPasses struct {
	Diffuse  [][]Color.Color
	Lighting [][]Color.Color
	// Reflection data
	ReflectedDiffuse  [][]Color.Color
	ReflectedLighting [][]Color.Color
	// Debug passes
	Side                   [][]bool
	Normal                 [][]Maths.Vector3
	Depth                  [][]float32
	Specular               [][]Color.Color
	Glossiness             [][]float32
	AccumulatedPassesCount int
	Width                  int
	Height                 int
}

func NewRenderPasses(width int, height int) *RenderPasses {
	passes := new(RenderPasses)
	passes.Diffuse = Maths.MakeColor2DArray(width, height)
	passes.Lighting = Maths.MakeColor2DArray(width, height)
	passes.ReflectedDiffuse = Maths.MakeColor2DArray(width, height)
	passes.ReflectedLighting = Maths.MakeColor2DArray(width, height)
	passes.Side = Maths.MakeBool2DArray(width, height)
	passes.Normal = Maths.MakeVector2DArray(width, height)
	passes.Depth = Maths.MakeFloat2DArray(width, height)
	passes.Specular = Maths.MakeColor2DArray(width, height)
	passes.Glossiness = Maths.MakeFloat2DArray(width, height)
	passes.AccumulatedPassesCount = 0
	passes.Width = width
	passes.Height = height
	return passes
}

func (passes *RenderPasses) DrawRay(x int, y int, ray RayResult) {
	passes.Diffuse[x][y].IUAdd(ray.Diffuse)
	passes.Lighting[x][y].IUAdd(ray.Lighting)
	passes.ReflectedDiffuse[x][y].IUAdd(ray.ReflectedDiffuse)
	passes.ReflectedLighting[x][y].IUAdd(ray.ReflectedLighting)
	passes.Side[x][y] = ray.Side
	passes.Normal[x][y] = ray.Normal
	passes.Depth[x][y] = ray.Depth
	passes.Specular[x][y].IUAdd(ray.Specular)
	passes.Glossiness[x][y] = ray.Glossiness
}

func (passes *RenderPasses) GetDiffuse(x int, y int) Color.Color {
	return passes.Diffuse[x][y].DivF(float32(passes.AccumulatedPassesCount))
}

func (passes *RenderPasses) GetLighting(x int, y int) Color.Color {
	return passes.Lighting[x][y].DivF(float32(passes.AccumulatedPassesCount))
}

func (passes *RenderPasses) GetReflectedDiffuse(x int, y int) Color.Color {
	return passes.ReflectedDiffuse[x][y].DivF(float32(passes.AccumulatedPassesCount))
}

func (passes *RenderPasses) GetReflectedLighting(x int, y int) Color.Color {
	return passes.ReflectedLighting[x][y].DivF(float32(passes.AccumulatedPassesCount))
}

func (passes *RenderPasses) GetSpecular(x int, y int) Color.Color {
	return passes.Specular[x][y].DivF(float32(passes.AccumulatedPassesCount))
}

func (passes *RenderPasses) Combined() [][]Color.Color {
	if passes.AccumulatedPassesCount == 0 {
		log.Fatal("Trying to combine passes before render")
	}
	combined := Maths.MakeColor2DArray(passes.Width, passes.Height)
	for x := 0; x < passes.Width; x++ {
		for y := 0; y < passes.Height; y++ {
			litDiffuse := passes.GetDiffuse(x, y).Mul(passes.GetLighting(x, y))
			litReflection := passes.GetReflectedDiffuse(x, y).Mul(passes.GetReflectedLighting(x, y)).Mul(passes.GetSpecular(x, y))
			mixed := litDiffuse.Mix(litReflection, passes.Glossiness[x][y])
			combined[x][y] = mixed
		}
	}
	return combined
}

func (passes *RenderPasses) ClearPasses() {
	ClearColorArray(passes.Diffuse)
	ClearColorArray(passes.Lighting)
	ClearColorArray(passes.ReflectedDiffuse)
	ClearColorArray(passes.ReflectedLighting)
	ClearBoolArray(passes.Side)
	ClearVectorArray(passes.Normal)
	ClearFloatArray(passes.Depth)
	ClearColorArray(passes.Specular)
	ClearFloatArray(passes.Glossiness)
	passes.AccumulatedPassesCount = 0
}

func (passes *RenderPasses) ClearDepth() {
	ClearFloatArray(passes.Depth)
}

func (passes *RenderPasses) SavePassesToImages(saveDirectory string, ctx *RenderContext) {
	if err := os.MkdirAll(saveDirectory, os.ModePerm); err != nil {
		log.Println("Directory already exists")
	}
	// Combined pass
	cmb := passes.Combined()
	DrawColorArrayToImage(fmt.Sprintf("%s\\combined.png", saveDirectory), cmb)
	if !ctx.RenderSettings.DumpDebugPasses {
		return
	}
	// Diffuse pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\diffuse_pass.png", saveDirectory), passes.Diffuse)
	// Lighting pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\lighting_pass.png", saveDirectory), passes.Lighting)
	// Reflected diffuse pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\refl_diffuse_pass.png", saveDirectory), passes.ReflectedDiffuse)
	// Reflected lighting pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\refl_lighting_pass.png", saveDirectory), passes.ReflectedLighting)
	// Side pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\side_pass.png", saveDirectory), Bool2Color(passes.Side))
	// Normal pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\normal_pass.png", saveDirectory), Vector2Color(passes.Normal))
	// Depth pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\depth_pass.png", saveDirectory), Float2Color(passes.Depth, 50))
	// Specular pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\specular_pass.png", saveDirectory), passes.Specular)
	// Glossiness pass
	DrawColorArrayToImage(fmt.Sprintf("%s\\glossiness_pass.png", saveDirectory), Float2Color(passes.Glossiness, 1))
}
