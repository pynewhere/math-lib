package advanced

import "math"

// Pow 返回x的y次方
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Sqrt 返回x的平方根
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

// Cbrt 返回x的立方根
func Cbrt(x float64) float64 {
	return math.Cbrt(x)
}

// Exp 返回e的x次方
func Exp(x float64) float64 {
	return math.Exp(x)
}

// Log 返回x的自然对数
func Log(x float64) float64 {
	return math.Log(x)
}

// Log10 返回x的以10为底的对数
func Log10(x float64) float64 {
	return math.Log10(x)
}
