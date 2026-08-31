package geom3d

import (
	"fmt"
	"testing"
)

func TestAABBIsValid(t *testing.T) {
	box := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{1, 2, 3},
	}
	if !box.IsValid() {
		t.Fatal("expected valid AABB")
	}

	bad := AABB{
		Min: Vec3{2, 0, 0},
		Max: Vec3{1, 2, 3},
	}
	if bad.IsValid() {
		t.Fatal("expected invalid AABB")
	}
}

func TestAABBContains(t *testing.T) {
	box := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{10, 10, 10},
	}

	if !box.Contains(Vec3{5, 5, 5}) {
		t.Fatal("expected point inside box")
	}
	if box.Contains(Vec3{11, 5, 5}) {
		t.Fatal("expected point outside box")
	}
}

func TestAABBOverlaps(t *testing.T) {
	a := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{2, 2, 2},
	}
	b := AABB{
		Min: Vec3{1, 1, 1},
		Max: Vec3{3, 3, 3},
	}
	c := AABB{
		Min: Vec3{2.1, 2.1, 2.1},
		Max: Vec3{4, 4, 4},
	}

	if !a.Overlaps(b) {
		t.Fatal("expected overlap")
	}
	if a.Overlaps(c) {
		t.Fatal("expected no overlap")
	}
}

func TestAABBVolume(t *testing.T) {
	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 3, 4}}
	if got, want := box.Volume(), 24.0; !AlmostEqual(got, want) {
		t.Fatalf("Volume: got %v, want %v", got, want)
	}

	bad := AABB{Min: Vec3{2, 0, 0}, Max: Vec3{0, 0, 0}}
	if got := bad.Volume(); got != 0 {
		t.Fatalf("Volume for invalid box: got %v, want 0", got)
	}
}

func TestAABBSurfaceArea(t *testing.T) {
	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 3, 4}}
	want := 2 * (2*3 + 3*4 + 4*2)
	if got := box.SurfaceArea(); !AlmostEqual(got, float64(want)) {
		t.Fatalf("SurfaceArea: got %v, want %v", got, want)
	}

	bad := AABB{Min: Vec3{2, 0, 0}, Max: Vec3{0, 0, 0}}
	if got := bad.SurfaceArea(); got != 0 {
		t.Fatalf("SurfaceArea for invalid box: got %v, want 0", got)
	}
}

func TestAABBUnion(t *testing.T) {
	a := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 2, 2}}
	b := AABB{Min: Vec3{1, -1, 1}, Max: Vec3{3, 1, 3}}

	got := a.Union(b)
	want := AABB{Min: Vec3{0, -1, 0}, Max: Vec3{3, 2, 3}}

	if got != want {
		t.Fatalf("Union: got %#v, want %#v", got, want)
	}
}

func TestAABBExpandToInclude(t *testing.T) {
	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 2, 2}}

	got := box.ExpandToInclude(Vec3{5, -1, 1})
	want := AABB{Min: Vec3{0, -1, 0}, Max: Vec3{5, 2, 2}}

	if got != want {
		t.Fatalf("ExpandToInclude: got %#v, want %#v", got, want)
	}

	// A point already inside the box should not change it.
	unchanged := box.ExpandToInclude(Vec3{1, 1, 1})
	if unchanged != box {
		t.Fatalf("ExpandToInclude with interior point: got %#v, want %#v", unchanged, box)
	}
}

func TestAABBExpand(t *testing.T) {
	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 2, 2}}

	got := box.Expand(1)
	want := AABB{Min: Vec3{-1, -1, -1}, Max: Vec3{3, 3, 3}}

	if got != want {
		t.Fatalf("Expand: got %#v, want %#v", got, want)
	}
}

func TestAABBFromPoints(t *testing.T) {
	points := []Vec3{
		{1, 5, 0},
		{-2, 1, 3},
		{4, -1, -1},
	}

	got := AABBFromPoints(points)
	want := AABB{Min: Vec3{-2, -1, -1}, Max: Vec3{4, 5, 3}}

	if got != want {
		t.Fatalf("AABBFromPoints: got %#v, want %#v", got, want)
	}
}

func TestAABBFromPointsEmpty(t *testing.T) {
	got := AABBFromPoints(nil)
	want := AABB{}

	if got != want {
		t.Fatalf("AABBFromPoints empty: got %#v, want %#v", got, want)
	}
}

func ExampleAABBFromPoints() {
	points := []Vec3{
		{X: 1, Y: 5, Z: 0},
		{X: -2, Y: 1, Z: 3},
		{X: 4, Y: -1, Z: -1},
	}

	box := AABBFromPoints(points)
	fmt.Println(box.Min)
	fmt.Println(box.Max)

	// Output:
	// {-2 -1 -1}
	// {4 5 3}
}

func ExampleAABB_Union() {
	a := AABB{Min: Vec3{X: 0, Y: 0, Z: 0}, Max: Vec3{X: 2, Y: 2, Z: 2}}
	b := AABB{Min: Vec3{X: 1, Y: -1, Z: 1}, Max: Vec3{X: 3, Y: 1, Z: 3}}

	u := a.Union(b)
	fmt.Println(u.Min)
	fmt.Println(u.Max)

	// Output:
	// {0 -1 0}
	// {3 2 3}
}
