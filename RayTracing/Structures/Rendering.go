package Structures

import (
	"RayTracer/Maths"
	"sync"
)

func GenerateViewportGrid(imgWidth int, imgHeight int, viewportCorners [4]Maths.Vector3, lensSize Maths.Vector2) [][]Maths.Vector3 {
	points := make([][]Maths.Vector3, imgWidth)
	for i := 0; i < imgHeight; i++ {
		points[i] = make([]Maths.Vector3, imgHeight)
	}
	stepSizeX := lensSize.X / float32(imgWidth)
	stepSizeY := lensSize.Y / float32(imgHeight)
	for x := 0; x < imgWidth; x++ {
		point := viewportCorners[0].Lerp(viewportCorners[1], float32(x)*stepSizeX)
		point2 := viewportCorners[2].Lerp(viewportCorners[3], float32(x)*stepSizeX)
		for y := 0; y < imgHeight; y++ {
			points[x][y] = point.Lerp(point2, float32(y)*stepSizeY)
		}
	}
	return points
}

func RenderTile(points [][]Maths.Vector3, focalPoint Maths.Vector3, startFragment Maths.Vector2, endFragment Maths.Vector2, imgBuff *ImageBuffer, ctx *RenderContext, mutex *sync.Mutex, rSamples *int) {
	for i := 0; i < ctx.RenderSettings.ProgressiveRenderingPassQuantity; i++ {
		for fragmentX := int(startFragment.X); fragmentX < int(endFragment.X); fragmentX++ {
			if fragmentX >= imgBuff.GetWidth() {
				continue
			}
			for fragmentY := int(startFragment.Y); fragmentY < int(endFragment.Y); fragmentY++ {
				if fragmentY >= imgBuff.GetHeight() {
					continue
				}
				point := points[fragmentX][fragmentY].Add(ctx.Scene.Camera.Position)
				rayDirection := point.Sub(focalPoint).Normalized()
				ray := TraceStartRay(point, rayDirection, ctx)
				mutex.Lock()
				imgBuff.Add(fragmentX, fragmentY, ray)
				ctx.Scene.RenderedPixels++
				*rSamples++
				mutex.Unlock()
			}
		}
	}
	ctx.Scene.RenderedTiles++
}
