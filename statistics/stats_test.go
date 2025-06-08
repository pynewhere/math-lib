package statistics

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{0, 0, 0}, 0},
		{[]float64{-1, 1}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := Mean(tt.numbers); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Mean(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}

func TestMedian(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{1, 2, 3, 4}, 2.5},
		{[]float64{0, 0, 0}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := Median(tt.numbers); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Median(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 2},
		{[]float64{0, 0, 0}, 0},
		{[]float64{1, 1, 1}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := Variance(tt.numbers); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("Variance(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}

func TestStdDev(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, math.Sqrt(2)},
		{[]float64{0, 0, 0}, 0},
		{[]float64{1, 1, 1}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := StdDev(tt.numbers); math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("StdDev(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 5},
		{[]float64{-1, -2, -3}, -1},
		{[]float64{0, 0, 0}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := Max(tt.numbers); got != tt.want {
			t.Errorf("Max(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		numbers []float64
		want    float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 1},
		{[]float64{-1, -2, -3}, -3},
		{[]float64{0, 0, 0}, 0},
		{[]float64{}, 0},
	}

	for _, tt := range tests {
		if got := Min(tt.numbers); got != tt.want {
			t.Errorf("Min(%v) = %.2f; want %.2f", tt.numbers, got, tt.want)
		}
	}
}
