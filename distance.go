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
