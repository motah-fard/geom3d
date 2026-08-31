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

func TestTriangleOverlapsNonCoplanarIntersecting(t *testing.T) {
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}} // z=0 plane
	// t2 lies in the y=1 plane; its A-B edge runs straight through z=0
	// at (1, 1, 0), which is inside t1.
	t2 := Triangle{A: Vec3{1, 1, -5}, B: Vec3{1, 1, 5}, C: Vec3{3, 1, 5}}

	if !t1.Overlaps(t2) {
		t.Fatal("expected non-coplanar triangles to overlap")
	}
	if !t2.Overlaps(t1) {
		t.Fatal("expected Overlaps to be symmetric")
	}
}

func TestTriangleOverlapsNonCoplanarDisjoint(t *testing.T) {
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	t2 := Triangle{A: Vec3{100, 100, -5}, B: Vec3{100, 100, 5}, C: Vec3{103, 100, 5}}

	if t1.Overlaps(t2) {
		t.Fatal("expected far-apart, non-coplanar triangles not to overlap")
	}
}

func TestTriangleOverlapsParallelDistinctPlanes(t *testing.T) {
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	t2 := Triangle{A: Vec3{0, 0, 5}, B: Vec3{4, 0, 5}, C: Vec3{0, 4, 5}}

	if t1.Overlaps(t2) {
		t.Fatal("expected triangles in parallel, distinct planes not to overlap")
	}
}

func TestTriangleOverlapsCoplanarPartial(t *testing.T) {
	// Same shape, one shifted diagonally so their footprints partially
	// overlap. Every edge of either triangle lies in the shared z=0 plane,
	// so IntersectSegmentTriangle rejects all of them as "parallel" — this
	// case only passes if the coplanar 2D fallback in Overlaps runs.
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	t2 := Triangle{A: Vec3{1, 1, 0}, B: Vec3{5, 1, 0}, C: Vec3{1, 5, 0}}

	if !t1.Overlaps(t2) {
		t.Fatal("expected coplanar, partially overlapping triangles to overlap")
	}
}

func TestTriangleOverlapsCoplanarPartialXPlane(t *testing.T) {
	// Same shape and offset as TestTriangleOverlapsCoplanarPartial, but
	// lying in the x=5 plane instead of z=0, so the shared normal is
	// dominated by X rather than Z. Exercises dropDominantAxis's other
	// branches.
	t1 := Triangle{A: Vec3{5, 0, 0}, B: Vec3{5, 4, 0}, C: Vec3{5, 0, 4}}
	t2 := Triangle{A: Vec3{5, 1, 1}, B: Vec3{5, 5, 1}, C: Vec3{5, 1, 5}}

	if !t1.Overlaps(t2) {
		t.Fatal("expected coplanar triangles in the x=5 plane to overlap")
	}
}

func TestTriangleOverlapsCoplanarPartialYPlane(t *testing.T) {
	t1 := Triangle{A: Vec3{0, 3, 0}, B: Vec3{4, 3, 0}, C: Vec3{0, 3, 4}}
	t2 := Triangle{A: Vec3{1, 3, 1}, B: Vec3{5, 3, 1}, C: Vec3{1, 3, 5}}

	if !t1.Overlaps(t2) {
		t.Fatal("expected coplanar triangles in the y=3 plane to overlap")
	}
}

func TestTriangleOverlapsCoplanarContainment(t *testing.T) {
	big := Triangle{A: Vec3{0, 0, 0}, B: Vec3{10, 0, 0}, C: Vec3{0, 10, 0}}
	small := Triangle{A: Vec3{1, 1, 0}, B: Vec3{2, 1, 0}, C: Vec3{1, 2, 0}}

	if !big.Overlaps(small) {
		t.Fatal("expected a triangle fully containing another to overlap it")
	}
	if !small.Overlaps(big) {
		t.Fatal("expected Overlaps to be symmetric for containment")
	}
}

func TestTriangleOverlapsCoplanarDisjoint(t *testing.T) {
	t1 := Triangle{A: Vec3{0, 0, 0}, B: Vec3{1, 0, 0}, C: Vec3{0, 1, 0}}
	t2 := Triangle{A: Vec3{5, 5, 0}, B: Vec3{6, 5, 0}, C: Vec3{5, 6, 0}}

	if t1.Overlaps(t2) {
		t.Fatal("expected disjoint coplanar triangles not to overlap")
	}
}

func TestTriangleOverlapsDegenerate(t *testing.T) {
	valid := Triangle{A: Vec3{0, 0, 0}, B: Vec3{4, 0, 0}, C: Vec3{0, 4, 0}}
	degenerate := Triangle{A: Vec3{0, 0, 0}, B: Vec3{2, 0, 0}, C: Vec3{4, 0, 0}}

	if valid.Overlaps(degenerate) {
		t.Fatal("expected Overlaps to be false when the other triangle is degenerate")
	}
	if degenerate.Overlaps(valid) {
		t.Fatal("expected Overlaps to be false when the receiver is degenerate")
	}
}

func ExampleTriangle_Overlaps() {
	t1 := Triangle{A: Vec3{X: 0, Y: 0, Z: 0}, B: Vec3{X: 4, Y: 0, Z: 0}, C: Vec3{X: 0, Y: 4, Z: 0}}
	t2 := Triangle{A: Vec3{X: 1, Y: 1, Z: 0}, B: Vec3{X: 5, Y: 1, Z: 0}, C: Vec3{X: 1, Y: 5, Z: 0}}

	fmt.Println(t1.Overlaps(t2))

	// Output:
	// true
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
