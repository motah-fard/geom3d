package geom3d

import "testing"

func TestAlmostZero(t *testing.T) {
	tests := []struct {
		name string
		x    float64
		want bool
	}{
		{name: "exact zero", x: 0, want: true},
		{name: "within epsilon positive", x: Epsilon / 2, want: true},
		{name: "within epsilon negative", x: -Epsilon / 2, want: true},
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
