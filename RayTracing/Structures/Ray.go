package Structures

import "RayTracer/Maths"

const (
	Epsilon = 0.0000001
)

type Ray struct {
	Origin    Maths.Vector3
	Direction Maths.Vector3
}

func (ray *Ray) GetPoint(f float32) Maths.Vector3 {
	return ray.Origin.Add(ray.Direction.MulF(f))
}

// Intersect preforms Ray - Triangle intersection with Moller-Trumbore algorithm
func (ray *Ray) Intersect(tri Triangle, intersectionPoint Maths.Vector3, barycentricIntersection Maths.Vector3) bool {
	var h, s, q Maths.Vector3
	e1 := tri.CalcFirstEdge()
	e2 := tri.CalcSecondEdge()
	var a, f, u, v float32
	h = ray.Direction.Cross(e2)
	a = e1.Dot(h)

	if a > -Epsilon && a < Epsilon {
		return false
	}

	f = 1 / a
	s = ray.Origin.Sub(tri.V1.Position)
	u = f * s.Dot(h)

	if u < 0 || u > 1 { // U + V must be less than 1
		return false
	}

	q = s.Cross(e1)
	v = f * ray.Direction.Dot(q)

	if v < 0 || u+v > 1 {
		return false
	}

	t := f * e2.Dot(q)

	if t <= Epsilon {
		return false
	}
	intersectionPoint = ray.GetPoint(t)
	barycentricIntersection = Maths.Vector3{X: u, Y: v, Z: 1 - u - v}
	return true
}
