package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	r := geom3d.Ray3{
		Origin: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
	}

	p1 := geom3d.Vec3{X: 2, Y: 3, Z: 0}
	p2 := geom3d.Vec3{X: -2, Y: 1, Z: 0}

	fmt.Println("closest point to", p1, "is", geom3d.ClosestPointOnRay(p1, r))
	fmt.Println("closest point to", p2, "is", geom3d.ClosestPointOnRay(p2, r))
}
