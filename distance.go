package geom3d

// DistancePointToPlane returns the signed perpendicular distance from point p
// to plane pl.
//
// The result is positive if p lies in the direction of the plane normal,
// negative if it lies in the opposite direction.
//
// If the plane is invalid, it returns 0.
func DistancePointToPlane(p Vec3, pl Plane) float64 {
	if !pl.IsValid() {
		return 0
	}

	n := pl.UnitNormal()
	return p.Sub(pl.Point).Dot(n)
}

// DistancePointToSegment returns the Euclidean distance from point p
// to segment s.
//
// If the segment is degenerate, it returns the distance from p to s.A.
func DistancePointToSegment(p Vec3, s Segment3) float64 {
	cp := ClosestPointOnSegment(p, s)
	return p.Distance(cp)
}

// DistancePointToLine returns the Euclidean distance from point p
// to the infinite line passing through points a and b.
//
// If a and b are the same point, it returns the distance from p to a.
func DistancePointToLine(p Vec3, a, b Vec3) float64 {
	proj := ProjectPointToLine(p, a, b)
	return p.Distance(proj)
}

// DistancePointToTriangle returns the Euclidean distance from point p
// to triangle t.
//
// If the triangle is degenerate, it returns the distance from p to the
// closest point on the triangle's longest non-degenerate edge. If all
// three vertices coincide, it returns the distance from p to t.A.
func DistancePointToTriangle(p Vec3, t Triangle) float64 {
	cp := ClosestPointOnTriangle(p, t)
	return p.Distance(cp)
}

// DistanceBetweenSegments returns the Euclidean distance between segments s1
// and s2.
func DistanceBetweenSegments(s1, s2 Segment3) float64 {
	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	return c1.Distance(c2)
}

// DistancePointToRay returns the Euclidean distance from point p to ray r.
//
// If the ray is invalid, it returns 0.
func DistancePointToRay(p Vec3, r Ray3) float64 {
	if !r.IsValid() {
		return 0
	}

	cp := ClosestPointOnRay(p, r)
	return p.Distance(cp)
}

// DistancePointToAABB returns the Euclidean distance from point p
// to axis-aligned bounding box b.
//
// If p lies inside the box, it returns 0.
//
// If the box is invalid, it returns 0.
func DistancePointToAABB(p Vec3, b AABB) float64 {
	if !b.IsValid() {
		return 0
	}

	cp := ClosestPointOnAABB(p, b)
	return p.Distance(cp)
}
