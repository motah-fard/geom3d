package geom3d

import (
	"fmt"
	"testing"
)

func TestDistancePointToPlane(t *testing.T) {
	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 0},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	tests := []struct {
		name string
		p    Vec3
		want float64
	}{
		{
			name: "point above plane",
			p:    Vec3{X: 1, Y: 2, Z: 5},
			want: 5,
		},
		{
			name: "point below plane",
			p:    Vec3{X: -1, Y: 0, Z: -3},
			want: -3,
		},
		{
			name: "point on plane",
			p:    Vec3{X: 4, Y: 5, Z: 0},
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DistancePointToPlane(tc.p, pl)
			if !AlmostEqual(got, tc.want) {
				t.Fatalf("DistancePointToPlane(%v, pl) = %v, want %v", tc.p, got, tc.want)
			}
		})
	}
}

func TestDistancePointToPlaneInvalidPlane(t *testing.T) {
	p := Vec3{X: 1, Y: 2, Z: 3}
	got := DistancePointToPlane(p, Plane{})

	if got != 0 {
		t.Fatalf("DistancePointToPlane with invalid plane = %v, want 0", got)
	}
}

func ExampleDistancePointToPlane() {
	pl := Plane{
		Point:  Vec3{X: 0, Y: 0, Z: 0},
		Normal: Vec3{X: 0, Y: 0, Z: 1},
	}

	p := Vec3{X: 2, Y: 3, Z: 5}
	fmt.Println(DistancePointToPlane(p, pl))

	// Output:
	// 5
}
