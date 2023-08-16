package Structures

import (
	"RayTracer/Maths"
	"math"
)

type Mesh struct {
	rotation       Maths.Vector3
	Position       Maths.Vector3
	scale          Maths.Vector3
	rotationMatrix *Maths.Mat3
	scaleMatrix    *Maths.Mat3
	MeshName       string
	Triangles      []Triangle
	Volume         *BoxVolume
	TriangleTree   *BVHNode
	// Transform cache
	transformed []Triangle
}

func MakeMesh(meshName string) Mesh {
	mesh := Mesh{}
	mesh.MeshName = meshName
	mesh.rotationMatrix = Maths.Mat3Identity()
	mesh.scaleMatrix = Maths.Mat3Identity()
	return mesh
}

func (mesh *Mesh) GetTransformed() []Triangle {
	return mesh.transformed
}

func (mesh *Mesh) BakeTransform() {
	mesh.transformed = make([]Triangle, len(mesh.Triangles))
	for i := 0; i < len(mesh.Triangles); i++ {
		mesh.transformed[i] = mesh.Triangles[i].MatMul(mesh.rotationMatrix).MatMul(mesh.scaleMatrix).Moved(mesh.Position)
	}
}

// Setters

func (mesh *Mesh) SetRotation(rotation Maths.Vector3) {
	mesh.rotation = rotation
	mesh.rotationMatrix = Maths.Mat3FromEulerAngles(rotation)
}

func (mesh *Mesh) SetScale(scale Maths.Vector3) {
	mesh.scale = scale
	mesh.scaleMatrix = Maths.Mat3Scale(scale)
}

// Getters

func (mesh *Mesh) GetRotation() Maths.Vector3 {
	return mesh.rotation
}

func (mesh *Mesh) GetScale() Maths.Vector3 {
	return mesh.scale
}

func (mesh *Mesh) GetTransformMatrix() *Maths.Mat3 {
	return mesh.rotationMatrix.MatMul(mesh.scaleMatrix)
}

func (mesh *Mesh) GetBoundaries() (Maths.Vector3, Maths.Vector3) {
	var min, max = Maths.InfiniteVector3(1), Maths.InfiniteVector3(-1)
	for i := 0; i < len(mesh.transformed); i++ {
		tri := mesh.transformed[i]
		minX := float32(math.Min(float64(tri.V1.Position.X), math.Min(float64(tri.V2.Position.X), float64(tri.V3.Position.X))))
		minY := float32(math.Min(float64(tri.V1.Position.Y), math.Min(float64(tri.V2.Position.Y), float64(tri.V3.Position.Y))))
		minZ := float32(math.Min(float64(tri.V1.Position.Z), math.Min(float64(tri.V2.Position.Z), float64(tri.V3.Position.Z))))
		maxX := float32(math.Max(float64(tri.V1.Position.X), math.Max(float64(tri.V2.Position.X), float64(tri.V3.Position.X))))
		maxY := float32(math.Max(float64(tri.V1.Position.Y), math.Max(float64(tri.V2.Position.Y), float64(tri.V3.Position.Y))))
		maxZ := float32(math.Max(float64(tri.V1.Position.Z), math.Max(float64(tri.V2.Position.Z), float64(tri.V3.Position.Z))))
		if minX < min.X {
			min.X = minX
		}
		if minY < min.Y {
			min.Y = minY
		}
		if minZ < min.Z {
			min.Z = minZ
		}
		if maxX > max.X {
			max.X = maxX
		}
		if maxY > max.Y {
			max.Y = maxY
		}
		if max.Z > max.Z {
			max.Z = maxZ
		}
	}
	return min, max
}

func (mesh *Mesh) BuildMeshTree() {
	var nodes = make([]BVHNode, len(mesh.transformed))
	for i := 0; i < len(mesh.transformed); i++ {
		nodes[i] = *BVHNodeFromTri(&mesh.transformed[i])
	}
	for len(nodes) > 1 {
		node := nodes[0]
		var nearestNode = nodes[1]
		var nearestDist = nearestNode.GetMidPoint().Sub(node.GetMidPoint()).Length()
		var secondNodeIndex = 1
		for i := 2; i < len(nodes); i++ {
			if nodes[i].GetMidPoint().Sub(node.GetMidPoint()).Length() < nearestDist {
				nearestNode = nodes[i]
				nearestDist = nearestNode.GetMidPoint().Sub(node.GetMidPoint()).Length()
			}
		}
		joined := JoinBVHNodes(&nearestNode, &node)
		nodes = append(append(nodes[1:secondNodeIndex], nodes[secondNodeIndex+1:]...), *joined)
	}
	mesh.TriangleTree = &nodes[0]
}
