package advanced

import (
	"math"
	"testing"
)

func TestPow(t *testing.T) {
	tests := []struct {
		x, y, want float64
	}{
		{2, 3, 8},
		{2, 0, 1},
		{2, -1, 0.5},
		{0, 2, 0},
	}

	for _, tt := range tests {
		if got := Pow(tt.x, tt.y); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Pow(%.2f, %.2f) = %.2f; want %.2f", tt.x, tt.y, got, tt.want)
		}
	}
}

func TestSqrt(t *testing.T) {
	tests := []struct {
		x, want float64
	}{
		{16, 4},
		{0, 0},
		{2, math.Sqrt(2)},
	}

	for _, tt := range tests {
		if got := Sqrt(tt.x); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Sqrt(%.2f) = %.2f; want %.2f", tt.x, got, tt.want)
		}
	}
}

func TestCbrt(t *testing.T) {
	tests := []struct {
		x, want float64
	}{
		{27, 3},
		{0, 0},
		{8, 2},
	}

	for _, tt := range tests {
		if got := Cbrt(tt.x); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Cbrt(%.2f) = %.2f; want %.2f", tt.x, got, tt.want)
		}
	}
}

func TestExp(t *testing.T) {
	tests := []struct {
		x, want float64
	}{
		{0, 1},
		{1, math.E},
		{2, math.E * math.E},
	}

	for _, tt := range tests {
		if got := Exp(tt.x); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Exp(%.2f) = %.2f; want %.2f", tt.x, got, tt.want)
		}
	}
}

func TestLog(t *testing.T) {
	tests := []struct {
		x, want float64
	}{
		{1, 0},
		{math.E, 1},
		{math.E * math.E, 2},
	}

	for _, tt := range tests {
		if got := Log(tt.x); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Log(%.2f) = %.2f; want %.2f", tt.x, got, tt.want)
		}
	}
}

func TestLog10(t *testing.T) {
	tests := []struct {
		x, want float64
	}{
		{1, 0},
		{10, 1},
		{100, 2},
	}

	for _, tt := range tests {
		if got := Log10(tt.x); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Log10(%.2f) = %.2f; want %.2f", tt.x, got, tt.want)
		}
	}
}
