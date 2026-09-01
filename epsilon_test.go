package geom3d

import (
	"fmt"
	"math"
	"testing"
)

func TestAlmostZero(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want bool
	}{
		{name: "exact zero", x: 0, want: true},
		{name: "within epsilon positive", x: Epsilon / 2, want: true},
		{name: "within epsilon negative", x: -Epsilon / 2, want: true},
		{name: "at epsilon positive", x: Epsilon, want: true},
		{name: "at epsilon negative", x: -Epsilon, want: true},
		{name: "outside epsilon positive", x: 2 * Epsilon, want: false},
		{name: "outside epsilon negative", x: -2 * Epsilon, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AlmostZero(tc.x)
			if got != tc.want {
				t.Fatalf("AlmostZero(%v) = %v, want %v", tc.x, got, tc.want)
			}
		})
	}
}

func TestAlmostEqual(t *testing.T) {
	tests := []struct {
		name string
		a    float64
		b    float64
		want bool
	}{
		{name: "equal", a: 1, b: 1, want: true},
		{name: "within epsilon", a: 1, b: 1 + Epsilon/2, want: true},
		{name: "just inside epsilon", a: 1, b: 1 + 0.9*Epsilon, want: true},
		{name: "outside epsilon", a: 1, b: 1 + 2*Epsilon, want: false},
		{name: "negative within epsilon", a: -2, b: -2 + Epsilon/2, want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AlmostEqual(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("AlmostEqual(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestAlmostEqualSymmetric(t *testing.T) {
	a := 1.0
	b := 1.0 + Epsilon/2

	if AlmostEqual(a, b) != AlmostEqual(b, a) {
		t.Fatal("AlmostEqual should be symmetric")
	}
}

func TestAlmostEqualRelative(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want bool
	}{
		{name: "exact equal", a: 1, b: 1, want: true},
		{name: "small values within absolute epsilon", a: 0, b: Epsilon / 2, want: true},
		{name: "small values clearly different", a: 0, b: 1e-3, want: false},
		// The motivating case: at large magnitudes, values that agree to
		// many significant figures differ by far more than the fixed
		// Epsilon in absolute terms, but AlmostEqualRelative still
		// recognizes them as equal.
		{name: "large values within relative tolerance", a: 1_000_000_000, b: 1_000_000_000.5, want: true},
		{name: "large values outside relative tolerance", a: 1_000_000_000, b: 1_000_000_010, want: false},
		{name: "both infinite same sign", a: math.Inf(1), b: math.Inf(1), want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AlmostEqualRelative(tc.a, tc.b)
			if got != tc.want {
				t.Fatalf("AlmostEqualRelative(%v, %v) = %v, want %v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestAlmostEqualRelativeDiffersFromAlmostEqualAtScale(t *testing.T) {
	// This is the actual gap AlmostEqualRelative closes: two large values
	// that AlmostEqual (fixed absolute Epsilon) correctly reports as
	// different, but that are "the same" for any practical purpose at
	// this scale.
	a, b := 1_000_000_000.0, 1_000_000_000.5

	if AlmostEqual(a, b) {
		t.Fatal("expected AlmostEqual to reject this pair (that's the documented limitation)")
	}
	if !AlmostEqualRelative(a, b) {
		t.Fatal("expected AlmostEqualRelative to accept this pair")
	}
}

func TestAlmostEqualRelativeSymmetric(t *testing.T) {
	a, b := 1_000_000_000.0, 1_000_000_000.5

	if AlmostEqualRelative(a, b) != AlmostEqualRelative(b, a) {
		t.Fatal("AlmostEqualRelative should be symmetric")
	}
}

func TestAlmostZeroAtScale(t *testing.T) {
	tests := []struct {
		name  string
		x     float64
		scale float64
		want  bool
	}{
		{name: "within relative tolerance of large scale", x: 0.5, scale: 1_000_000_000, want: true},
		{name: "outside relative tolerance of large scale", x: 10, scale: 1_000_000_000, want: false},
		{name: "within absolute epsilon at zero scale", x: Epsilon / 2, scale: 0, want: true},
		{name: "outside absolute epsilon at zero scale", x: 1e-6, scale: 0, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := AlmostZeroAtScale(tc.x, tc.scale)
			if got != tc.want {
				t.Fatalf("AlmostZeroAtScale(%v, %v) = %v, want %v", tc.x, tc.scale, got, tc.want)
			}
		})
	}
}

func ExampleAlmostEqualRelative() {
	a, b := 1_000_000_000.0, 1_000_000_000.5

	fmt.Println("AlmostEqual:", AlmostEqual(a, b))
	fmt.Println("AlmostEqualRelative:", AlmostEqualRelative(a, b))

	// Output:
	// AlmostEqual: false
	// AlmostEqualRelative: true
}
