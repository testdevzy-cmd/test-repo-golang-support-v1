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

	// Initialize User Layers
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Initialize Product Layers
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize Order Layers
	orderRepo := repositories.NewOrderRepository()
	orderService := services.NewOrderService(orderRepo, userService, productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Flow via User Handler
	userHandler.HandleCreateUser(1, "Alice", "alice@example.com", 30)
	userHandler.HandleGetReport(1)

	// Flow via Product Handler
	productHandler.HandleAddProduct("Laptop", 1200.50, "LAP-001")
	productHandler.HandleAddProduct("Mouse", 25.00, "MOU-002")

	// Flow via Order Handler
	orderHandler.HandlePlaceOrder(101, 1, []int{1, 2})
	orderHandler.PrintOrderSummary(101)

	// Direct Utils usage
	nums := []int{10, 20, 30}
	sum := utils.CalculateSum(nums)
	fmt.Printf("Final Sum Check: %d\n", sum)

	fmt.Println("=== Application Complete ===")
}
