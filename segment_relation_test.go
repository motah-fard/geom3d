package geom3d

import (
	"fmt"
	"testing"
)

func TestSegmentsOverlapTrue(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 0, 0},
		B: Vec3{6, 0, 0},
	}

	if !SegmentsOverlap(s1, s2) {
		t.Fatal("expected collinear overlapping segments to overlap")
	}
}

func TestSegmentsOverlapFalseDisjointCollinear(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 0, 0},
		B: Vec3{3, 0, 0},
	}

	if SegmentsOverlap(s1, s2) {
		t.Fatal("expected disjoint collinear segments not to overlap")
	}
}

func TestSegmentsOverlapFalseEndpointTouchOnly(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 0, 0},
		B: Vec3{2, 0, 0},
	}

	if SegmentsOverlap(s1, s2) {
		t.Fatal("expected endpoint-touching segments not to count as overlap")
	}
}

func TestSegmentsOverlapFalseParallelNonCollinear(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 1, 0},
		B: Vec3{2, 1, 0},
	}

	if SegmentsOverlap(s1, s2) {
		t.Fatal("expected parallel non-collinear segments not to overlap")
	}
}

func TestSegmentsOverlapFalseSkew(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 0, 1},
		B: Vec3{3, 0, 1},
	}

	if SegmentsOverlap(s1, s2) {
		t.Fatal("expected skew segments not to overlap")
	}
}

func TestSegmentsOverlapFalseDegenerate(t *testing.T) {
	s1 := Segment3{
		A: Vec3{1, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}

	if SegmentsOverlap(s1, s2) {
		t.Fatal("expected degenerate segment not to count as overlap")
	}
}

func ExampleSegmentsOverlap() {
	s1 := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 4, Y: 0, Z: 0},
	}
	s2 := Segment3{
		A: Vec3{X: 2, Y: 0, Z: 0},
		B: Vec3{X: 6, Y: 0, Z: 0},
	}

	fmt.Println(SegmentsOverlap(s1, s2))

	// Output:
	// true
}
