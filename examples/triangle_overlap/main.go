package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	t1 := geom3d.Triangle{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 4, Y: 0, Z: 0},
		C: geom3d.Vec3{X: 0, Y: 4, Z: 0},
	}

	// Coplanar, partially overlapping.
	t2 := geom3d.Triangle{
		A: geom3d.Vec3{X: 1, Y: 1, Z: 0},
		B: geom3d.Vec3{X: 5, Y: 1, Z: 0},
		C: geom3d.Vec3{X: 1, Y: 5, Z: 0},
	}

	// Non-coplanar, piercing straight through t1.
	t3 := geom3d.Triangle{
		A: geom3d.Vec3{X: 1, Y: 1, Z: -5},
		B: geom3d.Vec3{X: 1, Y: 1, Z: 5},
		C: geom3d.Vec3{X: 3, Y: 1, Z: 5},
	}

	// Coplanar but disjoint.
	t4 := geom3d.Triangle{
		A: geom3d.Vec3{X: 10, Y: 10, Z: 0},
		B: geom3d.Vec3{X: 11, Y: 10, Z: 0},
		C: geom3d.Vec3{X: 10, Y: 11, Z: 0},
	}

	fmt.Println("coplanar partial overlap:", t1.Overlaps(t2))
	fmt.Println("non-coplanar piercing overlap:", t1.Overlaps(t3))
	fmt.Println("coplanar disjoint:", t1.Overlaps(t4))
}
