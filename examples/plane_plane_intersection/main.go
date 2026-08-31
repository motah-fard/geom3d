package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	p1 := geom3d.Plane{Point: geom3d.Vec3{X: 0, Y: 0, Z: 5}, Normal: geom3d.Vec3{X: 0, Y: 0, Z: 1}} // z = 5
	p2 := geom3d.Plane{Point: geom3d.Vec3{X: 3, Y: 0, Z: 0}, Normal: geom3d.Vec3{X: 1, Y: 0, Z: 0}} // x = 3

	line, ok := geom3d.IntersectPlanePlane(p1, p2)
	fmt.Println("planes intersect:", ok)
	fmt.Println("a point on the intersection line:", line.Point)
	fmt.Println("the intersection line's direction:", line.Dir)

	// A Line3 extends in both directions, unlike Ray3.
	pl := geom3d.Plane{Point: geom3d.Vec3{X: 0, Y: 0, Z: -10}, Normal: geom3d.Vec3{X: 0, Y: 0, Z: 1}}
	l := geom3d.Line3{Point: geom3d.Vec3{X: 0, Y: 0, Z: 0}, Dir: geom3d.Vec3{X: 0, Y: 0, Z: 1}}
	p, hit := geom3d.IntersectLinePlane(l, pl)
	fmt.Println("line hits plane behind its own point:", hit, p)
}
