package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	points := []geom3d.Vec3{
		{X: 1, Y: 5, Z: 0},
		{X: -2, Y: 1, Z: 3},
		{X: 4, Y: -1, Z: -1},
	}

	box := geom3d.AABBFromPoints(points)

	fmt.Println("points:", points)
	fmt.Println("bounding box min:", box.Min)
	fmt.Println("bounding box max:", box.Max)
	fmt.Println("volume:", box.Volume())

	margin := geom3d.AABB{Min: geom3d.Vec3{X: 10, Y: 10, Z: 10}, Max: geom3d.Vec3{X: 12, Y: 12, Z: 12}}
	fmt.Println("union with a distant box:", box.Union(margin))
}
