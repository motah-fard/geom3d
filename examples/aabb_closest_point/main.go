package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	box := geom3d.AABB{
		Min: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Max: geom3d.Vec3{X: 2, Y: 2, Z: 2},
	}

	p1 := geom3d.Vec3{X: 3, Y: -1, Z: 1}
	p2 := geom3d.Vec3{X: 1, Y: 1, Z: 1}

	fmt.Println("closest point to", p1, "is", geom3d.ClosestPointOnAABB(p1, box))
	fmt.Println("closest point to", p2, "is", geom3d.ClosestPointOnAABB(p2, box))
}
