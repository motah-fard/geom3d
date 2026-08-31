package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	a := geom3d.Sphere{Center: geom3d.Vec3{X: 0, Y: 0, Z: 0}, Radius: 2}
	b := geom3d.Sphere{Center: geom3d.Vec3{X: 3, Y: 0, Z: 0}, Radius: 2}
	c := geom3d.Sphere{Center: geom3d.Vec3{X: 20, Y: 0, Z: 0}, Radius: 2}

	fmt.Println("a overlaps b:", a.Overlaps(b))
	fmt.Println("a overlaps c:", a.Overlaps(c))
	fmt.Println("distance between a and c:", geom3d.DistanceBetweenSpheres(a, c))

	box := geom3d.AABB{Min: geom3d.Vec3{X: -1, Y: -1, Z: -1}, Max: geom3d.Vec3{X: 1, Y: 1, Z: 1}}
	fmt.Println("box overlaps sphere a:", geom3d.IntersectAABBSphere(box, a))
}
