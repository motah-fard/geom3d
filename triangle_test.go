package geom3d

import "testing"

func TestTriangleNormal(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	got := tri.Normal()
	want := Vec3{0, 0, 1}

	if got != want {
		t.Fatalf("Normal: got %#v, want %#v", got, want)
	}
}

func TestTriangleArea(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	if !AlmostEqual(tri.Area(), 0.5) {
		t.Fatalf("Area: got %v, want 0.5", tri.Area())
	}
}

func TestTriangleIsDegenerate(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 1, 1},
		C: Vec3{2, 2, 2},
	}

	if !tri.IsDegenerate() {
		t.Fatal("expected degenerate triangle")
	}
}
