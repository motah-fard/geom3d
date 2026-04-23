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

	ray1 := geom3d.Ray3{
		Origin: geom3d.Vec3{X: -1, Y: 1, Z: 1},
		Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
	}

	ray2 := geom3d.Ray3{
		Origin: geom3d.Vec3{X: -1, Y: 3, Z: 1},
		Dir:    geom3d.Vec3{X: 1, Y: 0, Z: 0},
	}

	hit1, tMin1, tMax1 := geom3d.IntersectRayAABB(ray1, box)
	hit2, _, _ := geom3d.IntersectRayAABB(ray2, box)

	fmt.Println("box center:", box.Center())
	fmt.Println("box size:", box.Size())

	fmt.Printf("ray1 hit=%v tMin=%.0f tMax=%.0f\n", hit1, tMin1, tMax1)
	fmt.Printf("ray2 hit=%v\n", hit2)
}
