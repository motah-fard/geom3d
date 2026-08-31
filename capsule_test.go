package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestCapsuleIsValid(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	if !c.IsValid() {
		t.Fatal("expected valid capsule")
	}

	bad := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: -1}
	if bad.IsValid() {
		t.Fatal("expected invalid capsule for negative radius")
	}
}

func TestCapsuleIsDegenerate(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 0}
	if !c.IsDegenerate() {
		t.Fatal("expected zero-radius capsule to be degenerate")
	}

	c2 := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	if c2.IsDegenerate() {
		t.Fatal("expected non-zero-radius capsule to be non-degenerate")
	}
}

func TestCapsuleSegment(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	got := c.Segment()
	want := Segment3{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}}

	if got != want {
		t.Fatalf("Segment: got %#v, want %#v", got, want)
	}
}

func TestCapsuleVolume(t *testing.T) {
	// Cylinder height 4, radius 1: cylinder volume = pi*r^2*h = 4*pi;
	// end caps together form one sphere: (4/3)*pi*r^3 = 4/3*pi.
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	want := 4*math.Pi + (4.0/3.0)*math.Pi

	if got := c.Volume(); !AlmostEqual(got, want) {
		t.Fatalf("Volume: got %v, want %v", got, want)
	}

	bad := Capsule{Radius: -1}
	if got := bad.Volume(); got != 0 {
		t.Fatalf("Volume for invalid capsule: got %v, want 0", got)
	}
}

func TestCapsuleSurfaceArea(t *testing.T) {
	// Cylinder height 4, radius 1: lateral area = 2*pi*r*h = 8*pi;
	// end caps together form one sphere: 4*pi*r^2 = 4*pi.
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	want := 8*math.Pi + 4*math.Pi

	if got := c.SurfaceArea(); !AlmostEqual(got, want) {
		t.Fatalf("SurfaceArea: got %v, want %v", got, want)
	}

	bad := Capsule{Radius: -1}
	if got := bad.SurfaceArea(); got != 0 {
		t.Fatalf("SurfaceArea for invalid capsule: got %v, want 0", got)
	}
}

func TestCapsuleContains(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}

	if !c.Contains(Vec3{0, 0, 2}) {
		t.Fatal("expected point on the core axis to be contained")
	}
	if !c.Contains(Vec3{0.5, 0, -0.5}) {
		t.Fatal("expected point near the end cap to be contained")
	}
	if c.Contains(Vec3{2, 0, 2}) {
		t.Fatal("expected point outside the capsule")
	}

	bad := Capsule{Radius: -1}
	if bad.Contains(Vec3{0, 0, 0}) {
		t.Fatal("expected Contains to be false for invalid capsule")
	}
}

func TestClosestPointOnCapsuleOutsideCylinder(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	p := Vec3{3, 0, 2}

	got := ClosestPointOnCapsule(p, c)
	want := Vec3{1, 0, 2}

	if got != want {
		t.Fatalf("ClosestPointOnCapsule outside cylinder: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnCapsuleBeyondEndCap(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	p := Vec3{0, 0, 6}

	got := ClosestPointOnCapsule(p, c)
	want := Vec3{0, 0, 5}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ClosestPointOnCapsule beyond end cap: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnCapsuleInside(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	p := Vec3{0.2, 0, 2}

	got := ClosestPointOnCapsule(p, c)
	if got != p {
		t.Fatalf("ClosestPointOnCapsule inside: got %#v, want %#v", got, p)
	}
}

func TestClosestPointOnCapsuleInvalid(t *testing.T) {
	bad := Capsule{Radius: -1}
	got := ClosestPointOnCapsule(Vec3{1, 1, 1}, bad)
	want := Vec3{}

	if got != want {
		t.Fatalf("ClosestPointOnCapsule invalid: got %#v, want %#v", got, want)
	}
}

func TestDistancePointToCapsule(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}

	if got, want := DistancePointToCapsule(Vec3{3, 0, 2}, c), 2.0; !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToCapsule outside: got %v, want %v", got, want)
	}
	if got, want := DistancePointToCapsule(Vec3{0.2, 0, 2}, c), 0.0; !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToCapsule inside: got %v, want %v", got, want)
	}

	bad := Capsule{Radius: -1}
	if got := DistancePointToCapsule(Vec3{0, 0, 0}, bad); got != 0 {
		t.Fatalf("DistancePointToCapsule invalid: got %v, want 0", got)
	}
}

func ExampleClosestPointOnCapsule() {
	c := Capsule{A: Vec3{X: 0, Y: 0, Z: 0}, B: Vec3{X: 0, Y: 0, Z: 4}, Radius: 1}
	p := Vec3{X: 3, Y: 0, Z: 2}

	fmt.Println(ClosestPointOnCapsule(p, c))

	// Output:
	// {1 0 2}
}
