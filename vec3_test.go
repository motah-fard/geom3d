package geom3d

import (
	"fmt"
	"math"
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

func TestVec3Midpoint(t *testing.T) {
	a := Vec3{0, 0, 0}
	b := Vec3{4, 2, -2}
	got := a.Midpoint(b)
	want := Vec3{2, 1, -1}

	if got != want {
		t.Fatalf("Midpoint: got %#v, want %#v", got, want)
	}
}

func TestVec3Lerp(t *testing.T) {
	a := Vec3{0, 0, 0}
	b := Vec3{4, 8, -4}

	if got, want := a.Lerp(b, 0), a; got != want {
		t.Fatalf("Lerp t=0: got %#v, want %#v", got, want)
	}
	if got, want := a.Lerp(b, 1), b; got != want {
		t.Fatalf("Lerp t=1: got %#v, want %#v", got, want)
	}
	if got, want := a.Lerp(b, 0.5), (Vec3{2, 4, -2}); got != want {
		t.Fatalf("Lerp t=0.5: got %#v, want %#v", got, want)
	}
}

func TestVec3Reflect(t *testing.T) {
	a := Vec3{1, -1, 0}
	n := Vec3{0, 1, 0}
	got := a.Reflect(n)
	want := Vec3{1, 1, 0}

	if got != want {
		t.Fatalf("Reflect: got %#v, want %#v", got, want)
	}
}

func TestVec3Project(t *testing.T) {
	a := Vec3{3, 4, 0}
	b := Vec3{2, 0, 0}
	got := a.Project(b)
	want := Vec3{3, 0, 0}

	if got != want {
		t.Fatalf("Project: got %#v, want %#v", got, want)
	}
}

func TestVec3ProjectOntoZero(t *testing.T) {
	a := Vec3{3, 4, 0}
	got := a.Project(Vec3{})
	want := Vec3{}

	if got != want {
		t.Fatalf("Project onto zero vector: got %#v, want %#v", got, want)
	}
}

func TestVec3Angle(t *testing.T) {
	a := Vec3{1, 0, 0}
	b := Vec3{0, 1, 0}
	got := a.Angle(b)
	want := math.Pi / 2

	if !AlmostEqual(got, want) {
		t.Fatalf("Angle: got %v, want %v", got, want)
	}
}

func TestVec3AngleSameDirection(t *testing.T) {
	a := Vec3{2, 0, 0}
	b := Vec3{5, 0, 0}
	got := a.Angle(b)

	if !AlmostEqual(got, 0) {
		t.Fatalf("Angle for parallel vectors: got %v, want 0", got)
	}
}

func TestVec3AngleWithZeroVector(t *testing.T) {
	a := Vec3{1, 0, 0}
	got := a.Angle(Vec3{})

	if !AlmostEqual(got, 0) {
		t.Fatalf("Angle with zero vector: got %v, want 0", got)
	}
}

func TestVec3ClampLength(t *testing.T) {
	a := Vec3{3, 4, 0}
	got := a.ClampLength(2)
	want := Vec3{1.2, 1.6, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ClampLength: got %#v, want %#v", got, want)
	}
}

func TestVec3ClampLengthUnderLimit(t *testing.T) {
	a := Vec3{1, 0, 0}
	got := a.ClampLength(5)

	if got != a {
		t.Fatalf("ClampLength under limit: got %#v, want %#v", got, a)
	}
}

func TestVec3Abs(t *testing.T) {
	a := Vec3{-1, 2, -3}
	got := a.Abs()
	want := Vec3{1, 2, 3}

	if got != want {
		t.Fatalf("Abs: got %#v, want %#v", got, want)
	}
}

func TestVec3Min(t *testing.T) {
	a := Vec3{1, 5, -3}
	b := Vec3{4, 2, -1}
	got := a.Min(b)
	want := Vec3{1, 2, -3}

	if got != want {
		t.Fatalf("Min: got %#v, want %#v", got, want)
	}
}

func TestVec3Max(t *testing.T) {
	a := Vec3{1, 5, -3}
	b := Vec3{4, 2, -1}
	got := a.Max(b)
	want := Vec3{4, 5, -1}

	if got != want {
		t.Fatalf("Max: got %#v, want %#v", got, want)
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
func ExampleVec3_Lerp() {
	a := Vec3{X: 0, Y: 0, Z: 0}
	b := Vec3{X: 10, Y: 0, Z: 0}

	fmt.Println(a.Lerp(b, 0.25))

	// Output:
	// {2.5 0 0}
}

func ExampleVec3_Reflect() {
	incoming := Vec3{X: 1, Y: -1, Z: 0}
	surfaceNormal := Vec3{X: 0, Y: 1, Z: 0}

	fmt.Println(incoming.Reflect(surfaceNormal))

	// Output:
	// {1 1 0}
}

func ExampleVec3_ClampLength() {
	v := Vec3{X: 3, Y: 4, Z: 0}
	clamped := v.ClampLength(2)

	fmt.Printf("%.2f %.2f %.2f\n", clamped.X, clamped.Y, clamped.Z)

	// Output:
	// 1.20 1.60 0.00
}

func ExampleVec3_Normalize() {
	v := Vec3{X: 3, Y: 0, Z: 4}
	u := v.Normalize()

	fmt.Printf("%.1f %.1f %.1f\n", u.X, u.Y, u.Z)

	// Output:
	// 0.6 0.0 0.8
}
