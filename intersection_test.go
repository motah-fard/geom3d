package geom3d

import "testing"

func TestIntersectRayPlaneHit(t *testing.T) {
	r := Ray3{
		Origin: Vec3{0, 0, 0},
		Dir:    Vec3{0, 0, 1},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	got, ok := IntersectRayPlane(r, pl)
	want := Vec3{0, 0, 5}

	if !ok {
		t.Fatal("expected ray-plane intersection")
	}
	if got != want {
		t.Fatalf("IntersectRayPlane: got %#v, want %#v", got, want)
	}
}

func TestIntersectRayPlaneParallel(t *testing.T) {
	r := Ray3{
		Origin: Vec3{0, 0, 0},
		Dir:    Vec3{1, 0, 0},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	_, ok := IntersectRayPlane(r, pl)
	if ok {
		t.Fatal("expected no intersection for parallel ray")
	}
}

func TestIntersectRayPlaneBehindOrigin(t *testing.T) {
	r := Ray3{
		Origin: Vec3{0, 0, 10},
		Dir:    Vec3{0, 0, 1},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	_, ok := IntersectRayPlane(r, pl)
	if ok {
		t.Fatal("expected no intersection behind ray origin")
	}
}

func TestIntersectSegmentPlaneHit(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{0, 0, 10},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	got, ok := IntersectSegmentPlane(s, pl)
	want := Vec3{0, 0, 5}

	if !ok {
		t.Fatal("expected segment-plane intersection")
	}
	if got != want {
		t.Fatalf("IntersectSegmentPlane: got %#v, want %#v", got, want)
	}
}

func TestIntersectSegmentPlaneOutsideSegment(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{0, 0, 2},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	_, ok := IntersectSegmentPlane(s, pl)
	if ok {
		t.Fatal("expected no segment-plane intersection outside segment")
	}
}

func TestIntersectSegmentPlaneParallel(t *testing.T) {
	s := Segment3{
		A: Vec3{0, 0, 1},
		B: Vec3{1, 0, 1},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	_, ok := IntersectSegmentPlane(s, pl)
	if ok {
		t.Fatal("expected no intersection for parallel segment")
	}
}

func TestIntersectRayAABBHit(t *testing.T) {
	r := Ray3{
		Origin: Vec3{-1, 0.5, 0.5},
		Dir:    Vec3{1, 0, 0},
	}
	b := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{1, 1, 1},
	}

	if !IntersectRayAABB(r, b) {
		t.Fatal("expected ray to intersect AABB")
	}
}

func TestIntersectRayAABBMiss(t *testing.T) {
	r := Ray3{
		Origin: Vec3{-1, 2, 0.5},
		Dir:    Vec3{1, 0, 0},
	}
	b := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{1, 1, 1},
	}

	if IntersectRayAABB(r, b) {
		t.Fatal("expected ray to miss AABB")
	}
}

func TestIntersectRayAABBStartsInside(t *testing.T) {
	r := Ray3{
		Origin: Vec3{0.5, 0.5, 0.5},
		Dir:    Vec3{1, 1, 0},
	}
	b := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{1, 1, 1},
	}

	if !IntersectRayAABB(r, b) {
		t.Fatal("expected ray starting inside AABB to intersect")
	}
}

func TestIntersectRayAABBParallelOutside(t *testing.T) {
	r := Ray3{
		Origin: Vec3{2, 0.5, 0.5},
		Dir:    Vec3{0, 1, 0},
	}
	b := AABB{
		Min: Vec3{0, 0, 0},
		Max: Vec3{1, 1, 1},
	}

	if IntersectRayAABB(r, b) {
		t.Fatal("expected parallel ray outside slab to miss AABB")
	}
}
