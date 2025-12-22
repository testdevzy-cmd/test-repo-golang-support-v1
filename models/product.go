package models

import "time"

type Product struct {
	ID        int
	Name      string
	Price     float64
	SKU       string
	CreatedAt time.Time
}

func NewProduct(name string, price float64, sku string) *Product {
	return &Product{
		Name:      name,
		Price:     price,
		SKU:       sku,
		CreatedAt: time.Now(),
	}
}

func (p *Product) IsAvailable() bool {
	return p.SKU != "" && p.Price > 0
}
