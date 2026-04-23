package geom3d

import (
	"fmt"
	"testing"
)

func TestBarycentricCoordinatesAtVertices(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	u, v, w := BarycentricCoordinates(tri.A, tri)
	if !AlmostEqual(u, 1) || !AlmostEqual(v, 0) || !AlmostEqual(w, 0) {
		t.Fatalf("at A: got (%v, %v, %v), want (1, 0, 0)", u, v, w)
	}

	u, v, w = BarycentricCoordinates(tri.B, tri)
	if !AlmostEqual(u, 0) || !AlmostEqual(v, 1) || !AlmostEqual(w, 0) {
		t.Fatalf("at B: got (%v, %v, %v), want (0, 1, 0)", u, v, w)
	}

	u, v, w = BarycentricCoordinates(tri.C, tri)
	if !AlmostEqual(u, 0) || !AlmostEqual(v, 0) || !AlmostEqual(w, 1) {
		t.Fatalf("at C: got (%v, %v, %v), want (0, 0, 1)", u, v, w)
	}
}

func TestBarycentricCoordinatesInsideTriangle(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	p := Vec3{0.25, 0.25, 0}
	u, v, w := BarycentricCoordinates(p, tri)

	if !AlmostEqual(u, 0.5) || !AlmostEqual(v, 0.25) || !AlmostEqual(w, 0.25) {
		t.Fatalf("inside triangle: got (%v, %v, %v), want (0.5, 0.25, 0.25)", u, v, w)
	}
}

func TestBarycentricCoordinatesOutsideTriangle(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
		C: Vec3{0, 1, 0},
	}

	p := Vec3{1.5, 0.25, 0}
	u, v, w := BarycentricCoordinates(p, tri)

	if AlmostEqual(u, 0) && AlmostEqual(v, 0) && AlmostEqual(w, 0) {
		t.Fatal("outside triangle should still produce barycentric coordinates for a non-degenerate triangle")
	}
	if !AlmostEqual(u+v+w, 1) {
		t.Fatalf("outside triangle: got u+v+w = %v, want 1", u+v+w)
	}
}

func TestBarycentricCoordinatesDegenerateTriangle(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 1, 1},
		C: Vec3{2, 2, 2},
	}

	u, v, w := BarycentricCoordinates(Vec3{1, 0, 0}, tri)
	if u != 0 || v != 0 || w != 0 {
		t.Fatalf("degenerate triangle: got (%v, %v, %v), want (0, 0, 0)", u, v, w)
	}
}

func ExampleBarycentricCoordinates() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 1, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 1, Z: 0},
	}

	p := Vec3{X: 0.25, Y: 0.25, Z: 0}
	u, v, w := BarycentricCoordinates(p, tri)

	fmt.Printf("%.2f %.2f %.2f\n", u, v, w)

	// Output:
	// 0.50 0.25 0.25
}
