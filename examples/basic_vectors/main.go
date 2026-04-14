package main

import (
	"fmt"

	"github.com/motah-fard/geom3d"
)

func main() {
	a := geom3d.Vec3{X: 1, Y: 2, Z: 3}
	b := geom3d.Vec3{X: 4, Y: 5, Z: 6}

	fmt.Println("a + b =", a.Add(b))
	fmt.Println("a - b =", a.Sub(b))
	fmt.Println("a · b =", a.Dot(b))
	fmt.Println("a × b =", a.Cross(b))
	fmt.Println("distance =", a.Distance(b))
}
