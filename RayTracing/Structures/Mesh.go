package Structures

import "RayTracer/Maths"

type Mesh struct {
	rotation       Maths.Vector3
	Position       Maths.Vector3
	scale          Maths.Vector3
	rotationMatrix *Maths.Mat3
	scaleMatrix    *Maths.Mat3
	MeshName       string
	Triangles      []Triangle
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
