package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestIdentityMat3MulVec(t *testing.T) {
	m := IdentityMat3()
	v := Vec3{1, 2, 3}

	got := m.MulVec(v)
	want := v

	if got != want {
		t.Fatalf("IdentityMat3 MulVec: got %#v, want %#v", got, want)
	}
}

func TestRotationZ90(t *testing.T) {
	m := RotationZ(math.Pi / 2)
	v := Vec3{1, 0, 0}

	got := m.MulVec(v)
	want := Vec3{0, 1, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("RotationZ 90: got %#v, want %#v", got, want)
	}
}

func TestRotationX90(t *testing.T) {
	m := RotationX(math.Pi / 2)
	v := Vec3{0, 1, 0}

	got := m.MulVec(v)
	want := Vec3{0, 0, 1}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("RotationX 90: got %#v, want %#v", got, want)
	}
}

func TestMat3Mul(t *testing.T) {
	a := IdentityMat3()
	b := RotationY(math.Pi / 4)

	got := a.Mul(b)
	want := b

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !AlmostEqual(got.M[i][j], want.M[i][j]) {
				t.Fatalf("Mat3 Mul mismatch at (%d,%d): got %v, want %v", i, j, got.M[i][j], want.M[i][j])
			}
		}
	}
}

func TestMat3Transpose(t *testing.T) {
	m := Mat3{
		M: [3][3]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	got := m.Transpose()
	want := Mat3{
		M: [3][3]float64{
			{1, 4, 7},
			{2, 5, 8},
			{3, 6, 9},
		},
	}

	if got != want {
		t.Fatalf("Transpose: got %#v, want %#v", got, want)
	}
}
func TestMat3TransposeTwice(t *testing.T) {
	m := Mat3{
		M: [3][3]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}

	got := m.Transpose().Transpose()
	if got != m {
		t.Fatalf("Transpose twice: got %#v, want %#v", got, m)
	}
}
func TestRotationZPreservesNorm(t *testing.T) {
	m := RotationZ(math.Pi / 3)
	v := Vec3{3, 4, 0}

	got := m.MulVec(v)

	if !AlmostEqual(got.Norm(), v.Norm()) {
		t.Fatalf("rotation should preserve norm: got %v, want %v", got.Norm(), v.Norm())
	}
}
func TestMat3Determinant(t *testing.T) {
	if got, want := IdentityMat3().Determinant(), 1.0; !AlmostEqual(got, want) {
		t.Fatalf("Determinant of identity: got %v, want %v", got, want)
	}

	m := Mat3{M: [3][3]float64{
		{2, 0, 0},
		{0, 3, 0},
		{0, 0, 4},
	}}
	if got, want := m.Determinant(), 24.0; !AlmostEqual(got, want) {
		t.Fatalf("Determinant of diagonal matrix: got %v, want %v", got, want)
	}

	singular := Mat3{M: [3][3]float64{
		{1, 2, 3},
		{2, 4, 6},
		{1, 1, 1},
	}}
	if got, want := singular.Determinant(), 0.0; !AlmostEqual(got, want) {
		t.Fatalf("Determinant of singular matrix: got %v, want %v", got, want)
	}
}

func TestMat3InverseIdentity(t *testing.T) {
	inv, ok := IdentityMat3().Inverse()
	if !ok {
		t.Fatal("expected identity matrix to be invertible")
	}
	if inv != IdentityMat3() {
		t.Fatalf("Inverse of identity: got %#v, want identity", inv)
	}
}

func TestMat3InverseRoundTrip(t *testing.T) {
	m := RotationX(0.4).Mul(RotationY(0.9)).Mul(RotationZ(1.3))

	inv, ok := m.Inverse()
	if !ok {
		t.Fatal("expected rotation matrix to be invertible")
	}

	got := m.Mul(inv)
	want := IdentityMat3()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !AlmostEqual(got.M[i][j], want.M[i][j]) {
				t.Fatalf("m * m.Inverse() mismatch at (%d,%d): got %v, want %v", i, j, got.M[i][j], want.M[i][j])
			}
		}
	}
}

func TestMat3InverseSingular(t *testing.T) {
	singular := Mat3{M: [3][3]float64{
		{1, 2, 3},
		{2, 4, 6},
		{1, 1, 1},
	}}

	_, ok := singular.Inverse()
	if ok {
		t.Fatal("expected singular matrix to have no inverse")
	}
}

func TestMat3ToQuaternionRoundTrip(t *testing.T) {
	m := RotationX(0.3).Mul(RotationY(0.7)).Mul(RotationZ(1.1))

	got := m.ToQuaternion().ToMat3()

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !AlmostEqual(got.M[i][j], m.M[i][j]) {
				t.Fatalf("Mat3->Quaternion->Mat3 mismatch at (%d,%d): got %v, want %v", i, j, got.M[i][j], m.M[i][j])
			}
		}
	}
}

func TestMat3ToQuaternionMatchesAxisAngle(t *testing.T) {
	got := RotationZ(math.Pi / 2).ToQuaternion()
	want := QuaternionFromAxisAngle(Vec3{0, 0, 1}, math.Pi/2)

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) ||
		!AlmostEqual(got.Z, want.Z) || !AlmostEqual(got.W, want.W) {
		t.Fatalf("ToQuaternion: got %#v, want %#v", got, want)
	}
}

func TestMat3ToQuaternionLargestDiagonalBranches(t *testing.T) {
	// Each of these is a 180-degree rotation with trace <= 0, forcing
	// ToQuaternion into its "largest diagonal element" branches instead of
	// the trace > 0 fast path: X for RotationX, Y for RotationY, and the
	// default (Z) branch for RotationZ.
	rotations := []Mat3{
		RotationX(math.Pi),
		RotationY(math.Pi),
		RotationZ(math.Pi),
	}

	for i, m := range rotations {
		got := m.ToQuaternion().ToMat3()
		for r := 0; r < 3; r++ {
			for c := 0; c < 3; c++ {
				if !AlmostEqual(got.M[r][c], m.M[r][c]) {
					t.Fatalf("rotation %d: ToQuaternion round trip mismatch at (%d,%d): got %v, want %v", i, r, c, got.M[r][c], m.M[r][c])
				}
			}
		}
	}
}

func ExampleMat3_Inverse() {
	m := RotationZ(math.Pi / 2)
	inv, ok := m.Inverse()

	v := Vec3{X: 0, Y: 1, Z: 0}
	back := inv.MulVec(m.MulVec(v))

	fmt.Println(ok)
	fmt.Printf("%.0f %.0f %.0f\n", back.X, back.Y, back.Z)

	// Output:
	// true
	// 0 1 0
}

func ExampleRotationZ() {
	m := RotationZ(math.Pi / 2)
	v := Vec3{X: 1, Y: 0, Z: 0}
	out := m.MulVec(v)

	fmt.Printf("%.0f %.0f %.0f\n", out.X, out.Y, out.Z)

	// Output:
	// 0 1 0
}
func ExampleIdentityMat3() {
	m := IdentityMat3()
	v := Vec3{X: 1, Y: 2, Z: 3}
	out := m.MulVec(v)

	fmt.Printf("%.0f %.0f %.0f\n", out.X, out.Y, out.Z)

	// Output:
	// 1 2 3
}
