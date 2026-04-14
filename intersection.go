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
// If the ray is invalid or the box is invalid, it returns false.
func IntersectRayAABB(r Ray3, b AABB) bool {
	if !r.IsValid() || !b.IsValid() {
		return false
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
				return false
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
			return false
		}
	}

	return tMax >= 0
}
