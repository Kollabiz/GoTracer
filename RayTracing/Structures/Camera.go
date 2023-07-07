package Structures

import "RayTracer/Maths"

type Camera struct {
	Position       Maths.Vector3
	rotation       Maths.Vector3
	LensSize       Maths.Vector2
	FocalLength    float32
	rotationMatrix Maths.Mat3
}

func MakeCamera(position Maths.Vector3, rotation Maths.Vector3, lensSize Maths.Vector2, focalLength float32) Camera {
	return Camera{position, rotation, lensSize, focalLength, *Maths.Mat3FromEulerAngles(rotation)}
}

func (camera *Camera) GetRotation() Maths.Vector3 {
	return camera.rotation
}

func (camera *Camera) SetRotation(rotation Maths.Vector3) {
	camera.rotation = rotation
	camera.rotationMatrix = *Maths.Mat3FromEulerAngles(rotation)
}

func (camera *Camera) ResetRotation() {
	camera.rotation = Maths.ZeroVector3()
	camera.rotationMatrix = *Maths.Mat3Identity()
}
