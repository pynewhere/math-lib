package basic

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, tt := range tests {
		if got := Add(tt.a, tt.b); got != tt.want {
			t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{3, 2, 1},
		{1, 1, 0},
		{0, 0, 0},
		{-1, -1, 0},
	}

	for _, tt := range tests {
		if got := Sub(tt.a, tt.b); got != tt.want {
			t.Errorf("Sub(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{2, 3, 6},
		{-2, 3, -6},
		{0, 5, 0},
		{-2, -2, 4},
	}

	for _, tt := range tests {
		if got := Mul(tt.a, tt.b); got != tt.want {
			t.Errorf("Mul(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDiv(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{6, 2, 3},
		{5, 2, 2},
		{0, 5, 0},
		{-6, 2, -3},
	}

	for _, tt := range tests {
		if got := Div(tt.a, tt.b); got != tt.want {
			t.Errorf("Div(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}

	// 测试除以0的情况
	if got := Div(5, 0); got != 0 {
		t.Errorf("Div(5, 0) = %d; want 0", got)
	}
}

func TestMod(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{7, 3, 1},
		{8, 4, 0},
		{5, 2, 1},
		{-7, 3, -1},
	}

	for _, tt := range tests {
		if got := Mod(tt.a, tt.b); got != tt.want {
			t.Errorf("Mod(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}

	// 测试除以0的情况
	if got := Mod(5, 0); got != 0 {
		t.Errorf("Mod(5, 0) = %d; want 0", got)
	}
}
