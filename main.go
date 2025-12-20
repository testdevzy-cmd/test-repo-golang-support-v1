package main

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/handlers"
	"github.com/test-repo/test-repo-golang-support-v1/repositories"
	"github.com/test-repo/test-repo-golang-support-v1/services"
	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

func main() {
	fmt.Println("=== Go Application Started ===")

	// Initialize Layers
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Flow via Handler
	userHandler.HandleCreateUser(1, "Alice", "alice@example.com", 30)
	userHandler.HandleGetReport(1)

	// Direct Utils usage
	nums := []int{10, 20, 30}
	sum := utils.CalculateSum(nums)
	fmt.Printf("Final Sum Check: %d\n", sum)

	fmt.Println("=== Application Complete ===")
}
