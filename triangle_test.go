package geom3d

import (
	"fmt"
	"testing"
)

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
func TestTriangleEdgeAB(t *testing.T) {
	tri := Triangle{
		A: Vec3{1, 2, 3},
		B: Vec3{4, 6, 8},
		C: Vec3{0, 0, 0},
	}

	got := tri.EdgeAB()
	want := Vec3{3, 4, 5}

	if got != want {
		t.Fatalf("EdgeAB: got %#v, want %#v", got, want)
	}
}

func TestTriangleEdgeAC(t *testing.T) {
	tri := Triangle{
		A: Vec3{1, 2, 3},
		B: Vec3{0, 0, 0},
		C: Vec3{4, 6, 8},
	}

	got := tri.EdgeAC()
	want := Vec3{3, 4, 5}

	if got != want {
		t.Fatalf("EdgeAC: got %#v, want %#v", got, want)
	}
}
func ExampleTriangle_Normal() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 1, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 1, Z: 0},
	}

	n := tri.Normal()

	fmt.Printf("%.0f %.0f %.0f\n", n.X, n.Y, n.Z)
	fmt.Printf("%.1f\n", tri.Area())
	fmt.Println(tri.IsDegenerate())

	// Output:
	// 0 0 1
	// 0.5
	// false
}
