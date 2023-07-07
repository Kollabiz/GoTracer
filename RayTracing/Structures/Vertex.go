package Structures

import "RayTracer/Maths"

type Vertex struct {
	Position     Maths.Vector3
	UV           Maths.Vector3
	VertexNormal Maths.Vector3
}

func (vert *Vertex) Moved(offset Maths.Vector3) *Vertex {
	return &Vertex{vert.Position.Add(offset), vert.UV, vert.VertexNormal}
}
