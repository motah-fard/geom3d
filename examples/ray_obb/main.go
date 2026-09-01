package main

import (
	"fmt"
	"math"

	"github.com/motah-fard/geom3d"
)

func main() {
	box := geom3d.OBB{
		Center:      geom3d.Vec3{X: 0, Y: 0, Z: 0},
		HalfExtents: geom3d.Vec3{X: 2, Y: 1, Z: 1},
		Orientation: geom3d.RotationZ(math.Pi / 2),
	}

	r := geom3d.Ray3{
		Origin: geom3d.Vec3{X: -5, Y: 0, Z: 0},
		Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
	}

	hit, tMin, tMax := geom3d.IntersectRayOBB(r, box)

	fmt.Println("hit:", hit)
	fmt.Println("tMin:", tMin)
	fmt.Println("tMax:", tMax)
}
