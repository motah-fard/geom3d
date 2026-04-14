package geom3d

import (
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
