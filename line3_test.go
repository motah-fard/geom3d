package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestLine3PointAt(t *testing.T) {
	l := Line3{Point: Vec3{1, 1, 1}, Dir: Vec3{2, 0, 0}}

	got := l.PointAt(2)
	want := Vec3{5, 1, 1}

	if got != want {
		t.Fatalf("PointAt: got %#v, want %#v", got, want)
	}
}

func TestLine3PointAtNegativeT(t *testing.T) {
	l := Line3{Point: Vec3{1, 1, 1}, Dir: Vec3{2, 0, 0}}

	got := l.PointAt(-1)
	want := Vec3{-1, 1, 1}

	if got != want {
		t.Fatalf("PointAt negative t: got %#v, want %#v", got, want)
	}
}

func TestLine3IsValid(t *testing.T) {
	l := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{1, 0, 0}}
	if !l.IsValid() {
		t.Fatal("expected valid line")
	}

	bad := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{0, 0, 0}}
	if bad.IsValid() {
		t.Fatal("expected invalid line for zero direction")
	}
}

func TestIntersectLinePlaneHit(t *testing.T) {
	l := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{0, 0, 1}}
	pl := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}}

	got, ok := IntersectLinePlane(l, pl)
	want := Vec3{0, 0, 5}

	if !ok {
		t.Fatal("expected line-plane intersection")
	}
	if got != want {
		t.Fatalf("IntersectLinePlane: got %#v, want %#v", got, want)
	}
}

func TestIntersectLinePlaneBehindLineOrigin(t *testing.T) {
	// Unlike IntersectRayPlane, a negative t is still a valid hit for an
	// infinite line.
	l := Line3{Point: Vec3{0, 0, 10}, Dir: Vec3{0, 0, 1}}
	pl := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}}

	got, ok := IntersectLinePlane(l, pl)
	want := Vec3{0, 0, 5}

	if !ok {
		t.Fatal("expected line-plane intersection behind the line's own point")
	}
	if got != want {
		t.Fatalf("IntersectLinePlane: got %#v, want %#v", got, want)
	}
}

func TestIntersectLinePlaneParallel(t *testing.T) {
	l := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{1, 0, 0}}
	pl := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}}

	_, ok := IntersectLinePlane(l, pl)
	if ok {
		t.Fatal("expected no intersection for line parallel to plane")
	}
}

func TestIntersectLinePlaneInvalid(t *testing.T) {
	bad := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{0, 0, 0}}
	pl := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}}

	_, ok := IntersectLinePlane(bad, pl)
	if ok {
		t.Fatal("expected no intersection for invalid line")
	}
}

func TestClosestPointsBetweenLinesSkew(t *testing.T) {
	l1 := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{1, 0, 0}}
	l2 := Line3{Point: Vec3{0, 1, 1}, Dir: Vec3{0, 0, 1}}

	c1, c2 := ClosestPointsBetweenLines(l1, l2)
	want1 := Vec3{0, 0, 0}
	want2 := Vec3{0, 1, 0}

	if c1 != want1 || c2 != want2 {
		t.Fatalf("ClosestPointsBetweenLines: got %#v, %#v, want %#v, %#v", c1, c2, want1, want2)
	}
}

func TestClosestPointsBetweenLinesParallel(t *testing.T) {
	l1 := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{1, 0, 0}}
	l2 := Line3{Point: Vec3{0, 3, 0}, Dir: Vec3{1, 0, 0}}

	c1, c2 := ClosestPointsBetweenLines(l1, l2)

	if got, want := c1.Distance(c2), 3.0; !AlmostEqual(got, want) {
		t.Fatalf("distance between parallel lines: got %v, want %v", got, want)
	}
}

func TestClosestPointsBetweenLinesBothDegenerate(t *testing.T) {
	l1 := Line3{Point: Vec3{1, 2, 3}, Dir: Vec3{}}
	l2 := Line3{Point: Vec3{4, 5, 6}, Dir: Vec3{}}

	c1, c2 := ClosestPointsBetweenLines(l1, l2)

	if c1 != l1.Point || c2 != l2.Point {
		t.Fatalf("ClosestPointsBetweenLines both degenerate: got %#v, %#v, want %#v, %#v", c1, c2, l1.Point, l2.Point)
	}
}

func TestDistanceBetweenLinesSkew(t *testing.T) {
	l1 := Line3{Point: Vec3{0, 0, 0}, Dir: Vec3{1, 0, 0}}
	l2 := Line3{Point: Vec3{0, 1, 1}, Dir: Vec3{0, 0, 1}}

	got := DistanceBetweenLines(l1, l2)
	want := 1.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenLines: got %v, want %v", got, want)
	}
}

func TestIntersectPlanePlaneHit(t *testing.T) {
	p1 := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}} // z = 5
	p2 := Plane{Point: Vec3{3, 0, 0}, Normal: Vec3{1, 0, 0}} // x = 3

	l, ok := IntersectPlanePlane(p1, p2)
	if !ok {
		t.Fatal("expected plane-plane intersection")
	}

	// The intersection is the line x=3, z=5, running parallel to Y.
	if !AlmostEqual(l.Point.X, 3) || !AlmostEqual(l.Point.Z, 5) {
		t.Fatalf("IntersectPlanePlane point: got %#v, want X=3, Z=5", l.Point)
	}
	dir := l.Dir.Normalize()
	if !AlmostZero(dir.X) || !AlmostEqual(math.Abs(dir.Y), 1) || !AlmostZero(dir.Z) {
		t.Fatalf("IntersectPlanePlane direction: got %#v, want parallel to Y axis", l.Dir)
	}
}

func TestIntersectPlanePlaneParallel(t *testing.T) {
	p1 := Plane{Point: Vec3{0, 0, 0}, Normal: Vec3{0, 0, 1}}
	p2 := Plane{Point: Vec3{0, 0, 5}, Normal: Vec3{0, 0, 1}}

	_, ok := IntersectPlanePlane(p1, p2)
	if ok {
		t.Fatal("expected no intersection for parallel planes")
	}
}

func TestIntersectPlanePlaneCoincident(t *testing.T) {
	p1 := Plane{Point: Vec3{0, 0, 0}, Normal: Vec3{0, 0, 1}}
	p2 := Plane{Point: Vec3{5, 5, 0}, Normal: Vec3{0, 0, 1}}

	_, ok := IntersectPlanePlane(p1, p2)
	if ok {
		t.Fatal("expected no unique intersection for coincident planes")
	}
}

func TestIntersectPlanePlaneInvalid(t *testing.T) {
	bad := Plane{Point: Vec3{0, 0, 0}, Normal: Vec3{0, 0, 0}}
	p2 := Plane{Point: Vec3{0, 0, 0}, Normal: Vec3{1, 0, 0}}

	_, ok := IntersectPlanePlane(bad, p2)
	if ok {
		t.Fatal("expected no intersection for invalid plane")
	}
}

func ExampleIntersectPlanePlane() {
	p1 := Plane{Point: Vec3{X: 0, Y: 0, Z: 5}, Normal: Vec3{X: 0, Y: 0, Z: 1}}
	p2 := Plane{Point: Vec3{X: 3, Y: 0, Z: 0}, Normal: Vec3{X: 1, Y: 0, Z: 0}}

	l, ok := IntersectPlanePlane(p1, p2)
	fmt.Println(ok)
	// Adding 0 normalizes an exact negative zero (-0 == 0 numerically, but
	// they print differently) that this particular cross product produces
	// on its Y component.
	fmt.Printf("{%v %v %v}\n", l.Point.X+0, l.Point.Y+0, l.Point.Z+0)

	// Output:
	// true
	// {3 0 5}
}

func ExampleIntersectLinePlane() {
	l := Line3{Point: Vec3{X: 0, Y: 0, Z: 10}, Dir: Vec3{X: 0, Y: 0, Z: 1}}
	pl := Plane{Point: Vec3{X: 0, Y: 0, Z: 5}, Normal: Vec3{X: 0, Y: 0, Z: 1}}

	p, ok := IntersectLinePlane(l, pl)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {0 0 5}
}
