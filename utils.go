package main

import (
	"errors"
	"fmt"
	"strings"
)

func ValidateEmail(email string) bool {
	if email == "" {
		return false
	}
	if !strings.Contains(email, "@") {
		return false
	}
	return true
}

func FormatGreeting(name string) string {
	greeting := "Hello"
	if name != "" {
		greeting = fmt.Sprintf("Hello, %s!", name)
	}
	return greeting
}

func CalculateSum(numbers []int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func CalculateAverage(numbers []int) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot calculate average of empty slice")
	}
	sum := CalculateSum(numbers)
	avg := float64(sum) / float64(len(numbers))
	return avg, nil
}

func FilterPositive(numbers []int) []int {
	result := make([]int, 0)
	for _, num := range numbers {
		if num > 0 {
			result = append(result, num)
		}
	}
	return result
}

func MapToUpper(items []string) []string {
	result := make([]string, len(items))
	for i, item := range items {
		result[i] = strings.ToUpper(item)
	}
	return result
}

func Contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func MergeSlices(slice1, slice2 []int) []int {
	result := make([]int, 0, len(slice1)+len(slice2))
	result = append(result, slice1...)
	result = append(result, slice2...)
	return result
}

func ApplyDiscount(price float64, discount float64) float64 {
	if discount < 0 || discount > 100 {
		return price
	}
	discountAmount := price * (discount / 100)
	finalPrice := price - discountAmount
	return finalPrice
}

func FormatCurrency(amount float64, currency string) string {
	formatted := fmt.Sprintf("%.2f %s", amount, currency)
	return formatted
}

func ConvertToPercentage(value float64, total float64) float64 {
	if total == 0 {
		return 0
	}
	percentage := (value / total) * 100
	return percentage
}

func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}
