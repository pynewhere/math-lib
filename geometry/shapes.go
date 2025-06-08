package geometry

import "math"

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Triangle 三角形结构体
type Triangle struct {
	Base   float64
	Height float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area 计算三角形面积
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter 计算三角形周长（假设是等腰三角形）
func (t Triangle) Perimeter() float64 {
	// 使用勾股定理计算斜边
	side := math.Sqrt(t.Base*t.Base/4 + t.Height*t.Height)
	return t.Base + 2*side
}
