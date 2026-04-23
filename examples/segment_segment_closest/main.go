package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	s1 := geom3d.Segment3{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 2, Y: 0, Z: 0},
	}
	s2 := geom3d.Segment3{
		A: geom3d.Vec3{X: 0, Y: 1, Z: 0},
		B: geom3d.Vec3{X: 2, Y: 1, Z: 0},
	}

	fmt.Println("distance:", geom3d.DistanceBetweenSegments(s1, s2))
}
