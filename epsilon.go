package geom3d

import "math"

const Epsilon = 1e-9

func AlmostZero(x float64) bool {
	return math.Abs(x) <= Epsilon
}

func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= Epsilon
}
