package geometry

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	tests := []struct {
		radius        float64
		wantArea      float64
		wantPerimeter float64
	}{
		{5, math.Pi * 25, 2 * math.Pi * 5},
		{0, 0, 0},
		{1, math.Pi, 2 * math.Pi},
	}

	for _, tt := range tests {
		circle := Circle{Radius: tt.radius}
		if got := circle.Area(); math.Abs(got-tt.wantArea) > 1e-10 {
			t.Errorf("Circle.Area() = %.2f; want %.2f", got, tt.wantArea)
		}
		if got := circle.Perimeter(); math.Abs(got-tt.wantPerimeter) > 1e-10 {
			t.Errorf("Circle.Perimeter() = %.2f; want %.2f", got, tt.wantPerimeter)
		}
	}
}

func TestRectangle(t *testing.T) {
	tests := []struct {
		width, height float64
		wantArea      float64
		wantPerimeter float64
	}{
		{4, 6, 24, 20},
		{0, 5, 0, 10},
		{3, 3, 9, 12},
	}

	for _, tt := range tests {
		rect := Rectangle{Width: tt.width, Height: tt.height}
		if got := rect.Area(); math.Abs(got-tt.wantArea) > 1e-10 {
			t.Errorf("Rectangle.Area() = %.2f; want %.2f", got, tt.wantArea)
		}
		if got := rect.Perimeter(); math.Abs(got-tt.wantPerimeter) > 1e-10 {
			t.Errorf("Rectangle.Perimeter() = %.2f; want %.2f", got, tt.wantPerimeter)
		}
	}
}

func TestTriangle(t *testing.T) {
	tests := []struct {
		base, height float64
		wantArea     float64
	}{
		{4, 6, 12},
		{0, 5, 0},
		{3, 3, 4.5},
	}

	for _, tt := range tests {
		triangle := Triangle{Base: tt.base, Height: tt.height}
		if got := triangle.Area(); math.Abs(got-tt.wantArea) > 1e-10 {
			t.Errorf("Triangle.Area() = %.2f; want %.2f", got, tt.wantArea)
		}
	}

	// 测试等腰三角形的周长
	triangle := Triangle{Base: 4, Height: 3}
	expectedPerimeter := 4 + 2*math.Sqrt(13) // 使用勾股定理计算
	if got := triangle.Perimeter(); math.Abs(got-expectedPerimeter) > 1e-10 {
		t.Errorf("Triangle.Perimeter() = %.2f; want %.2f", got, expectedPerimeter)
	}
}
