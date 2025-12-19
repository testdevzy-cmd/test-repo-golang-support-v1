package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Go Application Started ===")

	user := NewUser("Alice", "alice@example.com", 30, true)
	fmt.Printf("Created user: %+v\n", user)

	isValid := ValidateEmail(user.EmailAddress)
	fmt.Printf("Email valid: %t\n", isValid)

	greeting := FormatGreeting(user.FullName, "Ms.")
	fmt.Printf("%s\n", greeting)

	nums := []int{1, 2, 3, 4, 5}
	total := CalculateSum(nums)
	fmt.Printf("Sum of numbers: %d\n", total)

	fmt.Println("=== Application Complete ===")
}
