package Structures

import (
	"RayTracer/Color"
	"RayTracer/Maths"
	"image"
	"log"
	"os"
)

type Texture2 struct {
	pixels [][]Color.Color
	Width  int
	Height int
}

func (texture *Texture2) GetPixel(x int, y int) Color.Color {
	return texture.pixels[x][y]
}

func (texture *Texture2) GetPixelFromUV(uv Maths.Vector2) Color.Color {
	if texture == nil { // Texture is not loaded
		return Color.Color{R: 1, G: 1, B: 1}
	}
	return texture.pixels[int(uv.X*float32(texture.Width))][int(uv.Y*float32(texture.Height))]
}

func Texture2FromFile(filePath string) *Texture2 {
	fl, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	im, _, err := image.Decode(fl)
	if err != nil {
		log.Fatal(err)
	}

	tex := new(Texture2)
	imWidth := im.Bounds().Max.X
	imHeight := im.Bounds().Min.Y
	tex.Width = imWidth
	tex.Height = imHeight

	// Allocating memory for texture
	tex.pixels = make([][]Color.Color, tex.Width)
	for i := 0; i < tex.Width; i++ {
		tex.pixels[i] = make([]Color.Color, tex.Height)
	}

	// Writing to allocated texture
	for x := 0; x < tex.Width; x++ {
		for y := 0; y < tex.Height; y++ {
			tex.pixels[x][y] = Color.FromImageColor(im.At(x, y))
		}
	}

	return tex
}
