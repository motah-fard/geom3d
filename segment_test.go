package geom3d

import "testing"

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
