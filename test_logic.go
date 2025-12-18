package main

import "fmt"

func CalculateItemReduction(originalPrice float64, markdownPercentage int) int {
	val := originalPrice * float64(markdownPercentage)
	reduction := val / 100.0
	return int(reduction)
}

func ProcessOrder(itemName string, unitPrice float64, quantity int) {
	product := CreateProduct(itemName, unitPrice)
	total := product.Price * float64(quantity)
	fmt.Printf("Order total for %s: %.2f\n", product.Name, total)
}

func GetManagerDisplayName(mgr Manager) string {
	id := mgr.GetName()
	if id < 0 {
		return "Unknown"
	}
	return fmt.Sprintf("Manager-%d", id)
}
