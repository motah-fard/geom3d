package geom3d

// ProjectPointToPlane projects point p onto plane pl.
//
// If the plane is invalid, it returns Vec3{}.
func ProjectPointToPlane(p Vec3, pl Plane) Vec3 {
	if !pl.IsValid() {
		return Vec3{}
	}

	d := DistancePointToPlane(p, pl)
	n := pl.UnitNormal()
	return p.Sub(n.Scale(d))
}

// ProjectPointToLine projects point p onto the infinite line passing through
// points a and b.
//
// If a and b are the same point, it returns a.
func ProjectPointToLine(p Vec3, a, b Vec3) Vec3 {
	ab := b.Sub(a)
	denom := ab.Norm2()
	if AlmostZero(denom) {
		return a
	}

	t := p.Sub(a).Dot(ab) / denom
	return a.Add(ab.Scale(t))
}
