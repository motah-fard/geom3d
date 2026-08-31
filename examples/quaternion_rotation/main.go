package main

import (
	"fmt"
	"math"

	"github.com/motah-fard/geom3d"
)

func main() {
	// A 90-degree rotation about the Z axis.
	q := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 0, Y: 0, Z: 1}, math.Pi/2)

	v := geom3d.Vec3{X: 1, Y: 0, Z: 0}
	fmt.Println("rotated vector:", q.RotateVector(v))

	// Quaternion composition: apply q2 first, then q1 (same order as
	// Mat3.Mul and Transform.Compose).
	q2 := geom3d.QuaternionFromAxisAngle(geom3d.Vec3{X: 1, Y: 0, Z: 0}, math.Pi/2)
	composed := q.Mul(q2)
	fmt.Println("composed rotation applied to vector:", composed.RotateVector(v))

	// Interpolating smoothly between two orientations.
	mid := geom3d.IdentityQuaternion().Slerp(q, 0.5)
	fmt.Println("halfway rotation applied to vector:", mid.RotateVector(v))

	// Converting to a Mat3 to apply the same rotation to many points cheaply.
	m := q.ToMat3()
	fmt.Println("same rotation via Mat3:", m.MulVec(v))
}
