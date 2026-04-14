package geom3d

// Segment3 represents a finite line segment from A to B.
type Segment3 struct {
	A Vec3
	B Vec3
}

// Direction returns the vector from A to B.
func (s Segment3) Direction() Vec3 {
	return s.B.Sub(s.A)
}

// Length2 returns the squared length of the segment.
func (s Segment3) Length2() float64 {
	return s.Direction().Norm2()
}

// Length returns the Euclidean length of the segment.
func (s Segment3) Length() float64 {
	return s.Direction().Norm()
}

// Midpoint returns the midpoint of the segment.
func (s Segment3) Midpoint() Vec3 {
	return Vec3{
		X: (s.A.X + s.B.X) / 2,
		Y: (s.A.Y + s.B.Y) / 2,
		Z: (s.A.Z + s.B.Z) / 2,
	}
}

// IsDegenerate reports whether the segment length is effectively zero.
func (s Segment3) IsDegenerate() bool {
	return AlmostZero(s.Length())
}
