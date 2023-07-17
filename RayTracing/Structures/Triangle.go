package Structures

import "RayTracer/Maths"

type Triangle struct {
	V1             *Vertex
	V2             *Vertex
	V3             *Vertex
	TriangleNormal Maths.Vector3
	Material       *Material
}

func (tri Triangle) CalcFirstEdge() Maths.Vector3 {
	return tri.V2.Position.Sub(tri.V1.Position)
}

func (tri Triangle) CalcSecondEdge() Maths.Vector3 {
	return tri.V3.Position.Sub(tri.V1.Position)
}

func (tri Triangle) Moved(offset Maths.Vector3) Triangle {
	return Triangle{tri.V1.Moved(offset), tri.V2.Moved(offset), tri.V3.Moved(offset), tri.TriangleNormal, tri.Material}
}

func MakeTriangle(v1 *Vertex, v2 *Vertex, v3 *Vertex, material *Material) Triangle {
	normal := v2.Position.Sub(v1.Position).Cross(v3.Position.Sub(v1.Position)).Normalized()
	return Triangle{v1, v2, v3, normal, material}
}

// Comparing

func AreEqualTriangles(tri1 *Triangle, tri2 *Triangle) bool {
	return Maths.AreEqualVectors(tri1.V1.Position, tri2.V1.Position) &&
		Maths.AreEqualVectors(tri1.V2.Position, tri2.V2.Position) &&
		Maths.AreEqualVectors(tri1.V3.Position, tri2.V3.Position)
}

// Getter methods

func (tri Triangle) GetUv(barycentricCoordinates Maths.Vector3) Maths.Vector2 {
	return Maths.TriInterpolate2(tri.V1.UV, tri.V2.UV, tri.V3.UV, barycentricCoordinates)
}

func (tri Triangle) GetSmoothNormal(barycentricCoordinates Maths.Vector3) Maths.Vector3 {
	return Maths.TriInterpolate(tri.V1.VertexNormal, tri.V2.VertexNormal, tri.V2.VertexNormal, barycentricCoordinates)
}

// Matrix multiplication

func (tri Triangle) MatMul(matrix *Maths.Mat3) Triangle {
	return MakeTriangle(tri.V1.MatMul(matrix), tri.V2.MatMul(matrix), tri.V3.MatMul(matrix), tri.Material)
}
