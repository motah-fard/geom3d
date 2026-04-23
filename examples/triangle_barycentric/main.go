package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	tri := geom3d.Triangle{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 1, Y: 0, Z: 0},
		C: geom3d.Vec3{X: 0, Y: 1, Z: 0},
	}

	p1 := geom3d.Vec3{X: 0.25, Y: 0.25, Z: 0}
	p2 := geom3d.Vec3{X: 1.5, Y: 0.25, Z: 0}

	u1, v1, w1 := geom3d.BarycentricCoordinates(p1, tri)
	u2, v2, w2 := geom3d.BarycentricCoordinates(p2, tri)

	fmt.Println("triangle:", tri)

	fmt.Printf("barycentric of %v: u=%.2f v=%.2f w=%.2f\n", p1, u1, v1, w1)
	fmt.Printf("barycentric of %v: u=%.2f v=%.2f w=%.2f\n", p2, u2, v2, w2)
}
