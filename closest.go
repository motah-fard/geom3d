package geom3d

// ClosestPointOnSegment returns the closest point on segment s to point p.
//
// If the segment is degenerate, it returns s.A.
func ClosestPointOnSegment(p Vec3, s Segment3) Vec3 {
	ab := s.B.Sub(s.A)
	denom := ab.Norm2()
	if AlmostZero(denom) {
		return s.A
	}

	t := p.Sub(s.A).Dot(ab) / denom

	if t <= 0 {
		return s.A
	}
	if t >= 1 {
		return s.B
	}

	return s.A.Add(ab.Scale(t))
}
