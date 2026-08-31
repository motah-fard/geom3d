package geom3d

import (
	"fmt"
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
