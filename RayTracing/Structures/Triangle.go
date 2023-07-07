package Structures

import "RayTracer/Maths"

type Triangle struct {
	V1             *Vertex
	V2             *Vertex
	V3             *Vertex
	TriangleNormal Maths.Vector3
}

func (tri Triangle) CalcFirstEdge() Maths.Vector3 {
	return tri.V2.Position.Sub(tri.V1.Position)
}

func (tri Triangle) CalcSecondEdge() Maths.Vector3 {
	return tri.V3.Position.Sub(tri.V1.Position)
}

func (tri Triangle) Moved(offset Maths.Vector3) Triangle {
	return Triangle{tri.V1.Moved(offset), tri.V2.Moved(offset), tri.V3.Moved(offset), tri.TriangleNormal}
}

func MakeTriangle(v1 *Vertex, v2 *Vertex, v3 *Vertex) Triangle {
	normal := v2.Position.Sub(v1.Position).Cross(v3.Position.Sub(v1.Position))
	return Triangle{v1, v2, v3, normal}
}
