package geom3d

import (
	"fmt"
	"testing"
)

func TestPlaneIsValid(t *testing.T) {
	p := Plane{Normal: Vec3{0, 0, 1}}
	if !p.IsValid() {
		t.Fatal("expected valid plane")
	}

	zero := Plane{}
	if zero.IsValid() {
		t.Fatal("expected zero-normal plane to be invalid")
	}
}

func TestPlaneUnitNormal(t *testing.T) {
	p := Plane{Normal: Vec3{0, 0, 5}}
	got := p.UnitNormal()
	want := Vec3{0, 0, 1}

	if got != want {
		t.Fatalf("UnitNormal: got %#v, want %#v", got, want)
	}
}

func ExamplePlane_UnitNormal() {
	p := Plane{
		Normal: Vec3{X: 0, Y: 0, Z: 5},
	}

	fmt.Println(p.UnitNormal())

	// Output:
	// {0 0 1}
}
