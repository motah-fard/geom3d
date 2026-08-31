package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestSphereIsValid(t *testing.T) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	if !s.IsValid() {
		t.Fatal("expected valid sphere")
	}

	bad := Sphere{Center: Vec3{0, 0, 0}, Radius: -1}
	if bad.IsValid() {
		t.Fatal("expected invalid sphere for negative radius")
	}
}

func TestSphereIsDegenerate(t *testing.T) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 0}
	if !s.IsDegenerate() {
		t.Fatal("expected zero-radius sphere to be degenerate")
	}

	s2 := Sphere{Center: Vec3{0, 0, 0}, Radius: 1}
	if s2.IsDegenerate() {
		t.Fatal("expected non-zero-radius sphere to be non-degenerate")
	}
}

func TestSphereSurfaceArea(t *testing.T) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	want := 4 * math.Pi * 4
	if !AlmostEqual(s.SurfaceArea(), want) {
		t.Fatalf("SurfaceArea: got %v, want %v", s.SurfaceArea(), want)
	}

	bad := Sphere{Radius: -1}
	if bad.SurfaceArea() != 0 {
		t.Fatal("expected 0 surface area for invalid sphere")
	}
}

func TestSphereVolume(t *testing.T) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 3}
	want := (4.0 / 3.0) * math.Pi * 27
	if !AlmostEqual(s.Volume(), want) {
		t.Fatalf("Volume: got %v, want %v", s.Volume(), want)
	}

	bad := Sphere{Radius: -1}
	if bad.Volume() != 0 {
		t.Fatal("expected 0 volume for invalid sphere")
	}
}

func TestSphereContains(t *testing.T) {
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 5}

	if !s.Contains(Vec3{3, 0, 0}) {
		t.Fatal("expected point inside sphere")
	}
	if !s.Contains(Vec3{5, 0, 0}) {
		t.Fatal("expected point on sphere surface to be contained")
	}
	if s.Contains(Vec3{6, 0, 0}) {
		t.Fatal("expected point outside sphere")
	}

	bad := Sphere{Center: Vec3{0, 0, 0}, Radius: -1}
	if bad.Contains(Vec3{0, 0, 0}) {
		t.Fatal("expected Contains to be false for invalid sphere")
	}
}

func TestSphereOverlaps(t *testing.T) {
	a := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	b := Sphere{Center: Vec3{3, 0, 0}, Radius: 2}
	c := Sphere{Center: Vec3{10, 0, 0}, Radius: 2}

	if !a.Overlaps(b) {
		t.Fatal("expected overlapping spheres")
	}
	if a.Overlaps(c) {
		t.Fatal("expected non-overlapping spheres")
	}

	bad := Sphere{Center: Vec3{0, 0, 0}, Radius: -1}
	if bad.Overlaps(a) {
		t.Fatal("expected Overlaps to be false for invalid sphere")
	}
}

func TestSphereOverlapsTouching(t *testing.T) {
	a := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	b := Sphere{Center: Vec3{4, 0, 0}, Radius: 2}

	if !a.Overlaps(b) {
		t.Fatal("expected touching spheres to count as overlapping")
	}
}

func ExampleSphere_Volume() {
	s := Sphere{Center: Vec3{X: 0, Y: 0, Z: 0}, Radius: 1}
	fmt.Printf("%.4f\n", s.Volume())

	// Output:
	// 4.1888
}
