package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	s := geom3d.Sphere{
		Center: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Radius: 2,
	}

	p1 := geom3d.Vec3{X: 4, Y: 0, Z: 0}
	p2 := geom3d.Vec3{X: 1, Y: 0, Z: 0}

	fmt.Println("closest point to", p1, "is", geom3d.ClosestPointOnSphere(p1, s))
	fmt.Println("closest point to", p2, "is", geom3d.ClosestPointOnSphere(p2, s))
	fmt.Println("distance from", p1, "is", geom3d.DistancePointToSphere(p1, s))
}
