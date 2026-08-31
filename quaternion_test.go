package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestIdentityQuaternionRotateVector(t *testing.T) {
	v := Vec3{1, 2, 3}
	got := IdentityQuaternion().RotateVector(v)

	if got != v {
		t.Fatalf("IdentityQuaternion RotateVector: got %#v, want %#v", got, v)
	}
}

func TestQuaternionFromAxisAngleRotateVector(t *testing.T) {
	q := QuaternionFromAxisAngle(Vec3{0, 0, 1}, math.Pi/2)
	got := q.RotateVector(Vec3{1, 0, 0})
	want := Vec3{0, 1, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("RotateVector: got %#v, want %#v", got, want)
	}
}

func TestQuaternionFromAxisAngleZeroAxis(t *testing.T) {
	got := QuaternionFromAxisAngle(Vec3{}, math.Pi/2)
	want := IdentityQuaternion()

	if got != want {
		t.Fatalf("QuaternionFromAxisAngle with zero axis: got %#v, want %#v", got, want)
	}
}

func TestQuaternionIsValid(t *testing.T) {
	if !IdentityQuaternion().IsValid() {
		t.Fatal("expected identity quaternion to be valid")
	}
	if (Quaternion{}).IsValid() {
		t.Fatal("expected zero quaternion to be invalid")
	}
}

func TestQuaternionNorm(t *testing.T) {
	q := Quaternion{X: 0, Y: 0, Z: 0, W: 1}
	if got, want := q.Norm(), 1.0; !AlmostEqual(got, want) {
		t.Fatalf("Norm: got %v, want %v", got, want)
	}
}

func TestQuaternionNormalize(t *testing.T) {
	q := Quaternion{X: 0, Y: 0, Z: 0, W: 2}
	got := q.Normalize()
	want := IdentityQuaternion()

	if got != want {
		t.Fatalf("Normalize: got %#v, want %#v", got, want)
	}
}

func TestQuaternionNormalizeZero(t *testing.T) {
	got := (Quaternion{}).Normalize()
	want := Quaternion{}

	if got != want {
		t.Fatalf("Normalize zero quaternion: got %#v, want %#v", got, want)
	}
}

func TestQuaternionConjugate(t *testing.T) {
	q := Quaternion{X: 1, Y: 2, Z: 3, W: 4}
	got := q.Conjugate()
	want := Quaternion{X: -1, Y: -2, Z: -3, W: 4}

	if got != want {
		t.Fatalf("Conjugate: got %#v, want %#v", got, want)
	}
}

func TestQuaternionInverse(t *testing.T) {
	q := QuaternionFromAxisAngle(Vec3{1, 1, 0}, 0.7)
	inv := q.Inverse()

	got := q.Mul(inv)
	want := IdentityQuaternion()

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) ||
		!AlmostEqual(got.Z, want.Z) || !AlmostEqual(got.W, want.W) {
		t.Fatalf("q.Mul(q.Inverse()): got %#v, want %#v", got, want)
	}
}

func TestQuaternionInverseZero(t *testing.T) {
	got := (Quaternion{}).Inverse()
	want := Quaternion{}

	if got != want {
		t.Fatalf("Inverse of zero quaternion: got %#v, want %#v", got, want)
	}
}

func TestQuaternionMulComposesLikeMat3(t *testing.T) {
	q1 := QuaternionFromAxisAngle(Vec3{0, 0, 1}, math.Pi/2)
	q2 := QuaternionFromAxisAngle(Vec3{1, 0, 0}, math.Pi/2)
	v := Vec3{0, 1, 0}

	// q1.Mul(q2) should apply q2 first, then q1 — the same order as
	// Mat3.Mul and Transform.Compose.
	composed := q1.Mul(q2).RotateVector(v)
	sequential := q1.RotateVector(q2.RotateVector(v))

	if !AlmostEqual(composed.X, sequential.X) || !AlmostEqual(composed.Y, sequential.Y) || !AlmostEqual(composed.Z, sequential.Z) {
		t.Fatalf("Mul composition order: got %#v, want %#v", composed, sequential)
	}
}

func TestQuaternionRotateVectorMatchesMat3(t *testing.T) {
	axis := Vec3{1, 2, 3}
	angle := 0.9
	q := QuaternionFromAxisAngle(axis, angle)
	m := q.ToMat3()

	v := Vec3{4, -1, 2}
	got := q.RotateVector(v)
	want := m.MulVec(v)

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("RotateVector vs ToMat3: got %#v, want %#v", got, want)
	}
}

func TestQuaternionRotateVectorInvalid(t *testing.T) {
	got := (Quaternion{}).RotateVector(Vec3{1, 2, 3})
	want := Vec3{}

	if got != want {
		t.Fatalf("RotateVector with invalid quaternion: got %#v, want %#v", got, want)
	}
}

func TestQuaternionToMat3Zero(t *testing.T) {
	got := (Quaternion{}).ToMat3()
	want := IdentityMat3()

	if got != want {
		t.Fatalf("ToMat3 of zero quaternion: got %#v, want %#v", got, want)
	}
}

func TestQuaternionSlerpEndpoints(t *testing.T) {
	q := IdentityQuaternion()
	other := QuaternionFromAxisAngle(Vec3{0, 1, 0}, math.Pi/2)

	start := q.Slerp(other, 0)
	if !AlmostEqual(start.X, q.X) || !AlmostEqual(start.Y, q.Y) ||
		!AlmostEqual(start.Z, q.Z) || !AlmostEqual(start.W, q.W) {
		t.Fatalf("Slerp t=0: got %#v, want %#v", start, q)
	}

	end := q.Slerp(other, 1)
	if !AlmostEqual(end.X, other.X) || !AlmostEqual(end.Y, other.Y) ||
		!AlmostEqual(end.Z, other.Z) || !AlmostEqual(end.W, other.W) {
		t.Fatalf("Slerp t=1: got %#v, want %#v", end, other)
	}
}

func TestQuaternionSlerpHalfway(t *testing.T) {
	q := IdentityQuaternion()
	other := QuaternionFromAxisAngle(Vec3{0, 1, 0}, math.Pi/2)

	mid := q.Slerp(other, 0.5)
	want := QuaternionFromAxisAngle(Vec3{0, 1, 0}, math.Pi/4)

	if !AlmostEqual(mid.X, want.X) || !AlmostEqual(mid.Y, want.Y) ||
		!AlmostEqual(mid.Z, want.Z) || !AlmostEqual(mid.W, want.W) {
		t.Fatalf("Slerp t=0.5: got %#v, want %#v", mid, want)
	}
}

func TestQuaternionSlerpShortestPath(t *testing.T) {
	q := QuaternionFromAxisAngle(Vec3{0, 1, 0}, math.Pi/3)
	// The antipodal quaternion represents the exact same rotation as q, but
	// has a negative dot product with it, forcing Slerp to take the
	// "shorter path" branch.
	antipodal := Quaternion{X: -q.X, Y: -q.Y, Z: -q.Z, W: -q.W}

	mid := q.Slerp(antipodal, 0.5)

	v := Vec3{1, 0, 0}
	got := mid.RotateVector(v)
	want := q.RotateVector(v)

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("Slerp shortest path: got %#v, want %#v", got, want)
	}
}

func TestQuaternionSlerpNearParallel(t *testing.T) {
	q := QuaternionFromAxisAngle(Vec3{0, 1, 0}, 0.5)
	other := QuaternionFromAxisAngle(Vec3{0, 1, 0}, 0.5+1e-9)

	// Should not panic or produce NaN when the two orientations are almost
	// identical (the near-zero-sine branch).
	mid := q.Slerp(other, 0.5)
	if math.IsNaN(mid.X) || math.IsNaN(mid.Y) || math.IsNaN(mid.Z) || math.IsNaN(mid.W) {
		t.Fatalf("Slerp of near-parallel quaternions produced NaN: %#v", mid)
	}
}

func ExampleQuaternionFromAxisAngle() {
	q := QuaternionFromAxisAngle(Vec3{X: 0, Y: 0, Z: 1}, math.Pi/2)
	v := q.RotateVector(Vec3{X: 1, Y: 0, Z: 0})

	fmt.Printf("%.0f %.0f %.0f\n", v.X, v.Y, v.Z)

	// Output:
	// 0 1 0
}
