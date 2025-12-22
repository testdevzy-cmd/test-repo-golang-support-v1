package repositories

import (
	"errors"

	"github.com/test-repo/test-repo-golang-support-v1/models"
)

type ProductRepository struct {
	products map[int]*models.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: make(map[int]*models.Product),
	}
}

func (r *ProductRepository) Save(p *models.Product) error {
	if p == nil {
		return errors.New("cannot save nil product")
	}
	r.products[p.ID] = p
	return nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	p, ok := r.products[id]
	if !ok {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (r *ProductRepository) ListAll() []*models.Product {
	list := make([]*models.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list
}
