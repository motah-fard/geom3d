package geom3d

// Line3 represents an infinite line in 3D, passing through Point in
// direction Dir. Unlike Ray3, a Line3 extends in both directions (t may be
// any real number, not just t >= 0). The direction is not automatically
// normalized.
//
// For distance and projection queries against a line given as two points,
// see DistancePointToLine and ProjectPointToLine, which predate this type
// and remain the simpler choice when you don't otherwise need a Line3
// value.
type Line3 struct {
	Point Vec3
	Dir   Vec3
}

// PointAt returns the point along the line at parameter t:
//
//	Point + t*Dir
func (l Line3) PointAt(t float64) Vec3 {
	return l.Point.Add(l.Dir.Scale(t))
}

// IsValid reports whether the line's direction is non-zero.
func (l Line3) IsValid() bool {
	return !AlmostZero(l.Dir.Norm())
}
