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

	r := geom3d.Ray3{
		Origin: geom3d.Vec3{X: -5, Y: 0, Z: 0},
		Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
	}

	hit, tMin, tMax := geom3d.IntersectRaySphere(r, s)

	fmt.Println("sphere:", s)
	fmt.Println("ray origin:", r.Origin)
	fmt.Println("ray direction:", r.Dir)
	fmt.Println("hit:", hit)
	fmt.Println("tMin:", tMin)
	fmt.Println("tMax:", tMax)
}
