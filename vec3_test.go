package geom3d

import "testing"

func TestVec3Add(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	got := a.Add(b)
	want := Vec3{5, 7, 9}

	if got != want {
		t.Fatalf("Add: got %#v, want %#v", got, want)
	}
}

func TestVec3Sub(t *testing.T) {
	a := Vec3{5, 7, 9}
	b := Vec3{1, 2, 3}
	got := a.Sub(b)
	want := Vec3{4, 5, 6}

	if got != want {
		t.Fatalf("Sub: got %#v, want %#v", got, want)
	}
}

func TestVec3Scale(t *testing.T) {
	a := Vec3{1, -2, 3}
	got := a.Scale(2)
	want := Vec3{2, -4, 6}

	if got != want {
		t.Fatalf("Scale: got %#v, want %#v", got, want)
	}
}

func TestVec3Dot(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	got := a.Dot(b)
	want := 32.0

	if !AlmostEqual(got, want) {
		t.Fatalf("Dot: got %v, want %v", got, want)
	}
}

func TestVec3Cross(t *testing.T) {
	a := Vec3{1, 0, 0}
	b := Vec3{0, 1, 0}
	got := a.Cross(b)
	want := Vec3{0, 0, 1}

	if got != want {
		t.Fatalf("Cross: got %#v, want %#v", got, want)
	}
}

func TestVec3Norm(t *testing.T) {
	a := Vec3{3, 4, 0}
	got := a.Norm()
	want := 5.0

	if !AlmostEqual(got, want) {
		t.Fatalf("Norm: got %v, want %v", got, want)
	}
}

func TestVec3Distance(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 6, 3}
	got := a.Distance(b)
	want := 5.0

	if !AlmostEqual(got, want) {
		t.Fatalf("Distance: got %v, want %v", got, want)
	}
}

func TestVec3Normalize(t *testing.T) {
	a := Vec3{3, 0, 4}
	got := a.Normalize()
	want := Vec3{0.6, 0, 0.8}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("Normalize: got %#v, want %#v", got, want)
	}
}

func TestVec3NormalizeZero(t *testing.T) {
	a := Vec3{}
	got := a.Normalize()
	want := Vec3{}

	if got != want {
		t.Fatalf("Normalize zero: got %#v, want %#v", got, want)
	}
}
