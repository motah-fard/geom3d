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

// ClosestPointOnTriangle returns the closest point on triangle t to point p.
//
// If the triangle is degenerate, it returns the closest point on the longest
// non-degenerate edge. If all three vertices coincide, it returns t.A.
func ClosestPointOnTriangle(p Vec3, t Triangle) Vec3 {
	a := t.A
	b := t.B
	c := t.C

	ab := b.Sub(a)
	ac := c.Sub(a)
	ap := p.Sub(a)

	d1 := ab.Dot(ap)
	d2 := ac.Dot(ap)
	if d1 <= 0 && d2 <= 0 {
		return a
	}

	bp := p.Sub(b)
	d3 := ab.Dot(bp)
	d4 := ac.Dot(bp)
	if d3 >= 0 && d4 <= d3 {
		return b
	}

	vc := d1*d4 - d3*d2
	if vc <= 0 && d1 >= 0 && d3 <= 0 {
		v := d1 / (d1 - d3)
		return a.Add(ab.Scale(v))
	}

	cp := p.Sub(c)
	d5 := ab.Dot(cp)
	d6 := ac.Dot(cp)
	if d6 >= 0 && d5 <= d6 {
		return c
	}

	vb := d5*d2 - d1*d6
	if vb <= 0 && d2 >= 0 && d6 <= 0 {
		w := d2 / (d2 - d6)
		return a.Add(ac.Scale(w))
	}

	va := d3*d6 - d5*d4
	if va <= 0 && (d4-d3) >= 0 && (d5-d6) >= 0 {
		bc := c.Sub(b)
		w := (d4 - d3) / ((d4 - d3) + (d5 - d6))
		return b.Add(bc.Scale(w))
	}

	denom := va + vb + vc
	if AlmostZero(denom) {
		ab2 := ab.Norm2()
		ac2 := ac.Norm2()
		bc := c.Sub(b)
		bc2 := bc.Norm2()

		switch {
		case ab2 >= ac2 && ab2 >= bc2 && !AlmostZero(ab2):
			return ClosestPointOnSegment(p, Segment3{A: a, B: b})
		case ac2 >= ab2 && ac2 >= bc2 && !AlmostZero(ac2):
			return ClosestPointOnSegment(p, Segment3{A: a, B: c})
		case !AlmostZero(bc2):
			return ClosestPointOnSegment(p, Segment3{A: b, B: c})
		default:
			return a
		}
	}

	invDenom := 1.0 / denom
	v := vb * invDenom
	w := vc * invDenom

	return a.Add(ab.Scale(v)).Add(ac.Scale(w))
}

func clamp01(x float64) float64 {
	if x < 0 {
		return 0
	}
	if x > 1 {
		return 1
	}
	return x
}

// ClosestPointsBetweenSegments returns the pair of closest points on segments
// s1 and s2.
//
// If either segment is degenerate, it falls back to the corresponding
// point-segment query. If both are degenerate, it returns their endpoints.
func ClosestPointsBetweenSegments(s1, s2 Segment3) (Vec3, Vec3) {
	p1 := s1.A
	q1 := s1.B
	p2 := s2.A
	q2 := s2.B

	d1 := q1.Sub(p1)
	d2 := q2.Sub(p2)
	r := p1.Sub(p2)

	a := d1.Dot(d1)
	e := d2.Dot(d2)
	f := d2.Dot(r)

	// Both segments degenerate to points.
	if AlmostZero(a) && AlmostZero(e) {
		return p1, p2
	}

	// First segment degenerates to a point.
	if AlmostZero(a) {
		t := clamp01(f / e)
		return p1, p2.Add(d2.Scale(t))
	}

	c := d1.Dot(r)

	// Second segment degenerates to a point.
	if AlmostZero(e) {
		s := clamp01(-c / a)
		return p1.Add(d1.Scale(s)), p2
	}

	b := d1.Dot(d2)
	denom := a*e - b*b

	var s float64
	if !AlmostZero(denom) {
		s = clamp01((b*f - c*e) / denom)
	} else {
		// Parallel or nearly parallel.
		s = 0
	}

	t := (b*s + f) / e

	if t < 0 {
		t = 0
		s = clamp01(-c / a)
	} else if t > 1 {
		t = 1
		s = clamp01((b - c) / a)
	}

	c1 := p1.Add(d1.Scale(s))
	c2 := p2.Add(d2.Scale(t))
	return c1, c2
}
