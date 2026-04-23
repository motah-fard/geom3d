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
