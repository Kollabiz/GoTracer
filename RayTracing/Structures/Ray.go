package Structures

import (
	"RayTracer/Maths"
	"math"
)

type Ray struct {
	Origin    Maths.Vector3
	Direction Maths.Vector3
}

func (ray *Ray) GetPoint(f float32) Maths.Vector3 {
	return ray.Origin.Add(ray.Direction.MulF(f))
}

// Intersect preforms Ray - Triangle intersection with Moller-Trumbore algorithm
func (ray *Ray) Intersect(tri *Triangle) (hit bool, intersectionPoint, barycentricIntersection Maths.Vector3) {
	var h, s, q Maths.Vector3
	e1 := tri.CalcFirstEdge()
	e2 := tri.CalcSecondEdge()
	var a, f, u, v float32
	h = ray.Direction.Cross(e2)
	a = e1.Dot(h)

	if a > -Epsilon && a < Epsilon {
		return false, Maths.ZeroVector3(), Maths.ZeroVector3()
	}

	f = 1 / a
	s = ray.Origin.Sub(tri.V1.Position)
	u = f * s.Dot(h)

	if u < 0 || u > 1 { // U + V must be less than 1
		return false, Maths.ZeroVector3(), Maths.ZeroVector3()
	}

	q = s.Cross(e1)
	v = f * ray.Direction.Dot(q)

	if v < 0 || u+v > 1 {
		return false, Maths.ZeroVector3(), Maths.ZeroVector3()
	}

	t := f * e2.Dot(q)

	if t <= Epsilon {
		return false, Maths.ZeroVector3(), Maths.ZeroVector3()
	}
	p := ray.GetPoint(t)
	return true, p, Maths.Vector3{u, v, 1 - u - v}
}

func (ray *Ray) IntersectAABB(box *BoxVolume, dirFrac Maths.Vector3, t *float64) bool {
	t1 := float64((box.Min.X - ray.Origin.X) * dirFrac.X)
	t2 := float64((box.Max.X - ray.Origin.X) * dirFrac.X)
	t3 := float64((box.Min.Y - ray.Origin.Y) * dirFrac.Y)
	t4 := float64((box.Max.Y - ray.Origin.Y) * dirFrac.Y)
	t5 := float64((box.Min.Z - ray.Origin.Z) * dirFrac.Z)
	t6 := float64((box.Max.Z - ray.Origin.Z) * dirFrac.Z)

	tmin := math.Max(math.Max(math.Min(t1, t2), math.Min(t3, t4)), math.Min(t5, t6))
	tmax := math.Min(math.Min(math.Max(t1, t2), math.Max(t3, t4)), math.Max(t5, t6))

	if tmax < 0 {
		t = &tmax
		return true
	}

	if tmin > tmax {
		t = &tmax
		return false
	}

	t = &tmin
	return true
}
