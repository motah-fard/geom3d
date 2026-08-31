package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	tri := geom3d.Triangle{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 4, Y: 0, Z: 0},
		C: geom3d.Vec3{X: 0, Y: 4, Z: 0},
	}

	r := geom3d.Ray3{
		Origin: geom3d.Vec3{X: 1, Y: 1, Z: 5},
		Dir:    geom3d.Vec3{X: 0, Y: 0, Z: -1},
	}

	p, ok := geom3d.IntersectRayTriangle(r, tri)

	fmt.Println("triangle:", tri)
	fmt.Println("ray origin:", r.Origin)
	fmt.Println("ray direction:", r.Dir)
	fmt.Println("hit:", ok)
	fmt.Println("intersection point:", p)
}
