package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	r := geom3d.Ray3{
		Origin: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Dir:    geom3d.Vec3{X: 0, Y: 0, Z: 1},
	}

	pl := geom3d.Plane{
		Point:  geom3d.Vec3{X: 0, Y: 0, Z: 5},
		Normal: geom3d.Vec3{X: 0, Y: 0, Z: 1},
	}

	p, ok := geom3d.IntersectRayPlane(r, pl)
	fmt.Println("hit:", ok)
	fmt.Println("point:", p)
}
