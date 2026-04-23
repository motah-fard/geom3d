package geom3d

import (
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
