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

// RelativeEpsilon is the default tolerance used by AlmostEqualRelative,
// expressed as a fraction of the larger operand's magnitude.
const RelativeEpsilon = 1e-9

// AlmostEqualRelative reports whether a and b are close to each other in
// relative terms: their difference is at most RelativeEpsilon times the
// larger of |a| and |b|, falling back to the fixed Epsilon near zero
// (where a relative tolerance is meaningless, since anything is "close"
// to zero in relative terms).
//
// AlmostEqual and AlmostZero — used throughout the rest of this package,
// including by every IsValid/IsDegenerate check — compare against a fixed
// absolute Epsilon (1e-9). That works well at "ordinary" coordinate
// magnitudes, but breaks down at large ones: two values that agree to 12
// significant figures can still differ by far more than 1e-9 in absolute
// terms once they're in the millions, even though they're "the same" for
// any practical purpose. AlmostEqualRelative is provided for exactly that
// situation — comparing geom3d's own outputs, or your own values, at
// scales where the fixed Epsilon no longer applies.
//
// This function does not change the behavior of any other function in the
// package: geom3d's internal validity/degeneracy checks continue to use
// the fixed Epsilon, and that won't change within v1 (see the README's
// "Error handling" section). Use AlmostEqualRelative in your own code when
// you need scale-aware comparisons; it does not retroactively make
// geom3d's own geometry queries scale-aware.
func AlmostEqualRelative(a, b float64) bool {
	if a == b {
		return true // also handles +Inf == +Inf and -Inf == -Inf
	}

	diff := math.Abs(a - b)
	largest := math.Max(math.Abs(a), math.Abs(b))

	return diff <= largest*RelativeEpsilon || diff <= Epsilon
}

// AlmostZeroAtScale reports whether x is negligible relative to scale,
// falling back to the fixed Epsilon when scale is small. Unlike
// AlmostEqualRelative, "is this near zero" has no second operand to scale
// against, so the caller must supply what magnitude "zero" should be
// judged relative to — for example, the norm of the vectors involved in
// computing x, if x came from a cross product or similar.
//
// As with AlmostEqualRelative, this exists for comparisons in your own
// code at large coordinate magnitudes; it does not change how any other
// function in this package behaves.
func AlmostZeroAtScale(x, scale float64) bool {
	return math.Abs(x) <= math.Abs(scale)*RelativeEpsilon || math.Abs(x) <= Epsilon
}
