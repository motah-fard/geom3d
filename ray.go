package geom3d

// Ray3 represents a 3D ray with an origin and direction.
// The direction is not automatically normalized.
type Ray3 struct {
	Origin Vec3
	Dir    Vec3
}

// PointAt returns the point along the ray at parameter t:
//
//	Origin + t*Dir
func (r Ray3) PointAt(t float64) Vec3 {
	return r.Origin.Add(r.Dir.Scale(t))
}

// IsValid reports whether the ray direction is non-zero.
func (r Ray3) IsValid() bool {
	return !AlmostZero(r.Dir.Norm())
}
