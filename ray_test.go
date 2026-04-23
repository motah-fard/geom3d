package geom3d

import (
	"fmt"
	"testing"
)

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
func ExampleRay3_PointAt() {
	r := Ray3{
		Origin: Vec3{X: 1, Y: 2, Z: 3},
		Dir:    Vec3{X: 2, Y: 0, Z: -1},
	}

	fmt.Println(r.PointAt(2))

	// Output:
	// {5 2 1}
}
func ExampleRay3_IsValid() {
	fmt.Println(Ray3{Dir: Vec3{X: 1, Y: 0, Z: 0}}.IsValid())
	fmt.Println(Ray3{}.IsValid())

	// Output:
	// true
	// false
}
