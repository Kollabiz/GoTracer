package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"fmt"
	"image"
	"image/png"
	"os"
)

func DrawColorArrayToImage(imgName string, colors [][]Color.Color) {
	img := image.NewRGBA(image.Rect(0, 0, len(colors), len(colors[0])))
	for x := 0; x < len(colors); x++ {
		for y := 0; y < len(colors[0]); y++ {
			img.Set(x, y, colors[x][y].ToImageColor())
		}
	}
	fl, err := os.Create(imgName)
	if err != nil {
		panic(fmt.Sprintf("cannot open or create file (%s)", imgName))
	}
	if err = png.Encode(fl, img); err != nil {
		panic(err)
	}
}

func Vector2Color(arr [][]Maths.Vector3) [][]Color.Color {
	conv := Maths.MakeColor2DArray(len(arr), len(arr[0]))
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			v := arr[x][y].Add(Maths.Vector3{1, 1, 1}).DivF(2)
			conv[x][y] = Color.Color{R: v.X, G: v.Y, B: v.Z}
		}
	}
	return conv
}

func Bool2Color(arr [][]bool) [][]Color.Color {
	conv := Maths.MakeColor2DArray(len(arr), len(arr[0]))
	tCl, fCl := Color.MakeColor(0, 0, 200), Color.MakeColor(200, 0, 0)
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			if arr[x][y] {
				conv[x][y] = tCl
			} else {
				conv[x][y] = fCl
			}
		}
	}
	return conv
}

func Float2Color(arr [][]float32, depth float32) [][]Color.Color {
	conv := Maths.MakeColor2DArray(len(arr), len(arr[0]))
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			clr := arr[x][y] / depth
			conv[x][y] = Color.Color{R: clr, G: clr, B: clr}
		}
	}
	return conv
}

func ClearVectorArray(arr [][]Maths.Vector3) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			arr[x][y] = Maths.ZeroVector3()
		}
	}
}

func ClearFloatArray(arr [][]float32) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			arr[x][y] = 0
		}
	}
}

func ClearBoolArray(arr [][]bool) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			arr[x][y] = false
		}
	}
}

func ClearColorArray(arr [][]Color.Color) {
	for x := 0; x < len(arr); x++ {
		for y := 0; y < len(arr[0]); y++ {
			arr[x][y] = Color.Color{R: 0, G: 0, B: 0}
		}
	}
}
