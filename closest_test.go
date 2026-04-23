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
func TestClosestPointsBetweenSegmentsIntersecting(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, -1, 0},
		B: Vec3{1, 1, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	want := Vec3{1, 0, 0}

	if c1 != want || c2 != want {
		t.Fatalf("intersecting segments: got %#v and %#v, want %#v and %#v", c1, c2, want, want)
	}
}

func TestClosestPointsBetweenSegmentsSkewSeparated(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 1, 1},
		B: Vec3{1, 1, -1},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	want1 := Vec3{1, 0, 0}
	want2 := Vec3{1, 1, 0}

	if !AlmostEqual(c1.X, want1.X) || !AlmostEqual(c1.Y, want1.Y) || !AlmostEqual(c1.Z, want1.Z) {
		t.Fatalf("skew segments c1: got %#v, want %#v", c1, want1)
	}
	if !AlmostEqual(c2.X, want2.X) || !AlmostEqual(c2.Y, want2.Y) || !AlmostEqual(c2.Z, want2.Z) {
		t.Fatalf("skew segments c2: got %#v, want %#v", c2, want2)
	}
}

func TestClosestPointsBetweenSegmentsParallel(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 1, 0},
		B: Vec3{2, 1, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)

	if !AlmostEqual(c1.Y, 0) || !AlmostEqual(c2.Y, 1) {
		t.Fatalf("parallel segments: got %#v and %#v", c1, c2)
	}
	if !AlmostEqual(c1.Distance(c2), 1) {
		t.Fatalf("parallel segments distance: got %v, want 1", c1.Distance(c2))
	}
}

func TestClosestPointsBetweenSegmentsDegenerate(t *testing.T) {
	s1 := Segment3{
		A: Vec3{1, 2, 3},
		B: Vec3{1, 2, 3},
	}
	s2 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	want1 := Vec3{1, 2, 3}
	want2 := Vec3{1, 0, 0}

	if c1 != want1 {
		t.Fatalf("degenerate segment c1: got %#v, want %#v", c1, want1)
	}
	if c2 != want2 {
		t.Fatalf("degenerate segment c2: got %#v, want %#v", c2, want2)
	}
}
func TestClosestPointsBetweenSegmentsOverlappingCollinear(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 0, 0},
		B: Vec3{6, 0, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)

	if !AlmostEqual(c1.Distance(c2), 0) {
		t.Fatalf("overlapping collinear segments: got points %#v and %#v with distance %v, want distance 0", c1, c2, c1.Distance(c2))
	}
}

func TestClosestPointsBetweenSegmentsEndpointToEndpoint(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 1, 0},
		B: Vec3{2, 2, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	want1 := Vec3{1, 0, 0}
	want2 := Vec3{2, 1, 0}

	if c1 != want1 || c2 != want2 {
		t.Fatalf("endpoint-to-endpoint nearest: got %#v and %#v, want %#v and %#v", c1, c2, want1, want2)
	}
}

func TestClosestPointsBetweenSegmentsOneDegenerateOneNormal(t *testing.T) {
	s1 := Segment3{
		A: Vec3{1, 2, 0},
		B: Vec3{1, 2, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)
	want1 := Vec3{1, 2, 0}
	want2 := Vec3{1, 0, 0}

	if c1 != want1 || c2 != want2 {
		t.Fatalf("one degenerate one normal: got %#v and %#v, want %#v and %#v", c1, c2, want1, want2)
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
func ExampleClosestPointsBetweenSegments() {
	s1 := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 0, Z: 0},
	}
	s2 := Segment3{
		A: Vec3{X: 1, Y: 1, Z: 0},
		B: Vec3{X: 1, Y: -1, Z: 0},
	}

	c1, c2 := ClosestPointsBetweenSegments(s1, s2)

	fmt.Println(c1)
	fmt.Println(c2)

	// Output:
	// {1 0 0}
	// {1 0 0}
}
