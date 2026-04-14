package geom3d

// Plane represents a plane in 3D using a point on the plane
// and a normal vector.
//
// The normal is not automatically normalized.
type Plane struct {
	Point  Vec3
	Normal Vec3
}

// IsValid reports whether the plane normal is non-zero.
func (p Plane) IsValid() bool {
	return !AlmostZero(p.Normal.Norm())
}

// UnitNormal returns the normalized plane normal.
// If the normal is zero, it returns Vec3{}.
func (p Plane) UnitNormal() Vec3 {
	return p.Normal.Normalize()
}
