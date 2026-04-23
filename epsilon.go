package geom3d

import "math"

// Epsilon is the default tolerance used for floating-point comparisons.
const Epsilon = 1e-9

// AlmostZero reports whether x is within Epsilon of zero.
func AlmostZero(x float64) bool {
	return math.Abs(x) <= Epsilon
}

// AlmostEqual reports whether a and b differ by no more than Epsilon.
func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= Epsilon
}
