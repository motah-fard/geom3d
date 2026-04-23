package geom3d

// BarycentricCoordinates returns the barycentric coordinates (u, v, w)
// of point p with respect to triangle t, where:
//
//	p = u*A + v*B + w*C
//
// and u + v + w = 1.
//
// The point p does not need to lie on the triangle plane or inside the triangle.
//
// If the triangle is degenerate, it returns 0, 0, 0.
func BarycentricCoordinates(p Vec3, t Triangle) (u, v, w float64) {
	v0 := t.B.Sub(t.A)
	v1 := t.C.Sub(t.A)
	v2 := p.Sub(t.A)

	d00 := v0.Dot(v0)
	d01 := v0.Dot(v1)
	d11 := v1.Dot(v1)
	d20 := v2.Dot(v0)
	d21 := v2.Dot(v1)

	denom := d00*d11 - d01*d01
	if AlmostZero(denom) {
		return 0, 0, 0
	}

	v = (d11*d20 - d01*d21) / denom
	w = (d00*d21 - d01*d20) / denom
	u = 1 - v - w

	return u, v, w
}
