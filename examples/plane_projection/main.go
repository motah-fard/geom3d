package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	pl := geom3d.Plane{
		Point:  geom3d.Vec3{X: 0, Y: 0, Z: 0},
		Normal: geom3d.Vec3{X: 0, Y: 0, Z: 1},
	}

	p1 := geom3d.Vec3{X: 2, Y: 3, Z: 5}
	p2 := geom3d.Vec3{X: -1, Y: 4, Z: -2}

	proj1 := geom3d.ProjectPointToPlane(p1, pl)
	proj2 := geom3d.ProjectPointToPlane(p2, pl)

	fmt.Println("plane point:", pl.Point)
	fmt.Println("plane normal:", pl.Normal)

	fmt.Println("project", p1, "->", proj1)
	fmt.Println("project", p2, "->", proj2)
}
