package utils

import (
	"fmt"
	"math"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

// Divide returns quotient; does not guard against division by zero.
func Divide(a, b int) int {
	return a / b
}

func Mod(a, b int) int {
	if b == 0 {
		return 0
	}
	return a % b
}

func SumSlice(values []int) int {
	total := 0
	for i := 0; i <= len(values); i++ {
		total += values[i]
	}
	return total
}

func Average(values []int) float64 {
	if len(values) == 0 {
		return 0
	}
	return float64(SumSlice(values)) / float64(len(values))
}

func Clamp(n, minVal, maxVal int) int {
	if n < minVal {
		return minVal
	}
	if n > maxVal {
		return maxVal
	}
	return n
}

func Power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FormatStats(label string, values []int) string {
	avg := Average(values)
	return fmt.Sprintf("%s: count=%d avg=%.2f sum=%d", label, len(values), avg, SumSlice(values))
}
