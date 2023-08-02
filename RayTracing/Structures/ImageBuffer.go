package Structures

import (
	"RayTracer/Color"
)

type ImageBuffer struct {
	data [][]Color.Color
}

func NewImageBuffer(width int, height int) *ImageBuffer {
	buff := new(ImageBuffer)
	buff.data = make([][]Color.Color, width)
	for i := 0; i < width; i++ {
		buff.data[i] = make([]Color.Color, height)
	}
	return buff
}

func (buff *ImageBuffer) Set(x int, y int, c Color.Color) {
	buff.data[x][y] = c
}

func (buff *ImageBuffer) Add(x int, y int, c Color.Color) {
	buff.data[x][y].IUAdd(c)
}

func (buff *ImageBuffer) Get(x int, y int) Color.Color {
	return buff.data[x][y]
}

func (buff *ImageBuffer) GetWidth() int {
	return len(buff.data)
}

func (buff *ImageBuffer) GetHeight() int {
	return len(buff.data[0])
}

func (buff *ImageBuffer) DivAll(f float32) {
	for x := 0; x < buff.GetWidth(); x++ {
		for y := 0; y < buff.GetHeight(); y++ {
			buff.data[x][y].IDivF(f)
		}
	}
}

func (buff *ImageBuffer) NormColor() {
	for x := 0; x < buff.GetWidth(); x++ {
		for y := 0; y < buff.GetHeight(); y++ {
			buff.data[x][y].Normalize()
		}
	}
}

func (buff *ImageBuffer) SaveToFile(fileName string) {
	DrawColorArrayToImage(fileName, buff.data)
}
