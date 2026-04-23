package geom3d

import (
	"fmt"
	"testing"
)

func TestClosestPointOnSegmentMiddle(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{10, 0, 0},
	}
	p := Vec3{4, 3, 0}

	got := ClosestPointOnSegment(p, s)
	want := Vec3{4, 0, 0}

	if got != want {
		t.Fatalf("ClosestPointOnSegment middle: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnSegmentBeforeStart(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{10, 0, 0},
	}
	p := Vec3{-2, 5, 0}

	got := ClosestPointOnSegment(p, s)
	want := Vec3{0, 0, 0}

	if got != want {
		t.Fatalf("ClosestPointOnSegment before start: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnSegmentAfterEnd(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{10, 0, 0},
	}
	p := Vec3{14, -1, 0}

	got := ClosestPointOnSegment(p, s)
	want := Vec3{10, 0, 0}

	if got != want {
		t.Fatalf("ClosestPointOnSegment after end: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnSegmentDegenerate(t *testing.T) {
	s := Segment3{
		A: Vec3{2, 2, 2},
		B: Vec3{2, 2, 2},
	}
	p := Vec3{9, 9, 9}

	got := ClosestPointOnSegment(p, s)
	want := Vec3{2, 2, 2}

	if got != want {
		t.Fatalf("ClosestPointOnSegment degenerate: got %#v, want %#v", got, want)
	}
}
func TestClosestPointOnTriangleVertexRegion(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	p := Vec3{-1, -1, 0}
	got := ClosestPointOnTriangle(p, tri)
	want := tri.A

	if got != want {
		t.Fatalf("vertex region: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnTriangleEdgeRegion(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{0, 2, 0},
	}

	p := Vec3{1, -1, 0}
	got := ClosestPointOnTriangle(p, tri)
	want := Vec3{1, 0, 0}

	if got != want {
		t.Fatalf("edge region: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnTriangleFaceRegion(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{0, 2, 0},
	}

	p := Vec3{0.5, 0.5, 3}
	got := ClosestPointOnTriangle(p, tri)
	want := Vec3{0.5, 0.5, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("face region: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnTriangleDegenerate(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{2, 0, 0},
	}

	p := Vec3{0.5, 2, 0}
	got := ClosestPointOnTriangle(p, tri)
	want := Vec3{0.5, 0, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("degenerate triangle: got %#v, want %#v", got, want)
	}
}
func ExampleClosestPointOnSegment() {
	seg := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 4, Y: 0, Z: 0},
	}

	p1 := Vec3{X: 2, Y: 3, Z: 0}
	p2 := Vec3{X: -1, Y: 1, Z: 0}
	p3 := Vec3{X: 6, Y: -2, Z: 0}

	fmt.Println(ClosestPointOnSegment(p1, seg))
	fmt.Println(ClosestPointOnSegment(p2, seg))
	fmt.Println(ClosestPointOnSegment(p3, seg))

	// Output:
	// {2 0 0}
	// {0 0 0}
	// {4 0 0}
}
func ExampleClosestPointOnTriangle() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 2, Z: 0},
	}

	p := Vec3{X: 0.5, Y: 0.5, Z: 3}
	cp := ClosestPointOnTriangle(p, tri)

	fmt.Printf("%.1f %.1f %.1f\n", cp.X, cp.Y, cp.Z)

	// Output:
	// 0.5 0.5 0.0
}
