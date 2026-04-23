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

	fmt.Println("distance from", p1, "to ray:", geom3d.DistancePointToRay(p1, r))
	fmt.Println("distance from", p2, "to ray:", geom3d.DistancePointToRay(p2, r))
}
