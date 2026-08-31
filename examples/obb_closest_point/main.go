package main

import (
	"fmt"
	"math"

	"github.com/motah-fard/geom3d"
)

func main() {
	box := geom3d.OBB{
		Center:      geom3d.Vec3{X: 0, Y: 0, Z: 0},
		HalfExtents: geom3d.Vec3{X: 1, Y: 1, Z: 1},
		Orientation: geom3d.RotationZ(math.Pi / 4),
	}

	p := geom3d.Vec3{X: 3, Y: 0, Z: 0}

	fmt.Println("contains origin:", box.Contains(geom3d.Vec3{X: 0, Y: 0, Z: 0}))
	fmt.Println("closest point on rotated box:", geom3d.ClosestPointOnOBB(p, box))
	fmt.Println("distance to rotated box:", geom3d.DistancePointToOBB(p, box))
	fmt.Println("volume:", box.Volume())
}
