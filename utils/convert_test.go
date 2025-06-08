package utils

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	tests := []struct {
		num       float64
		precision int
		want      float64
	}{
		{3.14159, 2, 3.14},
		{3.14159, 3, 3.142},
		{3.14159, 0, 3},
		{-3.14159, 2, -3.14},
	}

	for _, tt := range tests {
		if got := Round(tt.num, tt.precision); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Round(%.5f, %d) = %.5f; want %.5f", tt.num, tt.precision, got, tt.want)
		}
	}
}

func TestToRadians(t *testing.T) {
	tests := []struct {
		degrees float64
		want    float64
	}{
		{0, 0},
		{90, math.Pi / 2},
		{180, math.Pi},
		{360, 2 * math.Pi},
	}

	for _, tt := range tests {
		if got := ToRadians(tt.degrees); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("ToRadians(%.2f) = %.2f; want %.2f", tt.degrees, got, tt.want)
		}
	}
}

func TestToDegrees(t *testing.T) {
	tests := []struct {
		radians float64
		want    float64
	}{
		{0, 0},
		{math.Pi / 2, 90},
		{math.Pi, 180},
		{2 * math.Pi, 360},
	}

	for _, tt := range tests {
		if got := ToDegrees(tt.radians); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("ToDegrees(%.2f) = %.2f; want %.2f", tt.radians, got, tt.want)
		}
	}
}

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
		{1, false},
		{0, false},
		{-1, false},
	}

	for _, tt := range tests {
		if got := IsPrime(tt.n); got != tt.want {
			t.Errorf("IsPrime(%d) = %v; want %v", tt.n, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{12, 18, 6},
		{18, 12, 6},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{7, 13, 1},
	}

	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{12, 18, 36},
		{18, 12, 36},
		{0, 5, 0},
		{5, 0, 0},
		{7, 13, 91},
	}

	for _, tt := range tests {
		if got := LCM(tt.a, tt.b); got != tt.want {
			t.Errorf("LCM(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
