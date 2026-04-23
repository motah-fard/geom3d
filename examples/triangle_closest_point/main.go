package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	tri := geom3d.Triangle{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 2, Y: 0, Z: 0},
		C: geom3d.Vec3{X: 0, Y: 2, Z: 0},
	}

	p1 := geom3d.Vec3{X: 0.5, Y: 0.5, Z: 3}
	p2 := geom3d.Vec3{X: 1.5, Y: -1, Z: 0}
	p3 := geom3d.Vec3{X: -1, Y: -1, Z: 0}

	cp1 := geom3d.ClosestPointOnTriangle(p1, tri)
	cp2 := geom3d.ClosestPointOnTriangle(p2, tri)
	cp3 := geom3d.ClosestPointOnTriangle(p3, tri)

	fmt.Println("triangle:", tri)

	fmt.Println("closest point to", p1, "is", cp1)
	fmt.Println("closest point to", p2, "is", cp2)
	fmt.Println("closest point to", p3, "is", cp3)
}
