package statistics

import "math"

// Mean 计算平均值
func Mean(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

// Median 计算中位数
func Median(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	// 排序
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	mid := len(sorted) / 2
	if len(sorted)%2 == 0 {
		return (sorted[mid-1] + sorted[mid]) / 2
	}
	return sorted[mid]
}

// Variance 计算方差
func Variance(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	mean := Mean(numbers)
	sum := 0.0
	for _, num := range numbers {
		diff := num - mean
		sum += diff * diff
	}
	return sum / float64(len(numbers))
}

// StdDev 计算标准差
func StdDev(numbers []float64) float64 {
	return math.Sqrt(Variance(numbers))
}

// Max 返回最大值
func Max(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

// Min 返回最小值
func Min(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}
