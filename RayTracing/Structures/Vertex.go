package Structures

import "RayTracer/Maths"

type Vertex struct {
	Position     Maths.Vector3
	UV           Maths.Vector2
	VertexNormal Maths.Vector3
}

func (vert *Vertex) Moved(offset Maths.Vector3) *Vertex {
	return &Vertex{vert.Position.Add(offset), vert.UV, vert.VertexNormal}
}

func (vert *Vertex) MatMul(matrix *Maths.Mat3) *Vertex {
	r := new(Vertex)
	r.Position = vert.Position.MatMul(matrix)
	r.VertexNormal = vert.VertexNormal.MatMul(matrix)
	r.UV = vert.UV
	return r
}
