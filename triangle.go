package geom3d

import "math"

// Triangle represents a triangle in 3D with vertices A, B, and C.
type Triangle struct {
	A Vec3
	B Vec3
	C Vec3
}

// EdgeAB returns the edge vector from A to B.
func (t Triangle) EdgeAB() Vec3 {
	return t.B.Sub(t.A)
}

// EdgeAC returns the edge vector from A to C.
func (t Triangle) EdgeAC() Vec3 {
	return t.C.Sub(t.A)
}

// Normal returns the non-unit normal vector of the triangle,
// computed as (B-A) x (C-A).
func (t Triangle) Normal() Vec3 {
	return t.EdgeAB().Cross(t.EdgeAC())
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Normal().Norm()
}

// IsDegenerate reports whether the triangle area is effectively zero.
func (t Triangle) IsDegenerate() bool {
	return AlmostZero(t.Area())
}

// edges returns the triangle's three edges as segments, in winding order.
func (t Triangle) edges() [3]Segment3 {
	return [3]Segment3{
		{A: t.A, B: t.B},
		{A: t.B, B: t.C},
		{A: t.C, B: t.A},
	}
}

// Overlaps reports whether t and other intersect or touch, in 3D.
//
// If t or other is degenerate, it returns false.
func (t Triangle) Overlaps(other Triangle) bool {
	if t.IsDegenerate() || other.IsDegenerate() {
		return false
	}

	for _, e := range t.edges() {
		if _, ok := IntersectSegmentTriangle(e, other); ok {
			return true
		}
	}
	for _, e := range other.edges() {
		if _, ok := IntersectSegmentTriangle(e, t); ok {
			return true
		}
	}

	// If no edge of either triangle crosses the other, the only remaining
	// way they can overlap is if they are coplanar (a partial overlap or
	// one fully containing the other) — IntersectSegmentTriangle always
	// rejects edges lying within the triangle's own plane as "parallel,"
	// so that case needs a dedicated 2D test.
	n1 := t.Normal()
	n2 := other.Normal()
	if !AlmostZero(n1.Cross(n2).Norm()) {
		return false
	}

	plane := Plane{Point: t.A, Normal: n1}
	if !plane.IsValid() || !AlmostZero(DistancePointToPlane(other.A, plane)) {
		return false // parallel but distinct planes
	}

	return coplanarTrianglesOverlap(t, other, n1)
}

// point2 is a 2D point used by the coplanar-triangle overlap test below.
type point2 struct {
	x, y float64
}

// coplanarTrianglesOverlap tests two triangles known to lie in the same
// plane (with the given normal) for 2D overlap, using the separating axis
// theorem. This covers cases Triangle.Overlaps' edge-crossing test cannot:
// one triangle fully containing the other, or partial overlap where no
// edge of either triangle crosses the other's boundary from a 3D
// perspective (since coplanar edges are parallel to the other triangle's
// plane by definition).
func coplanarTrianglesOverlap(t, other Triangle, normal Vec3) bool {
	i0, i1 := dropDominantAxis(normal)
	proj := func(v Vec3) point2 {
		return point2{x: vec3Component(v, i0), y: vec3Component(v, i1)}
	}

	p := [3]point2{proj(t.A), proj(t.B), proj(t.C)}
	q := [3]point2{proj(other.A), proj(other.B), proj(other.C)}

	return !triangle2DHasSeparatingAxis(p, q) && !triangle2DHasSeparatingAxis(q, p)
}

// dropDominantAxis returns the indices of the two axes to keep when
// projecting a plane with the given normal down to 2D: the axis the normal
// points along most strongly is the one the plane varies least along, so
// it's dropped.
func dropDominantAxis(n Vec3) (keep0, keep1 int) {
	ax, ay, az := math.Abs(n.X), math.Abs(n.Y), math.Abs(n.Z)
	switch {
	case ax >= ay && ax >= az:
		return 1, 2
	case ay >= ax && ay >= az:
		return 0, 2
	default:
		return 0, 1
	}
}

// vec3Component returns v's component at index 0 (X), 1 (Y), or 2 (Z).
func vec3Component(v Vec3, i int) float64 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	default:
		return v.Z
	}
}

// triangle2DHasSeparatingAxis reports whether an axis perpendicular to one
// of p's edges separates 2D triangles p and q (the standard separating
// axis theorem test, checked from one triangle's perspective; the caller
// checks both).
func triangle2DHasSeparatingAxis(p, q [3]point2) bool {
	for i := 0; i < 3; i++ {
		a, b := p[i], p[(i+1)%3]
		axisX, axisY := -(b.y - a.y), b.x-a.x

		minP, maxP := triangle2DProjectOntoAxis(p, axisX, axisY)
		minQ, maxQ := triangle2DProjectOntoAxis(q, axisX, axisY)

		if maxP < minQ-Epsilon || maxQ < minP-Epsilon {
			return true
		}
	}
	return false
}

// triangle2DProjectOntoAxis projects each of pts onto the axis
// (axisX, axisY) and returns the resulting interval.
func triangle2DProjectOntoAxis(pts [3]point2, axisX, axisY float64) (min, max float64) {
	min = pts[0].x*axisX + pts[0].y*axisY
	max = min
	for i := 1; i < 3; i++ {
		v := pts[i].x*axisX + pts[i].y*axisY
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}
