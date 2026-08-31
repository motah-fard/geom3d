package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	box := geom3d.AABB{
		Min: geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Max: geom3d.Vec3{X: 1, Y: 1, Z: 1},
	}

	// A Ray3 is unbounded, so an infinite ray "hits" a box even if the box
	// is far beyond where the caller actually cares about. A Segment3 lets
	// the caller ask "does anything happen within this specific distance."
	near := geom3d.Segment3{A: geom3d.Vec3{X: -1, Y: 0.5, Z: 0.5}, B: geom3d.Vec3{X: 2, Y: 0.5, Z: 0.5}}
	short := geom3d.Segment3{A: geom3d.Vec3{X: -1, Y: 0.5, Z: 0.5}, B: geom3d.Vec3{X: -0.5, Y: 0.5, Z: 0.5}}

	hit, tMin, tMax := geom3d.IntersectSegmentAABB(near, box)
	fmt.Println("segment long enough to reach the box:", hit, tMin, tMax)

	hit, _, _ = geom3d.IntersectSegmentAABB(short, box)
	fmt.Println("segment too short to reach the box:", hit)
}
