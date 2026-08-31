package geom3d

import "math"

// IntersectRayPlane computes the intersection point between ray r and plane pl.
//
// It returns the intersection point and true if the ray intersects the plane
// at parameter t >= 0.
//
// If the ray is parallel to the plane, the plane is invalid, or the
// intersection lies behind the ray origin, it returns Vec3{} and false.
func IntersectRayPlane(r Ray3, pl Plane) (Vec3, bool) {
	if !r.IsValid() || !pl.IsValid() {
		return Vec3{}, false
	}

	n := pl.UnitNormal()
	denom := r.Dir.Dot(n)

	if AlmostZero(denom) {
		return Vec3{}, false
	}

	t := pl.Point.Sub(r.Origin).Dot(n) / denom
	if t < 0 {
		return Vec3{}, false
	}

	return r.PointAt(t), true
}

// IntersectLinePlane computes the intersection point between infinite line
// l and plane pl.
//
// Unlike IntersectRayPlane, the intersection parameter is not restricted to
// t >= 0, since a Line3 extends in both directions.
//
// If the line is parallel to the plane or the plane is invalid, it returns
// Vec3{} and false.
func IntersectLinePlane(l Line3, pl Plane) (Vec3, bool) {
	if !l.IsValid() || !pl.IsValid() {
		return Vec3{}, false
	}

	n := pl.UnitNormal()
	denom := l.Dir.Dot(n)

	if AlmostZero(denom) {
		return Vec3{}, false
	}

	t := pl.Point.Sub(l.Point).Dot(n) / denom
	return l.PointAt(t), true
}

// IntersectPlanePlane computes the line of intersection between planes p1
// and p2.
//
// It returns that line and true if the planes are not parallel. The
// returned line's Dir is p1's normal crossed with p2's normal; it is not
// normalized.
//
// If either plane is invalid, or the planes are parallel (including
// coincident planes), it returns Line3{} and false.
func IntersectPlanePlane(p1, p2 Plane) (Line3, bool) {
	if !p1.IsValid() || !p2.IsValid() {
		return Line3{}, false
	}

	n1 := p1.UnitNormal()
	n2 := p2.UnitNormal()

	dir := n1.Cross(n2)
	denom := dir.Dot(dir)
	if AlmostZero(denom) {
		return Line3{}, false
	}

	d1 := n1.Dot(p1.Point)
	d2 := n2.Dot(p2.Point)
	point := n2.Scale(d1).Sub(n1.Scale(d2)).Cross(dir).Scale(1 / denom)

	return Line3{Point: point, Dir: dir}, true
}

// IntersectSegmentPlane computes the intersection point between segment s and plane pl.
//
// It returns the intersection point and true if the segment intersects the plane
// at parameter t in [0, 1].
//
// If the segment is parallel to the plane, the plane is invalid, or the
// intersection lies outside the segment, it returns Vec3{} and false.
func IntersectSegmentPlane(s Segment3, pl Plane) (Vec3, bool) {
	if s.IsDegenerate() || !pl.IsValid() {
		return Vec3{}, false
	}

	dir := s.Direction()
	n := pl.UnitNormal()
	denom := dir.Dot(n)

	if AlmostZero(denom) {
		return Vec3{}, false
	}

	t := pl.Point.Sub(s.A).Dot(n) / denom
	if t < 0 || t > 1 {
		return Vec3{}, false
	}

	return s.A.Add(dir.Scale(t)), true
}

// IntersectRayAABB reports whether ray r intersects axis-aligned bounding box b.
//
// It uses the slab method. Touching the box counts as intersection.
//
// It returns whether an intersection occurs, along with the entry and exit
// ray parameters tMin and tMax.
//
// If the ray is invalid or the box is invalid, it returns false, 0, 0.
func IntersectRayAABB(r Ray3, b AABB) (bool, float64, float64) {
	if !r.IsValid() || !b.IsValid() {
		return false, 0, 0
	}

	tMin := 0.0
	tMax := math.Inf(1)

	origin := [3]float64{r.Origin.X, r.Origin.Y, r.Origin.Z}
	dir := [3]float64{r.Dir.X, r.Dir.Y, r.Dir.Z}
	mins := [3]float64{b.Min.X, b.Min.Y, b.Min.Z}
	maxs := [3]float64{b.Max.X, b.Max.Y, b.Max.Z}

	for i := 0; i < 3; i++ {
		if AlmostZero(dir[i]) {
			// Ray is parallel to slab. Origin must lie within slab.
			if origin[i] < mins[i] || origin[i] > maxs[i] {
				return false, 0, 0
			}
			continue
		}

		t1 := (mins[i] - origin[i]) / dir[i]
		t2 := (maxs[i] - origin[i]) / dir[i]

		if t1 > t2 {
			t1, t2 = t2, t1
		}

		if t1 > tMin {
			tMin = t1
		}
		if t2 < tMax {
			tMax = t2
		}

		if tMin > tMax {
			return false, 0, 0
		}
	}

	if tMax < 0 {
		return false, 0, 0
	}

	return true, tMin, tMax
}

// IntersectRaySphere reports whether ray r intersects sphere s.
//
// It returns whether an intersection occurs, along with the near and far
// ray parameters tMin and tMax. If the ray originates inside the sphere,
// tMin is clamped to 0.
//
// If the ray or the sphere is invalid, or the sphere lies entirely behind
// the ray origin, it returns false, 0, 0.
func IntersectRaySphere(r Ray3, s Sphere) (bool, float64, float64) {
	if !r.IsValid() || !s.IsValid() {
		return false, 0, 0
	}

	oc := r.Origin.Sub(s.Center)
	a := r.Dir.Dot(r.Dir)
	b := 2 * oc.Dot(r.Dir)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return false, 0, 0
	}

	sqrtDisc := math.Sqrt(discriminant)
	t1 := (-b - sqrtDisc) / (2 * a)
	t2 := (-b + sqrtDisc) / (2 * a)

	if t1 > t2 {
		t1, t2 = t2, t1
	}

	if t2 < 0 {
		return false, 0, 0
	}

	tMin := t1
	if tMin < 0 {
		tMin = 0
	}

	return true, tMin, t2
}

// IntersectRayTriangle computes the intersection point between ray r and
// triangle tri using the Möller–Trumbore algorithm.
//
// It returns the intersection point and true if the ray intersects the
// triangle at parameter t >= 0. Intersections are reported on either face of
// the triangle; the test does not perform back-face culling.
//
// If the ray is invalid, the triangle is degenerate, the ray is parallel to
// the triangle's plane, or the intersection lies outside the triangle or
// behind the ray origin, it returns Vec3{} and false.
func IntersectRayTriangle(r Ray3, tri Triangle) (Vec3, bool) {
	if !r.IsValid() || tri.IsDegenerate() {
		return Vec3{}, false
	}

	p, _, ok := intersectLineTriangle(r.Origin, r.Dir, tri, math.Inf(1))
	return p, ok
}

// IntersectSegmentTriangle computes the intersection point between segment s
// and triangle tri, using the same Möller–Trumbore test as
// IntersectRayTriangle but bounded to the segment's own length.
//
// It returns the intersection point and true if the segment intersects the
// triangle at parameter t in [0, 1]. As with IntersectRayTriangle, either
// face of the triangle counts as a hit.
//
// If the segment is degenerate, the triangle is degenerate, the segment is
// parallel to the triangle's plane, or the intersection lies outside the
// triangle or outside the segment, it returns Vec3{} and false.
func IntersectSegmentTriangle(s Segment3, tri Triangle) (Vec3, bool) {
	if s.IsDegenerate() || tri.IsDegenerate() {
		return Vec3{}, false
	}

	p, _, ok := intersectLineTriangle(s.A, s.Direction(), tri, 1)
	return p, ok
}

// intersectLineTriangle implements the Möller–Trumbore ray/triangle
// intersection test for the parametric line origin + t*dir, accepting any
// hit with 0 <= t <= maxT. IntersectRayTriangle and IntersectSegmentTriangle
// are thin wrappers around this shared implementation, differing only in
// maxT (unbounded for a ray, 1 for a segment).
func intersectLineTriangle(origin, dir Vec3, tri Triangle, maxT float64) (Vec3, float64, bool) {
	edge1 := tri.EdgeAB()
	edge2 := tri.EdgeAC()

	pvec := dir.Cross(edge2)
	det := edge1.Dot(pvec)

	if AlmostZero(det) {
		return Vec3{}, 0, false
	}

	invDet := 1 / det
	tvec := origin.Sub(tri.A)

	u := tvec.Dot(pvec) * invDet
	if u < 0 || u > 1 {
		return Vec3{}, 0, false
	}

	qvec := tvec.Cross(edge1)
	v := dir.Dot(qvec) * invDet
	if v < 0 || u+v > 1 {
		return Vec3{}, 0, false
	}

	t := edge2.Dot(qvec) * invDet
	if t < 0 || t > maxT {
		return Vec3{}, 0, false
	}

	return origin.Add(dir.Scale(t)), t, true
}

// IntersectAABBSphere reports whether axis-aligned bounding box b and sphere
// s intersect or touch.
//
// If either shape is invalid, it returns false.
func IntersectAABBSphere(b AABB, s Sphere) bool {
	if !b.IsValid() || !s.IsValid() {
		return false
	}
	return DistancePointToAABB(s.Center, b) <= s.Radius
}

// IntersectSegmentAABB reports whether segment s intersects axis-aligned
// bounding box b, using the same slab method as IntersectRayAABB but bounded
// to the segment's own length.
//
// It returns whether an intersection occurs, along with the entry and exit
// parameters tMin and tMax along the segment (0 at s.A, 1 at s.B).
//
// If the segment is degenerate, the box is invalid, or the box lies beyond
// s.B, it returns false, 0, 0.
func IntersectSegmentAABB(s Segment3, b AABB) (bool, float64, float64) {
	if s.IsDegenerate() || !b.IsValid() {
		return false, 0, 0
	}

	hit, tMin, tMax := IntersectRayAABB(Ray3{Origin: s.A, Dir: s.Direction()}, b)
	if !hit || tMin > 1 {
		return false, 0, 0
	}
	if tMax > 1 {
		tMax = 1
	}

	return true, tMin, tMax
}

// IntersectSegmentSphere reports whether segment s intersects sphere sph,
// using the same quadratic test as IntersectRaySphere but bounded to the
// segment's own length.
//
// It returns whether an intersection occurs, along with the entry and exit
// parameters tMin and tMax along the segment (0 at s.A, 1 at s.B).
//
// If the segment is degenerate, the sphere is invalid, or the sphere lies
// beyond s.B, it returns false, 0, 0.
func IntersectSegmentSphere(s Segment3, sph Sphere) (bool, float64, float64) {
	if s.IsDegenerate() || !sph.IsValid() {
		return false, 0, 0
	}

	hit, tMin, tMax := IntersectRaySphere(Ray3{Origin: s.A, Dir: s.Direction()}, sph)
	if !hit || tMin > 1 {
		return false, 0, 0
	}
	if tMax > 1 {
		tMax = 1
	}

	return true, tMin, tMax
}

// IntersectSegments reports whether segments s1 and s2 intersect at a single point.
//
// If they intersect at a single point, it returns that point and true.
//
// If they are disjoint, skew, parallel without intersection, or overlap over a
// non-zero interval, it returns Vec3{} and false.
func IntersectSegments(s1, s2 Segment3) (Vec3, bool) {
	c1, c2 := ClosestPointsBetweenSegments(s1, s2)

	if !AlmostZero(c1.Distance(c2)) {
		return Vec3{}, false
	}

	// Reject collinear overlap as a non-unique intersection.
	d1 := s1.Direction()
	d2 := s2.Direction()

	if AlmostZero(d1.Cross(d2).Norm()) {
		v := s2.A.Sub(s1.A)
		if AlmostZero(v.Cross(d1).Norm()) {
			return Vec3{}, false
		}
	}

	return c1, true
}
