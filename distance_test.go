package geom3d

import (
	"math"
	"testing"
)

func TestDistancePointToPlane(t *testing.T) {
	pl := Plane{
		Point:  Vec3{0, 0, 0},
		Normal: Vec3{0, 0, 2},
	}

	p := Vec3{0, 0, 5}
	got := DistancePointToPlane(p, pl)
	want := 5.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToPlane: got %v, want %v", got, want)
	}
}

func TestDistancePointToPlaneNegative(t *testing.T) {
	pl := Plane{
		Point:  Vec3{0, 0, 0},
		Normal: Vec3{0, 0, 1},
	}

	p := Vec3{0, 0, -3}
	got := DistancePointToPlane(p, pl)
	want := -3.0

	if !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToPlane negative: got %v, want %v", got, want)
	}
}

func TestDistancePointToPlaneInvalidPlane(t *testing.T) {
	pl := Plane{}
	p := Vec3{1, 2, 3}

	got := DistancePointToPlane(p, pl)
	if !AlmostEqual(got, 0) {
		t.Fatalf("DistancePointToPlane invalid plane: got %v, want 0", got)
	}
}

func TestDistancePointToPlaneUnsignedExample(t *testing.T) {
	pl := Plane{
		Point:  Vec3{0, 0, 0},
		Normal: Vec3{0, 0, 1},
	}

	p := Vec3{0, 0, -4}
	got := math.Abs(DistancePointToPlane(p, pl))
	want := 4.0

	if !AlmostEqual(got, want) {
		t.Fatalf("unsigned distance example: got %v, want %v", got, want)
	}
}
