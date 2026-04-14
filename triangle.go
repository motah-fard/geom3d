package geom3d

// Triangle represents a triangle in 3D with vertices A, B, and C.
type Triangle struct {
	A Vec3
	B Vec3
	C Vec3
}

// EdgeAB returns the edge vector from A to B.
func (t Triangle) EdgeAB() Vec3 {
	return t.B.Sub(t.A)
}

// EdgeAC returns the edge vector from A to C.
func (t Triangle) EdgeAC() Vec3 {
	return t.C.Sub(t.A)
}

// Normal returns the non-unit normal vector of the triangle,
// computed as (B-A) x (C-A).
func (t Triangle) Normal() Vec3 {
	return t.EdgeAB().Cross(t.EdgeAC())
}

// Area returns the area of the triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Normal().Norm()
}

// IsDegenerate reports whether the triangle area is effectively zero.
func (t Triangle) IsDegenerate() bool {
	return AlmostZero(t.Area())
}
