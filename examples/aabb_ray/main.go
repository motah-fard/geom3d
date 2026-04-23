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

	fmt.Println("box center:", box.Center())
	fmt.Println("box size:", box.Size())

	fmt.Println("ray1 intersects box:", geom3d.IntersectRayAABB(ray1, box))
	fmt.Println("ray2 intersects box:", geom3d.IntersectRayAABB(ray2, box))
}
