package handlers

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/services"
)

type ProductHandler struct {
	productService *services.ProductService
}

func NewProductHandler(s *services.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: s,
	}
}

func (h *ProductHandler) HandleAddProduct(name string, price float64, sku string) {
	product, err := h.productService.AddProduct(name, price, sku)
	if err != nil {
		fmt.Printf("Product Handler Error: %v\n", err)
		return
	}
	fmt.Printf("Product '%s' (SKU: %s) added successfully at price %.2f\n", product.Name, product.SKU, product.Price)
}
