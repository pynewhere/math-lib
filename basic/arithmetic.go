package basic

// Add 返回两个整数的和
func Add(a, b int) int {
	return a + b
}

// Sub 返回两个整数的差
func Sub(a, b int) int {
	return a - b
}

// Mul 返回两个整数的积
func Mul(a, b int) int {
	return a * b
}

// Div 返回两个整数的商，如果除数为0则返回0
func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// Mod 返回两个整数的余数，如果除数为0则返回0
func Mod(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}
