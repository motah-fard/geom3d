package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestOBBIsValid(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 2, 3}, Orientation: IdentityMat3()}
	if !box.IsValid() {
		t.Fatal("expected valid OBB")
	}

	bad := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{-1, 2, 3}, Orientation: IdentityMat3()}
	if bad.IsValid() {
		t.Fatal("expected invalid OBB for negative half-extent")
	}
}

func TestOBBVolume(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 2, 3}, Orientation: IdentityMat3()}
	if got, want := box.Volume(), 48.0; !AlmostEqual(got, want) {
		t.Fatalf("Volume: got %v, want %v", got, want)
	}

	bad := OBB{HalfExtents: Vec3{-1, 2, 3}}
	if got := bad.Volume(); got != 0 {
		t.Fatalf("Volume for invalid OBB: got %v, want 0", got)
	}
}

func TestOBBSurfaceArea(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 2, 3}, Orientation: IdentityMat3()}
	want := 8 * (1*2.0 + 2*3.0 + 3*1.0)
	if got := box.SurfaceArea(); !AlmostEqual(got, want) {
		t.Fatalf("SurfaceArea: got %v, want %v", got, want)
	}
}

func TestOBBSurfaceAreaInvalid(t *testing.T) {
	bad := OBB{HalfExtents: Vec3{-1, 1, 1}}
	if got := bad.SurfaceArea(); got != 0 {
		t.Fatalf("SurfaceArea for invalid OBB: got %v, want 0", got)
	}
}

func TestOBBContainsAxisAligned(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 1, 1}, Orientation: IdentityMat3()}

	if !box.Contains(Vec3{0.5, 0.5, 0.5}) {
		t.Fatal("expected point inside axis-aligned OBB")
	}
	if box.Contains(Vec3{2, 0, 0}) {
		t.Fatal("expected point outside axis-aligned OBB")
	}

	bad := OBB{HalfExtents: Vec3{-1, 1, 1}}
	if bad.Contains(Vec3{0, 0, 0}) {
		t.Fatal("expected Contains to be false for invalid OBB")
	}
}

func TestOBBContainsRotated(t *testing.T) {
	// A box rotated 45 degrees about Z, so a point that would be outside an
	// axis-aligned box at the same extents is inside the rotated one, and
	// vice versa.
	box := OBB{
		Center:      Vec3{0, 0, 0},
		HalfExtents: Vec3{1, 1, 1},
		Orientation: RotationZ(math.Pi / 4),
	}

	// (1.3, 0, 0) is outside an axis-aligned unit box, but the box's local
	// +X axis now points toward (1,1,0)/sqrt(2), so a point further out
	// along the original X axis can still fall inside the rotated box.
	if !box.Contains(Vec3{1.3, 0, 0}) {
		t.Fatal("expected point to be inside the rotated OBB")
	}
	if box.Contains(Vec3{0, 0, 3}) {
		t.Fatal("expected point far along Z to be outside the OBB")
	}
}

func TestClosestPointOnOBBOutside(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 1, 1}, Orientation: IdentityMat3()}
	p := Vec3{3, 0, 0}

	got := ClosestPointOnOBB(p, box)
	want := Vec3{1, 0, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ClosestPointOnOBB outside: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnOBBInside(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 1, 1}, Orientation: IdentityMat3()}
	p := Vec3{0.5, 0.5, 0.5}

	got := ClosestPointOnOBB(p, box)

	if got != p {
		t.Fatalf("ClosestPointOnOBB inside: got %#v, want %#v", got, p)
	}
}

func TestClosestPointOnOBBRotated(t *testing.T) {
	box := OBB{
		Center:      Vec3{0, 0, 0},
		HalfExtents: Vec3{1, 1, 1},
		Orientation: RotationZ(math.Pi / 2),
	}
	// After a 90-degree rotation about Z, the box's local +X axis points
	// along world +Y, so a point far out along world +Y should clamp back
	// to the box's local X extent, landing at world (0, 1, 0).
	p := Vec3{0, 5, 0}

	got := ClosestPointOnOBB(p, box)
	want := Vec3{0, 1, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ClosestPointOnOBB rotated: got %#v, want %#v", got, want)
	}
}

func TestClosestPointOnOBBInvalid(t *testing.T) {
	bad := OBB{HalfExtents: Vec3{-1, 1, 1}}
	got := ClosestPointOnOBB(Vec3{1, 1, 1}, bad)
	want := Vec3{}

	if got != want {
		t.Fatalf("ClosestPointOnOBB invalid: got %#v, want %#v", got, want)
	}
}

func TestDistancePointToOBB(t *testing.T) {
	box := OBB{Center: Vec3{0, 0, 0}, HalfExtents: Vec3{1, 1, 1}, Orientation: IdentityMat3()}

	if got, want := DistancePointToOBB(Vec3{3, 0, 0}, box), 2.0; !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToOBB outside: got %v, want %v", got, want)
	}
	if got, want := DistancePointToOBB(Vec3{0.5, 0.5, 0.5}, box), 0.0; !AlmostEqual(got, want) {
		t.Fatalf("DistancePointToOBB inside: got %v, want %v", got, want)
	}

	bad := OBB{HalfExtents: Vec3{-1, 1, 1}}
	if got := DistancePointToOBB(Vec3{0, 0, 0}, bad); got != 0 {
		t.Fatalf("DistancePointToOBB invalid: got %v, want 0", got)
	}
}

func ExampleClosestPointOnOBB() {
	box := OBB{
		Center:      Vec3{X: 0, Y: 0, Z: 0},
		HalfExtents: Vec3{X: 1, Y: 1, Z: 1},
		Orientation: IdentityMat3(),
	}

	fmt.Println(ClosestPointOnOBB(Vec3{X: 3, Y: 0, Z: 0}, box))

	// Output:
	// {1 0 0}
}
