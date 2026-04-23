package geom3d

import (
	"fmt"
	"testing"
)

func TestSegment3Length(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{3, 4, 0},
	}

	if !AlmostEqual(s.Length(), 5) {
		t.Fatalf("Length: got %v, want 5", s.Length())
	}
}

func TestSegment3Midpoint(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 4, 6},
	}

	got := s.Midpoint()
	want := Vec3{1, 2, 3}

	if got != want {
		t.Fatalf("Midpoint: got %#v, want %#v", got, want)
	}
}

func TestSegment3IsDegenerate(t *testing.T) {
	s := Segment3{
		A: Vec3{1, 1, 1},
		B: Vec3{1, 1, 1},
	}

	if !s.IsDegenerate() {
		t.Fatal("expected degenerate segment")
	}
}

func TestSegment3Length2(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{3, 4, 0},
	}

	if !AlmostEqual(s.Length2(), 25) {
		t.Fatalf("Length2: got %v, want 25", s.Length2())
	}
}
func ExampleSegment3_Midpoint() {
	s := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 4, Z: 6},
	}

	fmt.Println(s.Midpoint())

	// Output:
	// {1 2 3}
}
