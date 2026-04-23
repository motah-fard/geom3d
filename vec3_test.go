package geom3d

import (
	"fmt"
	"testing"
)

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
func TestVec3Norm2(t *testing.T) {
	a := Vec3{3, 4, 0}
	got := a.Norm2()
	want := 25.0

	if !AlmostEqual(got, want) {
		t.Fatalf("Norm2: got %v, want %v", got, want)
	}
}

func TestVec3Distance2(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 6, 3}
	got := a.Distance2(b)
	want := 25.0

	if !AlmostEqual(got, want) {
		t.Fatalf("Distance2: got %v, want %v", got, want)
	}
}
func TestVec3CrossOrthogonalToInputs(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{4, 5, 6}
	c := a.Cross(b)

	if !AlmostZero(c.Dot(a)) {
		t.Fatalf("cross product should be orthogonal to first input: got %v", c.Dot(a))
	}
	if !AlmostZero(c.Dot(b)) {
		t.Fatalf("cross product should be orthogonal to second input: got %v", c.Dot(b))
	}
}
func ExampleVec3_basicOperations() {
	a := Vec3{X: 1, Y: 2, Z: 3}
	b := Vec3{X: 4, Y: 5, Z: 6}

	sum := a.Add(b)
	diff := a.Sub(b)
	cross := a.Cross(b)

	fmt.Printf("%.0f %.0f %.0f\n", sum.X, sum.Y, sum.Z)
	fmt.Printf("%.0f %.0f %.0f\n", diff.X, diff.Y, diff.Z)
	fmt.Printf("%.0f\n", a.Dot(b))
	fmt.Printf("%.0f %.0f %.0f\n", cross.X, cross.Y, cross.Z)

	// Output:
	// 5 7 9
	// -3 -3 -3
	// 32
	// -3 6 -3
}
func ExampleVec3_Normalize() {
	v := Vec3{X: 3, Y: 0, Z: 4}
	u := v.Normalize()

	fmt.Printf("%.1f %.1f %.1f\n", u.X, u.Y, u.Z)

	// Output:
	// 0.6 0.0 0.8
}
