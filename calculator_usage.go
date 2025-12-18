package main

func CalculateFinalPrice(basePrice float64, discountPercent int) float64 {
	discount := CalculateDiscount(basePrice, discountPercent)
	finalPrice := basePrice - discount
	return finalPrice
}
