package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	seg := geom3d.Segment3{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 4, Y: 0, Z: 0},
	}

	p1 := geom3d.Vec3{X: 2, Y: 3, Z: 0}
	p2 := geom3d.Vec3{X: -1, Y: 1, Z: 0}
	p3 := geom3d.Vec3{X: 6, Y: -2, Z: 0}

	c1 := geom3d.ClosestPointOnSegment(p1, seg)
	c2 := geom3d.ClosestPointOnSegment(p2, seg)
	c3 := geom3d.ClosestPointOnSegment(p3, seg)

	fmt.Println("segment:", seg)
	fmt.Println("length:", seg.Length())

	fmt.Println("closest point to", p1, "is", c1)
	fmt.Println("closest point to", p2, "is", c2)
	fmt.Println("closest point to", p3, "is", c3)
}
