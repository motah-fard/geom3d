package geom3d

import (
	"fmt"
	"math"
	"testing"
)

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

	hit, tMin, tMax := IntersectRayAABB(r, b)
	if !hit {
		t.Fatal("expected ray to intersect AABB")
	}
	if !AlmostEqual(tMin, 1) || !AlmostEqual(tMax, 2) {
		t.Fatalf("IntersectRayAABB: got tMin=%v, tMax=%v, want 1 and 2", tMin, tMax)
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

	hit, _, _ := IntersectRayAABB(r, b)
	if hit {
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

	hit, tMin, tMax := IntersectRayAABB(r, b)
	if !hit {
		t.Fatal("expected ray starting inside AABB to intersect")
	}
	if !AlmostEqual(tMin, 0) {
		t.Fatalf("expected tMin = 0 for ray starting inside box, got %v", tMin)
	}
	if tMax <= 0 {
		t.Fatalf("expected positive tMax for ray starting inside box, got %v", tMax)
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

	hit, _, _ := IntersectRayAABB(r, b)
	if hit {
		t.Fatal("expected parallel ray outside slab to miss AABB")
	}
}

func TestIntersectRayOBBAxisAlignedMatchesAABB(t *testing.T) {
	// With an identity orientation, IntersectRayOBB must agree exactly
	// with IntersectRayAABB on an equivalent box.
	r := Ray3{Origin: Vec3{-1, 0.5, 0.5}, Dir: Vec3{1, 0, 0}}
	box := OBB{
		Center:      Vec3{0.5, 0.5, 0.5},
		HalfExtents: Vec3{0.5, 0.5, 0.5},
		Orientation: IdentityMat3(),
	}

	hit, tMin, tMax := IntersectRayOBB(r, box)
	if !hit {
		t.Fatal("expected ray to intersect axis-aligned OBB")
	}
	if !AlmostEqual(tMin, 1) || !AlmostEqual(tMax, 2) {
		t.Fatalf("IntersectRayOBB: got tMin=%v, tMax=%v, want 1 and 2", tMin, tMax)
	}
}

func TestIntersectRayOBBRotatedHit(t *testing.T) {
	// A box with non-uniform extents, rotated 90 degrees about Z: its
	// local X axis (half-extent 2) now points along world Y, and its
	// local Y axis (half-extent 1) now points along world X. A ray along
	// world X should therefore see a world-X half-extent of 1, entering
	// at x=-1 (t=4) and exiting at x=1 (t=6).
	box := OBB{
		Center:      Vec3{0, 0, 0},
		HalfExtents: Vec3{2, 1, 1},
		Orientation: RotationZ(math.Pi / 2),
	}
	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}

	hit, tMin, tMax := IntersectRayOBB(r, box)
	if !hit {
		t.Fatal("expected ray to intersect rotated OBB")
	}
	if !AlmostEqual(tMin, 4) || !AlmostEqual(tMax, 6) {
		t.Fatalf("IntersectRayOBB rotated: got tMin=%v, tMax=%v, want 4 and 6", tMin, tMax)
	}
}

func TestIntersectRayOBBMiss(t *testing.T) {
	box := OBB{
		Center:      Vec3{0, 0, 0},
		HalfExtents: Vec3{1, 1, 1},
		Orientation: IdentityMat3(),
	}
	r := Ray3{Origin: Vec3{-5, 10, 0.5}, Dir: Vec3{1, 0, 0}}

	hit, _, _ := IntersectRayOBB(r, box)
	if hit {
		t.Fatal("expected ray to miss OBB")
	}
}

func TestIntersectRayOBBInvalid(t *testing.T) {
	box := OBB{HalfExtents: Vec3{1, 1, 1}, Orientation: IdentityMat3()}
	badRay := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{0, 0, 0}}
	if hit, _, _ := IntersectRayOBB(badRay, box); hit {
		t.Fatal("expected no intersection for invalid ray")
	}

	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}
	badBox := OBB{HalfExtents: Vec3{-1, 1, 1}, Orientation: IdentityMat3()}
	if hit, _, _ := IntersectRayOBB(r, badBox); hit {
		t.Fatal("expected no intersection for invalid OBB")
	}
}

func ExampleIntersectRayOBB() {
	box := OBB{
		Center:      Vec3{X: 0, Y: 0, Z: 0},
		HalfExtents: Vec3{X: 2, Y: 1, Z: 1},
		Orientation: RotationZ(math.Pi / 2),
	}
	r := Ray3{Origin: Vec3{X: -5, Y: 0, Z: 0}, Dir: Vec3{X: 1, Y: 0, Z: 0}}

	hit, tMin, tMax := IntersectRayOBB(r, box)
	fmt.Println(hit, tMin, tMax)

	// Output:
	// true 4 6
}

func TestIntersectSegmentsCrossing(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, -1, 0},
		B: Vec3{1, 1, 0},
	}

	got, ok := IntersectSegments(s1, s2)
	want := Vec3{1, 0, 0}

	if !ok {
		t.Fatal("expected segments to intersect")
	}
	if got != want {
		t.Fatalf("IntersectSegments crossing: got %#v, want %#v", got, want)
	}
}

func TestIntersectSegmentsEndpointTouch(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{1, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 0, 0},
		B: Vec3{1, 1, 0},
	}

	got, ok := IntersectSegments(s1, s2)
	want := Vec3{1, 0, 0}

	if !ok {
		t.Fatal("expected endpoint-touching segments to intersect")
	}
	if got != want {
		t.Fatalf("IntersectSegments endpoint touch: got %#v, want %#v", got, want)
	}
}

func TestIntersectSegmentsParallelDisjoint(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{0, 1, 0},
		B: Vec3{2, 1, 0},
	}

	_, ok := IntersectSegments(s1, s2)
	if ok {
		t.Fatal("expected parallel disjoint segments not to intersect")
	}
}

func TestIntersectSegmentsSkewDisjoint(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{1, 1, 1},
		B: Vec3{1, 1, -1},
	}

	_, ok := IntersectSegments(s1, s2)
	if ok {
		t.Fatal("expected skew disjoint segments not to intersect")
	}
}

func TestIntersectSegmentsOverlappingCollinear(t *testing.T) {
	s1 := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
	}
	s2 := Segment3{
		A: Vec3{2, 0, 0},
		B: Vec3{6, 0, 0},
	}

	_, ok := IntersectSegments(s1, s2)
	if ok {
		t.Fatal("expected overlapping collinear segments to return false for non-unique intersection")
	}
}

func TestIntersectRaySphereHit(t *testing.T) {
	r := Ray3{
		Origin: Vec3{-5, 0, 0},
		Dir:    Vec3{1, 0, 0},
	}
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}

	hit, tMin, tMax := IntersectRaySphere(r, s)
	if !hit {
		t.Fatal("expected ray to intersect sphere")
	}
	if !AlmostEqual(tMin, 3) || !AlmostEqual(tMax, 7) {
		t.Fatalf("IntersectRaySphere: got tMin=%v, tMax=%v, want 3 and 7", tMin, tMax)
	}
}

func TestIntersectRaySphereMiss(t *testing.T) {
	r := Ray3{
		Origin: Vec3{-5, 5, 0},
		Dir:    Vec3{1, 0, 0},
	}
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}

	hit, _, _ := IntersectRaySphere(r, s)
	if hit {
		t.Fatal("expected ray to miss sphere")
	}
}

func TestIntersectRaySphereStartsInside(t *testing.T) {
	r := Ray3{
		Origin: Vec3{0, 0, 0},
		Dir:    Vec3{1, 0, 0},
	}
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}

	hit, tMin, tMax := IntersectRaySphere(r, s)
	if !hit {
		t.Fatal("expected ray starting inside sphere to intersect")
	}
	if !AlmostEqual(tMin, 0) {
		t.Fatalf("expected tMin = 0 for ray starting inside sphere, got %v", tMin)
	}
	if !AlmostEqual(tMax, 2) {
		t.Fatalf("expected tMax = 2, got %v", tMax)
	}
}

func TestIntersectRaySphereBehindOrigin(t *testing.T) {
	r := Ray3{
		Origin: Vec3{5, 0, 0},
		Dir:    Vec3{1, 0, 0},
	}
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}

	hit, _, _ := IntersectRaySphere(r, s)
	if hit {
		t.Fatal("expected no intersection when sphere is behind ray origin")
	}
}

func TestIntersectRaySphereInvalid(t *testing.T) {
	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{0, 0, 0}}
	s := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}

	hit, _, _ := IntersectRaySphere(r, s)
	if hit {
		t.Fatal("expected no intersection for invalid ray")
	}

	r2 := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}
	badSphere := Sphere{Center: Vec3{0, 0, 0}, Radius: -1}

	hit2, _, _ := IntersectRaySphere(r2, badSphere)
	if hit2 {
		t.Fatal("expected no intersection for invalid sphere")
	}
}

func ExampleIntersectRaySphere() {
	r := Ray3{
		Origin: Vec3{X: -5, Y: 0, Z: 0},
		Dir:    Vec3{X: 1, Y: 0, Z: 0},
	}
	s := Sphere{Center: Vec3{X: 0, Y: 0, Z: 0}, Radius: 2}

	hit, tMin, tMax := IntersectRaySphere(r, s)
	fmt.Println(hit, tMin, tMax)

	// Output:
	// true 3 7
}

func TestIntersectRayTriangleHit(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	r := Ray3{
		Origin: Vec3{1, 1, 5},
		Dir:    Vec3{0, 0, -1},
	}

	got, ok := IntersectRayTriangle(r, tri)
	want := Vec3{1, 1, 0}

	if !ok {
		t.Fatal("expected ray-triangle intersection")
	}
	if got != want {
		t.Fatalf("IntersectRayTriangle: got %#v, want %#v", got, want)
	}
}

func TestIntersectRayTriangleBackFace(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	r := Ray3{
		Origin: Vec3{1, 1, -5},
		Dir:    Vec3{0, 0, 1},
	}

	_, ok := IntersectRayTriangle(r, tri)
	if !ok {
		t.Fatal("expected back-face hit to be reported (no culling)")
	}
}

func TestIntersectRayTriangleMissOutsideEdge(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	r := Ray3{
		Origin: Vec3{5, 5, 5},
		Dir:    Vec3{0, 0, -1},
	}

	_, ok := IntersectRayTriangle(r, tri)
	if ok {
		t.Fatal("expected no intersection outside triangle bounds")
	}
}

func TestIntersectRayTriangleParallel(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	r := Ray3{
		Origin: Vec3{1, 1, 5},
		Dir:    Vec3{1, 0, 0},
	}

	_, ok := IntersectRayTriangle(r, tri)
	if ok {
		t.Fatal("expected no intersection for ray parallel to triangle plane")
	}
}

func TestIntersectRayTriangleBehindOrigin(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	r := Ray3{
		Origin: Vec3{1, 1, -5},
		Dir:    Vec3{0, 0, -1},
	}

	_, ok := IntersectRayTriangle(r, tri)
	if ok {
		t.Fatal("expected no intersection behind ray origin")
	}
}

func TestIntersectRayTriangleDegenerate(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{4, 0, 0},
	}
	r := Ray3{
		Origin: Vec3{1, 1, 5},
		Dir:    Vec3{0, 0, -1},
	}

	_, ok := IntersectRayTriangle(r, tri)
	if ok {
		t.Fatal("expected no intersection for degenerate triangle")
	}
}

func TestIntersectSegmentTriangleHit(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	s := Segment3{
		A: Vec3{1, 1, 5},
		B: Vec3{1, 1, -5},
	}

	got, ok := IntersectSegmentTriangle(s, tri)
	want := Vec3{1, 1, 0}

	if !ok {
		t.Fatal("expected segment-triangle intersection")
	}
	if got != want {
		t.Fatalf("IntersectSegmentTriangle: got %#v, want %#v", got, want)
	}
}

func TestIntersectSegmentTriangleBeyondSegment(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	// The triangle's plane is at z=0, but the segment stops at z=1,
	// short of the plane.
	s := Segment3{
		A: Vec3{1, 1, 5},
		B: Vec3{1, 1, 1},
	}

	_, ok := IntersectSegmentTriangle(s, tri)
	if ok {
		t.Fatal("expected no intersection when the triangle is beyond the segment's endpoint")
	}
}

func TestIntersectSegmentTriangleMissOutsideEdge(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{4, 0, 0},
		C: Vec3{0, 4, 0},
	}
	s := Segment3{
		A: Vec3{5, 5, 5},
		B: Vec3{5, 5, -5},
	}

	_, ok := IntersectSegmentTriangle(s, tri)
	if ok {
		t.Fatal("expected no intersection outside triangle bounds")
	}
}

func TestIntersectSegmentTriangleDegenerate(t *testing.T) {
	tri := Triangle{
		A: Vec3{0, 0, 0},
		B: Vec3{2, 0, 0},
		C: Vec3{4, 0, 0},
	}
	s := Segment3{A: Vec3{1, 1, 1}, B: Vec3{1, 1, 1}}

	_, ok := IntersectSegmentTriangle(s, tri)
	if ok {
		t.Fatal("expected no intersection for degenerate segment")
	}
}

func TestIntersectSegmentAABBHit(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	s := Segment3{A: Vec3{-1, 0.5, 0.5}, B: Vec3{2, 0.5, 0.5}}

	hit, tMin, tMax := IntersectSegmentAABB(s, b)
	if !hit {
		t.Fatal("expected segment to intersect AABB")
	}
	// The full segment spans x in [-1, 2] (length 3); the box spans x in
	// [0, 1], entered at t=1/3 and exited at t=2/3.
	if !AlmostEqual(tMin, 1.0/3.0) || !AlmostEqual(tMax, 2.0/3.0) {
		t.Fatalf("IntersectSegmentAABB: got tMin=%v, tMax=%v, want 1/3 and 2/3", tMin, tMax)
	}
}

func TestIntersectSegmentAABBEndsInsideBox(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{2, 2, 2}}
	// The segment ends at x=1, which is inside the box, so the raw ray
	// exit parameter (1.5) must be clamped down to the segment's own
	// endpoint (t=1).
	s := Segment3{A: Vec3{-1, 0.5, 0.5}, B: Vec3{1, 0.5, 0.5}}

	hit, tMin, tMax := IntersectSegmentAABB(s, b)
	if !hit {
		t.Fatal("expected segment ending inside the box to intersect")
	}
	if !AlmostEqual(tMin, 0.5) || !AlmostEqual(tMax, 1) {
		t.Fatalf("IntersectSegmentAABB: got tMin=%v, tMax=%v, want 0.5 and 1", tMin, tMax)
	}
}

func TestIntersectSegmentAABBBeyondSegment(t *testing.T) {
	b := AABB{Min: Vec3{5, 0, 0}, Max: Vec3{6, 1, 1}}
	s := Segment3{A: Vec3{-1, 0.5, 0.5}, B: Vec3{1, 0.5, 0.5}}

	hit, _, _ := IntersectSegmentAABB(s, b)
	if hit {
		t.Fatal("expected no intersection when the box is beyond the segment's endpoint")
	}
}

func TestIntersectSegmentAABBMiss(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	s := Segment3{A: Vec3{-1, 2, 0.5}, B: Vec3{2, 2, 0.5}}

	hit, _, _ := IntersectSegmentAABB(s, b)
	if hit {
		t.Fatal("expected segment to miss AABB")
	}
}

func TestIntersectSegmentAABBDegenerate(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	s := Segment3{A: Vec3{0.5, 0.5, 0.5}, B: Vec3{0.5, 0.5, 0.5}}

	hit, _, _ := IntersectSegmentAABB(s, b)
	if hit {
		t.Fatal("expected no intersection for degenerate segment")
	}
}

func TestIntersectSegmentSphereHit(t *testing.T) {
	sph := Sphere{Center: Vec3{0, 0, 0}, Radius: 2}
	s := Segment3{A: Vec3{-5, 0, 0}, B: Vec3{5, 0, 0}}

	hit, tMin, tMax := IntersectSegmentSphere(s, sph)
	if !hit {
		t.Fatal("expected segment to intersect sphere")
	}
	// The full segment spans x in [-5, 5] (length 10); the sphere is
	// entered at x=-2 (t=0.3) and exited at x=2 (t=0.7).
	if !AlmostEqual(tMin, 0.3) || !AlmostEqual(tMax, 0.7) {
		t.Fatalf("IntersectSegmentSphere: got tMin=%v, tMax=%v, want 0.3 and 0.7", tMin, tMax)
	}
}

func TestIntersectSegmentSphereEndsInsideSphere(t *testing.T) {
	sph := Sphere{Center: Vec3{0, 0, 0}, Radius: 3}
	// The segment ends at x=1, which is inside the sphere, so the raw ray
	// exit parameter (4/3) must be clamped down to the segment's own
	// endpoint (t=1).
	s := Segment3{A: Vec3{-5, 0, 0}, B: Vec3{1, 0, 0}}

	hit, tMin, tMax := IntersectSegmentSphere(s, sph)
	if !hit {
		t.Fatal("expected segment ending inside the sphere to intersect")
	}
	if !AlmostEqual(tMin, 1.0/3.0) || !AlmostEqual(tMax, 1) {
		t.Fatalf("IntersectSegmentSphere: got tMin=%v, tMax=%v, want 1/3 and 1", tMin, tMax)
	}
}

func TestIntersectSegmentSphereBeyondSegment(t *testing.T) {
	sph := Sphere{Center: Vec3{10, 0, 0}, Radius: 2}
	s := Segment3{A: Vec3{-5, 0, 0}, B: Vec3{5, 0, 0}}

	hit, _, _ := IntersectSegmentSphere(s, sph)
	if hit {
		t.Fatal("expected no intersection when the sphere is beyond the segment's endpoint")
	}
}

func TestIntersectSegmentSphereMiss(t *testing.T) {
	sph := Sphere{Center: Vec3{0, 5, 0}, Radius: 2}
	s := Segment3{A: Vec3{-5, 0, 0}, B: Vec3{5, 0, 0}}

	hit, _, _ := IntersectSegmentSphere(s, sph)
	if hit {
		t.Fatal("expected segment to miss sphere")
	}
}

func TestIntersectAABBSphereOverlap(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	s := Sphere{Center: Vec3{2, 0.5, 0.5}, Radius: 1.5}

	if !IntersectAABBSphere(b, s) {
		t.Fatal("expected AABB and sphere to overlap")
	}
}

func TestIntersectAABBSphereMiss(t *testing.T) {
	b := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	s := Sphere{Center: Vec3{10, 0.5, 0.5}, Radius: 1}

	if IntersectAABBSphere(b, s) {
		t.Fatal("expected AABB and sphere not to overlap")
	}
}

func TestIntersectAABBSphereInvalid(t *testing.T) {
	badBox := AABB{Min: Vec3{1, 1, 1}, Max: Vec3{0, 0, 0}}
	s := Sphere{Center: Vec3{0.5, 0.5, 0.5}, Radius: 1}

	if IntersectAABBSphere(badBox, s) {
		t.Fatal("expected false for invalid AABB")
	}

	box := AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}
	badSphere := Sphere{Center: Vec3{0.5, 0.5, 0.5}, Radius: -1}

	if IntersectAABBSphere(box, badSphere) {
		t.Fatal("expected false for invalid sphere")
	}
}

func ExampleIntersectSegmentTriangle() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 4, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 4, Z: 0},
	}
	s := Segment3{
		A: Vec3{X: 1, Y: 1, Z: 5},
		B: Vec3{X: 1, Y: 1, Z: -5},
	}

	p, ok := IntersectSegmentTriangle(s, tri)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {1 1 0}
}

func ExampleIntersectAABBSphere() {
	b := AABB{Min: Vec3{X: 0, Y: 0, Z: 0}, Max: Vec3{X: 1, Y: 1, Z: 1}}
	s := Sphere{Center: Vec3{X: 2, Y: 0.5, Z: 0.5}, Radius: 1.5}

	fmt.Println(IntersectAABBSphere(b, s))

	// Output:
	// true
}

func ExampleIntersectSegmentAABB() {
	b := AABB{Min: Vec3{X: 0, Y: 0, Z: 0}, Max: Vec3{X: 1, Y: 1, Z: 1}}
	s := Segment3{A: Vec3{X: -1, Y: 0.5, Z: 0.5}, B: Vec3{X: 2, Y: 0.5, Z: 0.5}}

	hit, tMin, tMax := IntersectSegmentAABB(s, b)
	fmt.Println(hit, tMin, tMax)

	// Output:
	// true 0.3333333333333333 0.6666666666666666
}

func ExampleIntersectSegmentSphere() {
	sph := Sphere{Center: Vec3{X: 0, Y: 0, Z: 0}, Radius: 2}
	s := Segment3{A: Vec3{X: -5, Y: 0, Z: 0}, B: Vec3{X: 5, Y: 0, Z: 0}}

	hit, tMin, tMax := IntersectSegmentSphere(s, sph)
	fmt.Println(hit, tMin, tMax)

	// Output:
	// true 0.3 0.7
}

func TestIntersectRayCapsuleThroughCylinderBody(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{-5, 0, 2}, Dir: Vec3{1, 0, 0}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	if !hit {
		t.Fatal("expected ray to intersect the capsule's cylindrical body")
	}
	if !AlmostEqual(tMin, 4) || !AlmostEqual(tMax, 6) {
		t.Fatalf("IntersectRayCapsule: got tMin=%v, tMax=%v, want 4 and 6", tMin, tMax)
	}
}

func TestIntersectRayCapsuleAlongAxisThroughBothCaps(t *testing.T) {
	// A ray parallel to the capsule's axis, straight down through the
	// center: it must enter through the top hemisphere (at z=5, since the
	// cap sphere at B=(0,0,4) has radius 1) and exit through the bottom
	// hemisphere (at z=-1). This exercises the degenerate "ray parallel to
	// the cylinder axis" branch, where the cylinder quadratic's leading
	// coefficient is zero and only the two cap spheres contribute.
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{0, 0, 10}, Dir: Vec3{0, 0, -1}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	if !hit {
		t.Fatal("expected ray to intersect the capsule along its axis")
	}
	if !AlmostEqual(tMin, 5) || !AlmostEqual(tMax, 11) {
		t.Fatalf("IntersectRayCapsule along axis: got tMin=%v, tMax=%v, want 5 and 11", tMin, tMax)
	}
}

func TestIntersectRayCapsuleAlongAxisReversed(t *testing.T) {
	// Same axis-aligned setup as the "through both caps" case above, but
	// approaching from below: entry is now via the bottom cap (sphere A)
	// and exit via the top cap (sphere B) — the mirror image of that
	// test's entry-via-B/exit-via-A, exercising the other two candidate
	// branches (sphere-A-as-entry, sphere-B-as-exit).
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{0, 0, -10}, Dir: Vec3{0, 0, 1}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	if !hit {
		t.Fatal("expected ray to intersect the capsule along its axis")
	}
	if !AlmostEqual(tMin, 9) || !AlmostEqual(tMax, 15) {
		t.Fatalf("IntersectRayCapsule along axis reversed: got tMin=%v, tMax=%v, want 9 and 15", tMin, tMax)
	}
}

func TestIntersectRayCapsuleMiss(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{-5, 10, 2}, Dir: Vec3{1, 0, 0}}

	hit, _, _ := IntersectRayCapsule(r, c)
	if hit {
		t.Fatal("expected ray to miss the capsule")
	}
}

func TestIntersectRayCapsuleStartsInside(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{0, 0, 2}, Dir: Vec3{1, 0, 0}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	if !hit {
		t.Fatal("expected ray starting inside the capsule to intersect")
	}
	if !AlmostEqual(tMin, 0) {
		t.Fatalf("expected tMin = 0 for ray starting inside capsule, got %v", tMin)
	}
	if !AlmostEqual(tMax, 1) {
		t.Fatalf("expected tMax = 1, got %v", tMax)
	}
}

func TestIntersectRayCapsuleParallelOutsideRadius(t *testing.T) {
	// Parallel to the axis, but offset beyond the radius: even though the
	// ray never diverges from the axis direction, it should never come
	// within range of either cap sphere.
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	r := Ray3{Origin: Vec3{3, 0, 10}, Dir: Vec3{0, 0, -1}}

	hit, _, _ := IntersectRayCapsule(r, c)
	if hit {
		t.Fatal("expected parallel ray beyond the radius to miss")
	}
}

func TestIntersectRayCapsuleDegenerateToSphere(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 0}, Radius: 2}
	r := Ray3{Origin: Vec3{-5, 0, 0}, Dir: Vec3{1, 0, 0}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	wantHit, wantMin, wantMax := IntersectRaySphere(r, Sphere{Center: Vec3{0, 0, 0}, Radius: 2})

	if hit != wantHit || !AlmostEqual(tMin, wantMin) || !AlmostEqual(tMax, wantMax) {
		t.Fatalf("IntersectRayCapsule degenerate: got (%v,%v,%v), want (%v,%v,%v)", hit, tMin, tMax, wantHit, wantMin, wantMax)
	}
}

func TestIntersectRayCapsuleConsistentWithDistancePointToCapsule(t *testing.T) {
	// Cross-check against the independently implemented
	// DistancePointToCapsule: at the reported entry and exit parameters,
	// the ray should be exactly on the capsule's surface (distance ==
	// radius), and strictly inside just past the entry point.
	cases := []struct {
		name string
		c    Capsule
		r    Ray3
	}{
		{"through cylinder body", Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}, Ray3{Origin: Vec3{-5, 0, 2}, Dir: Vec3{1, 0, 0}}},
		{"along axis through caps", Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}, Ray3{Origin: Vec3{0, 0, 10}, Dir: Vec3{0, 0, -1}}},
		{"diagonal approach", Capsule{A: Vec3{1, -2, 0}, B: Vec3{3, 2, 5}, Radius: 0.7}, Ray3{Origin: Vec3{-10, -2, 1}, Dir: Vec3{1, 0.15, 0.08}}},
	}

	for _, tc := range cases {
		hit, tMin, tMax := IntersectRayCapsule(tc.r, tc.c)
		if !hit {
			t.Fatalf("%s: expected a hit", tc.name)
		}

		entryDist := DistancePointToCapsule(tc.r.PointAt(tMin), tc.c)
		exitDist := DistancePointToCapsule(tc.r.PointAt(tMax), tc.c)
		if !AlmostEqual(entryDist, 0) {
			t.Fatalf("%s: entry point should be on the capsule surface, got distance %v", tc.name, entryDist)
		}
		if !AlmostEqual(exitDist, 0) {
			t.Fatalf("%s: exit point should be on the capsule surface, got distance %v", tc.name, exitDist)
		}

		midT := (tMin + tMax) / 2
		if got := DistancePointToCapsule(tc.r.PointAt(midT), tc.c); got != 0 {
			t.Fatalf("%s: midpoint between entry and exit should be strictly inside the capsule, got distance %v", tc.name, got)
		}
	}
}

func TestIntersectRayCapsuleInvalid(t *testing.T) {
	c := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: 1}
	badRay := Ray3{Origin: Vec3{-5, 0, 2}, Dir: Vec3{0, 0, 0}}
	if hit, _, _ := IntersectRayCapsule(badRay, c); hit {
		t.Fatal("expected no intersection for invalid ray")
	}

	r := Ray3{Origin: Vec3{-5, 0, 2}, Dir: Vec3{1, 0, 0}}
	badCapsule := Capsule{A: Vec3{0, 0, 0}, B: Vec3{0, 0, 4}, Radius: -1}
	if hit, _, _ := IntersectRayCapsule(r, badCapsule); hit {
		t.Fatal("expected no intersection for invalid capsule")
	}
}

func ExampleIntersectRayCapsule() {
	c := Capsule{A: Vec3{X: 0, Y: 0, Z: 0}, B: Vec3{X: 0, Y: 0, Z: 4}, Radius: 1}
	r := Ray3{Origin: Vec3{X: -5, Y: 0, Z: 2}, Dir: Vec3{X: 1, Y: 0, Z: 0}}

	hit, tMin, tMax := IntersectRayCapsule(r, c)
	fmt.Println(hit, tMin, tMax)

	// Output:
	// true 4 6
}

func ExampleIntersectSegmentPlane() {
	s := Segment3{
		A: Vec3{0, 0, 0},
		B: Vec3{0, 0, 10},
	}
	pl := Plane{
		Point:  Vec3{0, 0, 5},
		Normal: Vec3{0, 0, 1},
	}

	p, ok := IntersectSegmentPlane(s, pl)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {0 0 5}
}

func ExampleIntersectRayAABB() {
	box := AABB{
		Min: Vec3{X: 0, Y: 0, Z: 0},
		Max: Vec3{X: 2, Y: 2, Z: 2},
	}

	ray1 := Ray3{
		Origin: Vec3{X: -1, Y: 1, Z: 1},
		Dir:    Vec3{X: 1, Y: 0, Z: 0},
	}

	ray2 := Ray3{
		Origin: Vec3{X: -1, Y: 3, Z: 1},
		Dir:    Vec3{X: 1, Y: 0, Z: 0},
	}

	hit1, tMin1, tMax1 := IntersectRayAABB(ray1, box)
	hit2, _, _ := IntersectRayAABB(ray2, box)

	fmt.Println(hit1, tMin1, tMax1)
	fmt.Println(hit2)

	// Output:
	// true 1 3
	// false
}

func ExampleIntersectRayPlane() {
	r := Ray3{
		Origin: Vec3{X: 0, Y: 0, Z: 0},
		Dir:    Vec3{X: 0, Y: 0, Z: 1},
	}

	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 5},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	p, ok := IntersectRayPlane(r, pl)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {0 0 5}
}

func ExampleIntersectRayPlane_parallel() {
	r := Ray3{
		Origin: Vec3{X: 0, Y: 0, Z: 0},
		Dir:    Vec3{X: 1, Y: 0, Z: 0},
	}

	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 5},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	p, ok := IntersectRayPlane(r, pl)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// false
	// {0 0 0}
}
func ExampleIntersectRayTriangle() {
	tri := Triangle{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 4, Y: 0, Z: 0},
		C: Vec3{X: 0, Y: 4, Z: 0},
	}

	r := Ray3{
		Origin: Vec3{X: 1, Y: 1, Z: 5},
		Dir:    Vec3{X: 0, Y: 0, Z: -1},
	}

	p, ok := IntersectRayTriangle(r, tri)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {1 1 0}
}

func ExampleIntersectSegments() {
	s1 := Segment3{
		A: Vec3{X: 0, Y: 0, Z: 0},
		B: Vec3{X: 2, Y: 0, Z: 0},
	}
	s2 := Segment3{
		A: Vec3{X: 1, Y: -1, Z: 0},
		B: Vec3{X: 1, Y: 1, Z: 0},
	}

	p, ok := IntersectSegments(s1, s2)
	fmt.Println(ok)
	fmt.Println(p)

	// Output:
	// true
	// {1 0 0}
}
