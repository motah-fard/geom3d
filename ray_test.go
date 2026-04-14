package geom3d

import "testing"

func TestRay3PointAt(t *testing.T) {
	r := Ray3{
		Origin: Vec3{1, 2, 3},
		Dir:    Vec3{2, 0, -1},
	}

	got := r.PointAt(2)
	want := Vec3{5, 2, 1}

	if got != want {
		t.Fatalf("PointAt: got %#v, want %#v", got, want)
	}
}

func TestRay3IsValid(t *testing.T) {
	r := Ray3{Dir: Vec3{1, 0, 0}}
	if !r.IsValid() {
		t.Fatal("expected valid ray")
	}

	zero := Ray3{}
	if zero.IsValid() {
		t.Fatal("expected zero-direction ray to be invalid")
	}
}
