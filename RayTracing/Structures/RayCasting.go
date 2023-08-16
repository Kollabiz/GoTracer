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

func CastRay(origin Maths.Vector3, direction Maths.Vector3, ctx *RenderContext, excludeTriangle *Triangle) RayIntersection {
	ray := Ray{Origin: origin, Direction: direction.Normalized()}
	// Precomputing dirFrac
	dirFrac := Maths.MakeVector3(1/ray.Direction.X, 1/ray.Direction.Y, 1/ray.Direction.Z)
	var t *float64
	// Allocating a big slice to not bother GC much
	var meshes = make([]Mesh, len(ctx.Scene.Meshes))
	actualMeshCount := 0
	for i := 0; i < len(ctx.Scene.Meshes); i++ {
		if ray.IntersectAABB(ctx.Scene.Meshes[i].Volume, dirFrac, t) {
			meshes[actualMeshCount] = ctx.Scene.Meshes[i]
			actualMeshCount++
		}
	}
	// Intersecting meshes
	var nearestIntersection = MakeNonHitRay()
	var minDepth = float32(math.Inf(1))
	for i := 0; i < actualMeshCount; i++ {
		nodeMesh := meshes[i].GetTransformed()
		var triangle *Triangle
		for j := 0; j < len(nodeMesh); j++ {
			triangle = &nodeMesh[j]
			// Backface culling
			if triangle.TriangleNormal.Dot(direction) >= 0 {
				continue
			}
			rayHit, hitPos, bHitPos := ray.Intersect(triangle)
			if rayHit {
				rayDepth := hitPos.Sub(origin).Length()
				if rayDepth < minDepth {
					minDepth = rayDepth
					nearestIntersection = MakeRayIntersection(
						rayHit,
						&meshes[i],
						triangle,
						hitPos,
						bHitPos,
						rayDepth,
					)
				}
			}
		}
	}
	return nearestIntersection
}
