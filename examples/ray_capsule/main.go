package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	c := geom3d.Capsule{
		A:      geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B:      geom3d.Vec3{X: 0, Y: 0, Z: 4},
		Radius: 1,
	}

	throughBody := geom3d.Ray3{Origin: geom3d.Vec3{X: -5, Y: 0, Z: 2}, Dir: geom3d.Vec3{X: 1, Y: 0, Z: 0}}
	throughCaps := geom3d.Ray3{Origin: geom3d.Vec3{X: 0, Y: 0, Z: 10}, Dir: geom3d.Vec3{X: 0, Y: 0, Z: -1}}

	hit, tMin, tMax := geom3d.IntersectRayCapsule(throughBody, c)
	fmt.Println("through the cylindrical body:", hit, tMin, tMax)

	hit, tMin, tMax = geom3d.IntersectRayCapsule(throughCaps, c)
	fmt.Println("straight down through both end caps:", hit, tMin, tMax)
}
