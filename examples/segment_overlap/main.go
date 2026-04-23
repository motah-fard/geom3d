package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	s1 := geom3d.Segment3{
		A: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 4, Y: 0, Z: 0},
	}
	s2 := geom3d.Segment3{
		A: geom3d.Vec3{X: 2, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 6, Y: 0, Z: 0},
	}
	s3 := geom3d.Segment3{
		A: geom3d.Vec3{X: 4, Y: 0, Z: 0},
		B: geom3d.Vec3{X: 5, Y: 0, Z: 0},
	}

	fmt.Println("s1 overlaps s2:", geom3d.SegmentsOverlap(s1, s2))
	fmt.Println("s1 overlaps s3:", geom3d.SegmentsOverlap(s1, s3))
}
