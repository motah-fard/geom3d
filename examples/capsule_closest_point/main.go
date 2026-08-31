package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	// A capsule standing along the Z axis, radius 1.
	c := geom3d.Capsule{
		A:      geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B:      geom3d.Vec3{X: 0, Y: 0, Z: 4},
		Radius: 1,
	}

	beside := geom3d.Vec3{X: 3, Y: 0, Z: 2}   // out from the cylindrical body
	above := geom3d.Vec3{X: 0, Y: 0, Z: 6}    // beyond the top end cap
	inside := geom3d.Vec3{X: 0.5, Y: 0, Z: 2} // inside the capsule

	fmt.Println("closest point beside the capsule:", geom3d.ClosestPointOnCapsule(beside, c))
	fmt.Println("closest point above the capsule:", geom3d.ClosestPointOnCapsule(above, c))
	fmt.Println("closest point for an interior point:", geom3d.ClosestPointOnCapsule(inside, c))
	fmt.Println("distance from beside:", geom3d.DistancePointToCapsule(beside, c))
	fmt.Println("volume:", c.Volume())
}
