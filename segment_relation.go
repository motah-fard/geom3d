package geom3d

// SegmentsOverlap reports whether s1 and s2 are collinear and overlap over a
// non-zero interval.
//
// It returns false for disjoint segments, skew segments, parallel non-collinear
// segments, and segments that only touch at a single point.
func SegmentsOverlap(s1, s2 Segment3) bool {
	if s1.IsDegenerate() || s2.IsDegenerate() {
		return false
	}

	d1 := s1.Direction()
	d2 := s2.Direction()

	// Must be parallel.
	if !AlmostZero(d1.Cross(d2).Norm()) {
		return false
	}

	// Must be collinear.
	if !AlmostZero(s2.A.Sub(s1.A).Cross(d1).Norm()) {
		return false
	}

	// Project onto the dominant axis to test 1D overlap length.
	absX := d1.X
	if absX < 0 {
		absX = -absX
	}
	absY := d1.Y
	if absY < 0 {
		absY = -absY
	}
	absZ := d1.Z
	if absZ < 0 {
		absZ = -absZ
	}

	var a1, b1, a2, b2 float64
	switch {
	case absX >= absY && absX >= absZ:
		a1, b1 = s1.A.X, s1.B.X
		a2, b2 = s2.A.X, s2.B.X
	case absY >= absX && absY >= absZ:
		a1, b1 = s1.A.Y, s1.B.Y
		a2, b2 = s2.A.Y, s2.B.Y
	default:
		a1, b1 = s1.A.Z, s1.B.Z
		a2, b2 = s2.A.Z, s2.B.Z
	}

	if a1 > b1 {
		a1, b1 = b1, a1
	}
	if a2 > b2 {
		a2, b2 = b2, a2
	}

	left := a1
	if a2 > left {
		left = a2
	}
	right := b1
	if b2 < right {
		right = b2
	}

	// Require positive overlap length, not just a shared endpoint.
	return right-left > Epsilon
}
