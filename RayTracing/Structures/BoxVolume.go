package Structures

import "RayTracer/Maths"

type BoxVolume struct {
	Min Maths.Vector3
	Max Maths.Vector3
}

func (box *BoxVolume) GetBoundSize() Maths.Vector3 {
	return box.Max.Sub(box.Min).Abs()
}
