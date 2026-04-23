package geom3d

import (
	"fmt"
	"testing"
)

func TestDistancePointToPlane(t *testing.T) {
	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 0},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	tests := []struct {
		name string
		p    Vec3
		want float64
	}{
		{
			name: "point above plane",
			p:    Vec3{X: 1, Y: 2, Z: 5},
			want: 5,
		},
		{
			name: "point below plane",
			p:    Vec3{X: -1, Y: 0, Z: -3},
			want: -3,
		},
		{
			name: "point on plane",
			p:    Vec3{X: 4, Y: 5, Z: 0},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DistancePointToPlane(tc.p, pl)
			if !AlmostEqual(got, tc.want) {
				t.Fatalf("DistancePointToPlane(%v, pl) = %v, want %v", tc.p, got, tc.want)
			}
		})
	}
}

func TestDistancePointToPlaneInvalidPlane(t *testing.T) {
	p := Vec3{X: 1, Y: 2, Z: 3}
	got := DistancePointToPlane(p, Plane{})

	if got != 0 {
		t.Fatalf("DistancePointToPlane with invalid plane = %v, want 0", got)
	}
}

func TestDistancePointToSegment(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}

	tests := []struct {
		name string
		p    Vec3
		want float64
	}{
		{
			name: "point above middle of segment",
			p:    Vec3{2, 3, 0},
			want: 3,
		},
		{
			name: "point before start of segment",
			p:    Vec3{-1, 0, 0},
			want: 1,
		},
		{
			name: "point after end of segment",
			p:    Vec3{6, 0, 0},
			want: 2,
		},
		{
			name: "point on segment",
			p:    Vec3{1, 0, 0},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DistancePointToSegment(tc.p, s)
			if !AlmostEqual(got, tc.want) {
				t.Fatalf("DistancePointToSegment(%v, s) = %v, want %v", tc.p, got, tc.want)
			}
		})
	}
}

func TestDistancePointToSegmentDegenerate(t *testing.T) {
	s := Segment3{
		A: Vec3{1, 2, 3},
		B: Vec3{1, 2, 3},
	}
	p := Vec3{4, 6, 3}

	got := DistancePointToSegment(p, s)
	want := 5.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToSegment degenerate: got %v, want %v", got, want)
	}
}
func TestDistancePointToLine(t *testing.T) {
	a := Vec3{0, 0, 0}
	b := Vec3{4, 0, 0}

	tests := []struct {
		name string
		p    Vec3
		want float64
	}{
		{
			name: "point above line",
			p:    Vec3{2, 3, 0},
			want: 3,
		},
		{
			name: "point before line start but same infinite line",
			p:    Vec3{-1, 2, 0},
			want: 2,
		},
		{
			name: "point on line",
			p:    Vec3{10, 0, 0},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DistancePointToLine(tc.p, a, b)
			if !AlmostEqual(got, tc.want) {
				t.Fatalf("DistancePointToLine(%v, a, b) = %v, want %v", tc.p, got, tc.want)
			}
		})
	}
}

func TestDistancePointToLineDegenerate(t *testing.T) {
	a := Vec3{1, 2, 3}
	p := Vec3{4, 6, 3}

	got := DistancePointToLine(p, a, a)
	want := 5.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToLine degenerate: got %v, want %v", got, want)
	}
}
func TestDistancePointToTriangleFaceRegion(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{0, 2, 0},
	}

	p := Vec3{0.5, 0.5, 3}
	got := DistancePointToTriangle(p, tri)
	want := 3.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToTriangle face region: got %v, want %v", got, want)
	}
}

func TestDistancePointToTriangleEdgeRegion(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{0, 2, 0},
	}

	p := Vec3{1, -2, 0}
	got := DistancePointToTriangle(p, tri)
	want := 2.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToTriangle edge region: got %v, want %v", got, want)
	}
}

func TestDistancePointToTriangleDegenerate(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{2, 0, 0},
	}

	p := Vec3{1, 3, 0}
	got := DistancePointToTriangle(p, tri)
	want := 3.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToTriangle degenerate: got %v, want %v", got, want)
	}
}
func TestDistanceBetweenSegmentsIntersecting(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, -1, 0},
		B: Vec3{1, 1, 0},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := 0.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments intersecting: got %v, want %v", got, want)
	}
}

func TestDistanceBetweenSegmentsParallel(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 1, 0},
		B: Vec3{2, 1, 0},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := 1.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments parallel: got %v, want %v", got, want)
	}
}

func TestDistanceBetweenSegmentsSkew(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 1, 1},
		B: Vec3{1, 1, -1},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := 1.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments skew: got %v, want %v", got, want)
	}
}
func TestDistanceBetweenSegmentsOverlappingCollinear(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 0, 0},
		B: Vec3{6, 0, 0},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := 0.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments overlapping collinear: got %v, want %v", got, want)
	}
}

func TestDistanceBetweenSegmentsEndpointToEndpoint(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 1, 0},
		B: Vec3{2, 2, 0},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := Vec3{1, 0, 0}.Distance(Vec3{2, 1, 0})

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments endpoint-to-endpoint: got %v, want %v", got, want)
	}
}

func TestDistanceBetweenSegmentsOneDegenerateOneNormal(t *testing.T) {
	s1 := Segment3{
		A: Vec3{1, 2, 0},
		B: Vec3{1, 2, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}

	got := DistanceBetweenSegments(s1, s2)
	want := 2.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistanceBetweenSegments one degenerate one normal: got %v, want %v", got, want)
	}
}
func ExampleDistancePointToPlane() {
	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 0},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	p := Vec3{X: 2, Y: 3, Z: 5}
	fmt.Println(DistancePointToPlane(p, pl))

	// Output:
	// 5
}
func ExampleDistancePointToSegment() {
	s := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 4, Y: 0, Z: 0},
	}
	p := Vec3{X: 2, Y: 3, Z: 0}

	fmt.Println(DistancePointToSegment(p, s))

	// Output:
	// 3
}
func ExampleDistancePointToLine() {
	a := Vec3{X: 0, Y: 0, Z: 0}
	b := Vec3{X: 4, Y: 0, Z: 0}
	p := Vec3{X: 2, Y: 3, Z: 0}

	fmt.Println(DistancePointToLine(p, a, b))

	// Output:
	// 3
}
func ExampleDistancePointToTriangle() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 2, Z: 0},
	}

	p := Vec3{X: 0.5, Y: 0.5, Z: 3}
	fmt.Println(DistancePointToTriangle(p, tri))

	// Output:
	// 3
}
func ExampleDistanceBetweenSegments() {
	s1 := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 0, Z: 0},
	}
	s2 := Segment3{
		A: Vec3{X: 0, Y: 1, Z: 0},
		B: Vec3{X: 2, Y: 1, Z: 0},
	}

	fmt.Println(DistanceBetweenSegments(s1, s2))

	// Output:
	// 1
}
