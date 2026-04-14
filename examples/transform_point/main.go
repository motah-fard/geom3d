package main

import (
	"fmt"
	"math"

	"github.com/motah-fard/geom3d"
)

func main() {
	tf := geom3d.Transform{
		R: geom3d.RotationZ(math.Pi / 2),
		T: geom3d.Vec3{X: 10, Y: 0, Z: 0},
	}

	p := geom3d.Vec3{X: 1, Y: 0, Z: 0}
	q := tf.ApplyPoint(p)

	fmt.Println("original:", p)
	fmt.Println("transformed:", q)
}
