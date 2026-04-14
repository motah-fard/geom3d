package geom3d

import (
	"math"
	"testing"
)

func TestIdentityTransformApplyPoint(t *testing.T) {
	tf := IdentityTransform()
	p := Vec3{1, 2, 3}

	got := tf.ApplyPoint(p)
	want := p

	if got != want {
		t.Fatalf("IdentityTransform ApplyPoint: got %#v, want %#v", got, want)
	}
}

func TestTransformApplyPoint(t *testing.T) {
	tf := Transform{
		R: RotationZ(math.Pi / 2),
		T: Vec3{10, 0, 0},
	}
	p := Vec3{1, 0, 0}

	got := tf.ApplyPoint(p)
	want := Vec3{10, 1, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ApplyPoint: got %#v, want %#v", got, want)
	}
}

func TestTransformApplyVector(t *testing.T) {
	tf := Transform{
		R: RotationZ(math.Pi / 2),
		T: Vec3{10, 20, 30},
	}
	v := Vec3{1, 0, 0}

	got := tf.ApplyVector(v)
	want := Vec3{0, 1, 0}

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("ApplyVector: got %#v, want %#v", got, want)
	}
}

func TestTransformCompose(t *testing.T) {
	a := Transform{
		R: RotationZ(math.Pi / 2),
		T: Vec3{1, 0, 0},
	}
	b := Transform{
		R: IdentityMat3(),
		T: Vec3{0, 2, 0},
	}
	p := Vec3{1, 0, 0}

	got1 := a.Compose(b).ApplyPoint(p)
	got2 := a.ApplyPoint(b.ApplyPoint(p))

	if !AlmostEqual(got1.X, got2.X) || !AlmostEqual(got1.Y, got2.Y) || !AlmostEqual(got1.Z, got2.Z) {
		t.Fatalf("Compose mismatch: got %#v, want %#v", got1, got2)
	}
}

func TestTransformInverse(t *testing.T) {
	tf := Transform{
		R: RotationZ(math.Pi / 2),
		T: Vec3{3, 4, 0},
	}
	p := Vec3{1, 2, 0}

	q := tf.ApplyPoint(p)
	got := tf.Inverse().ApplyPoint(q)
	want := p

	if !AlmostEqual(got.X, want.X) || !AlmostEqual(got.Y, want.Y) || !AlmostEqual(got.Z, want.Z) {
		t.Fatalf("Inverse: got %#v, want %#v", got, want)
	}
}
