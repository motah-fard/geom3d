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

// IntersectRayOBB reports whether ray r intersects oriented bounding box b.
//
// It returns whether an intersection occurs, along with the entry and exit
// ray parameters tMin and tMax, exactly as IntersectRayAABB does. It works
// by transforming the ray into b's local (axis-aligned) frame and reusing
// IntersectRayAABB there — a rigid rotation and translation don't change
// the ray's parameter t, so the resulting tMin/tMax are valid in world
// space unchanged.
//
// If the ray is invalid or the box is invalid, it returns false, 0, 0.
func IntersectRayOBB(r Ray3, b OBB) (bool, float64, float64) {
	if !r.IsValid() || !b.IsValid() {
		return false, 0, 0
	}

	localRay := Ray3{
		Origin: b.toLocal(r.Origin),
		Dir:    b.Orientation.Transpose().MulVec(r.Dir),
	}
	localBox := AABB{Min: b.HalfExtents.Scale(-1), Max: b.HalfExtents}

	return IntersectRayAABB(localRay, localBox)
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

	t0, t1, ok := intersectLineSphere(r.Origin, r.Dir, s)
	if !ok || t1 < 0 {
		return false, 0, 0
	}

	tMin := t0
	if tMin < 0 {
		tMin = 0
	}

	return true, tMin, t1
}

// intersectLineSphere solves the quadratic for where the parametric line
// origin + t*dir meets sphere s, returning the two roots (t0 <= t1) and
// whether a real intersection exists. Unlike IntersectRaySphere, it does
// not restrict t to any range — callers are responsible for applying
// ray/segment/capsule-specific bounds on top of the raw roots.
//
// dir must be non-zero (every caller already guarantees this via its own
// IsValid check). Given that, the quadratic's leading coefficient
// dir.Dot(dir) is always positive, so (-b-sqrtDisc)/(2a) is always <=
// (-b+sqrtDisc)/(2a) — the roots come out pre-ordered with no swap needed.
func intersectLineSphere(origin, dir Vec3, s Sphere) (t0, t1 float64, ok bool) {
	oc := origin.Sub(s.Center)
	a := dir.Dot(dir)
	b := 2 * oc.Dot(dir)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, false
	}

	sqrtDisc := math.Sqrt(discriminant)
	return (-b - sqrtDisc) / (2 * a), (-b + sqrtDisc) / (2 * a), true
}

// IntersectRayCapsule reports whether ray r intersects capsule c.
//
// It returns whether an intersection occurs, along with the entry and exit
// ray parameters tMin and tMax, exactly as IntersectRaySphere does. If the
// ray originates inside the capsule, tMin is clamped to 0.
//
// The capsule's boundary is treated as three pieces: the infinite cylinder
// along its axis (valid only between the two end caps) and the two
// hemispherical end caps (valid only beyond the cylinder's own range). For
// each piece, a candidate crossing is kept only if it actually falls within
// that piece's own valid region — a ray's raw intersection with, say, the
// sphere at c.A only counts if the crossing point's projection onto the
// capsule's axis falls at or before c.A, not somewhere the cylinder's
// lateral surface would apply instead.
//
// If the ray or the capsule is invalid, or the capsule lies entirely
// behind the ray origin, it returns false, 0, 0.
func IntersectRayCapsule(r Ray3, c Capsule) (bool, float64, float64) {
	if !r.IsValid() || !c.IsValid() {
		return false, 0, 0
	}

	axis := c.B.Sub(c.A)
	dd := axis.Dot(axis)

	if AlmostZero(dd) {
		// Degenerate capsule (A == B): equivalent to a plain sphere.
		return IntersectRaySphere(r, Sphere{Center: c.A, Radius: c.Radius})
	}

	m := r.Origin.Sub(c.A)
	n := r.Dir
	nd := n.Dot(axis)
	md := m.Dot(axis)

	// axisParam returns the point at ray parameter t's projection onto the
	// capsule's axis, in [0, 1] between c.A and c.B (extending outside that
	// range beyond the caps).
	axisParam := func(t float64) float64 {
		return (md + t*nd) / dd
	}

	var haveEntry, haveExit bool
	var entry, exit float64

	considerEntry := func(t float64) {
		if !haveEntry || t < entry {
			entry, haveEntry = t, true
		}
	}
	considerExit := func(t float64) {
		if !haveExit || t > exit {
			exit, haveExit = t, true
		}
	}

	// The infinite cylinder along the capsule's axis, valid only where the
	// axis projection lands between the two end caps.
	//
	// a = dd*nn - nd*nd is, by the Cauchy-Schwarz identity, equal to
	// dd*nn*sin²θ where θ is the angle between the ray direction and the
	// axis — always >= 0, and strictly > 0 once the !AlmostZero(a) guard
	// below passes. So, exactly as in intersectLineSphere, the two roots
	// come out pre-ordered (t0 <= t1) with no swap needed.
	nn := n.Dot(n)
	mn := m.Dot(n)
	mm := m.Dot(m)

	if a := dd*nn - nd*nd; !AlmostZero(a) {
		bCoef := dd*mn - nd*md
		cCoef := dd*mm - md*md - c.Radius*c.Radius*dd
		disc := bCoef*bCoef - a*cCoef

		if disc >= 0 {
			sqrtDisc := math.Sqrt(disc)
			t0 := (-bCoef - sqrtDisc) / a
			t1 := (-bCoef + sqrtDisc) / a

			if s := axisParam(t0); s >= 0 && s <= 1 {
				considerEntry(t0)
			}
			if s := axisParam(t1); s >= 0 && s <= 1 {
				considerExit(t1)
			}
		}
	}

	// The two hemispherical end caps, each valid only beyond its own end
	// of the axis.
	if t0, t1, ok := intersectLineSphere(r.Origin, r.Dir, Sphere{Center: c.A, Radius: c.Radius}); ok {
		if axisParam(t0) <= 0 {
			considerEntry(t0)
		}
		if axisParam(t1) <= 0 {
			considerExit(t1)
		}
	}
	if t0, t1, ok := intersectLineSphere(r.Origin, r.Dir, Sphere{Center: c.B, Radius: c.Radius}); ok {
		if axisParam(t0) >= 1 {
			considerEntry(t0)
		}
		if axisParam(t1) >= 1 {
			considerExit(t1)
		}
	}

	if !haveEntry || !haveExit || exit < 0 {
		return false, 0, 0
	}
	if entry < 0 {
		entry = 0
	}

	return true, entry, exit
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
