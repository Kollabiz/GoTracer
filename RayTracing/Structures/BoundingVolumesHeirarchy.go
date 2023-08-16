package Structures

import (
	"RayTracer/Maths"
	"math"
)

type BVHNode struct {
	Triangle    *Triangle
	BoundingBox *BoxVolume
	Children    []BVHNode
}

func NewBVHNode() *BVHNode {
	return new(BVHNode)
}

func JoinBVHNodes(node1 *BVHNode, node2 *BVHNode) *BVHNode {
	minX := float32(math.Min(float64(node1.BoundingBox.Min.X), float64(node2.BoundingBox.Min.X)))
	minY := float32(math.Min(float64(node1.BoundingBox.Min.Y), float64(node2.BoundingBox.Min.Y)))
	minZ := float32(math.Min(float64(node1.BoundingBox.Min.Z), float64(node2.BoundingBox.Min.Z)))
	maxX := float32(math.Max(float64(node1.BoundingBox.Max.X), float64(node2.BoundingBox.Max.X)))
	maxY := float32(math.Max(float64(node1.BoundingBox.Max.Y), float64(node2.BoundingBox.Max.Y)))
	maxZ := float32(math.Max(float64(node1.BoundingBox.Max.Z), float64(node2.BoundingBox.Max.Z)))
	box := &BoxVolume{
		Min: Maths.Vector3{X: minX, Y: minY, Z: minZ},
		Max: Maths.Vector3{X: maxX, Y: maxY, Z: maxZ},
	}
	return &BVHNode{
		BoundingBox: box,
		Children:    []BVHNode{*node1, *node2},
	}
}

func BVHNodeFromTri(tri *Triangle) *BVHNode {
	box := tri.GetBoundingBox()
	return &BVHNode{
		Triangle:    tri,
		BoundingBox: &box,
	}
}

func (node *BVHNode) GetMidPoint() Maths.Vector3 {
	return node.BoundingBox.Min.Lerp(node.BoundingBox.Max, 0.5)
}

func (node *BVHNode) AddChild(child BVHNode) {
	node.Children = append(node.Children, child)
}

func (node *BVHNode) AddTriChild(childTri *Triangle) {
	node.Children = append(node.Children, *BVHNodeFromTri(childTri))
}
