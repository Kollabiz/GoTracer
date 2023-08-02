package Structures

import (
	"RayTracer/Maths"
	"math"
)

type RayIntersection struct {
	Hit            bool
	HitMesh        *Mesh
	HitTriangle    *Triangle
	HitPosition    Maths.Vector3
	BarycentricHit Maths.Vector3
	RayLength      float32
}

func MakeRayIntersection(hit bool, hitMesh *Mesh, hitTriangle *Triangle, hitPosition Maths.Vector3, barycentricHit Maths.Vector3, rayLength float32) RayIntersection {
	return RayIntersection{hit, hitMesh, hitTriangle, hitPosition, barycentricHit, rayLength}
}

func MakeNonHitRay() RayIntersection {
	return RayIntersection{false, nil, nil, Maths.ZeroVector3(), Maths.ZeroVector3(), -1}
}

func (ray RayIntersection) GetHitNormal() Maths.Vector3 {
	//return ray.HitTriangle.GetSmoothNormal(ray.BarycentricHit)
	return ray.HitTriangle.TriangleNormal
}

func (ray RayIntersection) GetHitUv() Maths.Vector2 {
	return ray.HitTriangle.GetUv(ray.BarycentricHit)
}

func (ray RayIntersection) GetHitMaterial() *Material {
	return ray.HitTriangle.Material
}

// Tracing a single ray

func triScale(tri *Triangle) float32 {
	return (tri.CalcFirstEdge().Length() + tri.CalcSecondEdge().Length() + tri.CalcThirdEdge().Length()) / 3
}

func canHit(rayOrig, rayDirection Maths.Vector3, tri *Triangle) bool {
	midpoint := tri.GetCenter()
	centerRayDir := midpoint.Sub(rayOrig).Normalized()
	longestEdge := float32(math.Max(math.Max(float64(tri.CalcFirstEdge().Length()), float64(tri.CalcSecondEdge().Length())), float64(tri.CalcThirdEdge().Length())))
	dot := 1 - centerRayDir.Dot(rayDirection)
	if dot*triScale(tri) > longestEdge {
		return false
	}
	return true
}

func CastRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext, excludeTriangle *Triangle) RayIntersection {
	ray := Ray{Origin: origin, Direction: direction}
	var minDepth float32 = math.MaxFloat32
	var tri *Triangle
	minResult := MakeNonHitRay()
	for i := 0; i < len(ctx.Scene.Meshes); i++ {
		mesh := ctx.Scene.Meshes[i].GetTransformed()
		for j := 0; j < len(mesh); j++ {
			tri = &mesh[j]
			if excludeTriangle != nil && AreEqualTriangles(tri, excludeTriangle) {
				continue
			}
			if tri.TriangleNormal.Dot(direction) >= 0 {
				continue
			}
			rayHit, hitPos, bHitPos := ray.Intersect(tri)
			if rayHit {
				rayLength := hitPos.Sub(origin).Length()
				if rayLength < minDepth {
					minDepth = rayLength
					minResult = MakeRayIntersection(
						true,
						&ctx.Scene.Meshes[i],
						tri,
						hitPos,
						bHitPos,
						rayLength,
					)
				}
			}
		}
	}
	return minResult
}