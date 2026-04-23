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

	n := tri.Normal()

	fmt.Println("triangle:", tri)
	fmt.Println("normal:", n)
	fmt.Println("area:", tri.Area())
	fmt.Println("degenerate:", tri.IsDegenerate())
}
