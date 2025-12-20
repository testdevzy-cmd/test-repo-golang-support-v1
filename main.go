package main

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/services"
	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

func main() {
	fmt.Println("=== Go Application Started ===")

	userService := services.NewUserService()

	user, err := userService.CreateUser(1, "Alice", "alice@example.com", 30)
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("Created user via service: %+v\n", user)

	isValid := utils.ValidateEmail(user.EmailAddress)
	fmt.Printf("Email valid: %t\n", isValid)

	greeting := utils.FormatGreeting(user.FullName, "Ms.")
	fmt.Printf("%s\n", greeting)

	report := userService.PrintUserReport(1)
	fmt.Println(report)

	nums := []int{1, 2, 3, 4, 5}
	total := utils.CalculateSum(nums)
	fmt.Printf("Sum of numbers: %d\n", total)

	fmt.Println("=== Application Complete ===")
}
