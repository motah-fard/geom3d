package geom3d

import (
	"fmt"
	"testing"
)

func TestProjectPointToPlane(t *testing.T) {
	pl := Plane{
		Point:  Vec3{0, 0, 0},
		Normal: Vec3{0, 0, 1},
	}

	p := Vec3{2, 3, 5}
	got := ProjectPointToPlane(p, pl)
	want := Vec3{2, 3, 0}

	if got != want {
		t.Fatalf("ProjectPointToPlane: got %#v, want %#v", got, want)
	}
}

func TestProjectPointToPlaneInvalidPlane(t *testing.T) {
	got := ProjectPointToPlane(Vec3{1, 2, 3}, Plane{})
	want := Vec3{}

	if got != want {
		t.Fatalf("ProjectPointToPlane invalid plane: got %#v, want %#v", got, want)
	}
}

func TestProjectPointToLine(t *testing.T) {
	a := Vec3{0, 0, 0}
	b := Vec3{2, 0, 0}
	p := Vec3{1, 3, 0}

	got := ProjectPointToLine(p, a, b)
	want := Vec3{1, 0, 0}

	if got != want {
		t.Fatalf("ProjectPointToLine: got %#v, want %#v", got, want)
	}
}

func TestProjectPointToLineDegenerate(t *testing.T) {
	a := Vec3{4, 5, 6}
	p := Vec3{1, 2, 3}

	got := ProjectPointToLine(p, a, a)
	want := a

	if got != want {
		t.Fatalf("ProjectPointToLine degenerate: got %#v, want %#v", got, want)
	}
}

func ExampleProjectPointToPlane() {
	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 0},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	p := Vec3{X: 2, Y: 3, Z: 5}
	fmt.Println(ProjectPointToPlane(p, pl))

	// Output:
	// {2 3 0}
}
func ExampleProjectPointToPlane_invalidPlane() {
	pl := Plane{}
	p := Vec3{X: 2, Y: 3, Z: 5}

	fmt.Println(ProjectPointToPlane(p, pl))

	// Output:
	// {0 0 0}
}
func ExampleProjectPointToLine() {
	a := Vec3{X: 0, Y: 0, Z: 0}
	b := Vec3{X: 2, Y: 0, Z: 0}
	p := Vec3{X: 1, Y: 3, Z: 0}

	fmt.Println(ProjectPointToLine(p, a, b))

	// Output:
	// {1 0 0}
}
