package utils

import "math"

// Round 四舍五入到指定小数位
func Round(num float64, precision int) float64 {
	scale := math.Pow10(precision)
	return math.Round(num*scale) / scale
}

// ToRadians 将角度转换为弧度
func ToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// ToDegrees 将弧度转换为角度
func ToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// IsPrime 判断一个数是否为质数
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// GCD 计算最大公约数
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM 计算最小公倍数
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}
