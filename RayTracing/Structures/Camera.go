package Structures

import (
	"RayTracer/Maths"
)

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

func (camera *Camera) GetForwardDirection() Maths.Vector3 {
	return Maths.MakeVector3(0, 0, 1).MatMul(&camera.rotationMatrix)
}

func (camera *Camera) GetViewportPlaneCorners() [4]Maths.Vector3 {
	hw := camera.LensSize.X / 2
	hh := camera.LensSize.Y / 2
	upperLeft := Maths.MakeVector3(-hw, hh, 0).MatMul(&camera.rotationMatrix)
	upperRight := Maths.MakeVector3(hw, hh, 0).MatMul(&camera.rotationMatrix)
	downLeft := Maths.MakeVector3(upperLeft.X, -upperLeft.Y, -upperLeft.Z)
	downRight := Maths.MakeVector3(upperRight.X, -upperRight.Y, -upperRight.Z)
	return [4]Maths.Vector3{upperLeft, upperRight, downLeft, downRight}
}
