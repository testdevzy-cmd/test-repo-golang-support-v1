package services

import (
	"errors"

	"github.com/test-repo/test-repo-golang-support-v1/models"
	"github.com/test-repo/test-repo-golang-support-v1/repositories"
	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) AddProduct(name string, price float64, sku string) (*models.Product, error) {
	if utils.IsEmpty(name) {
		return nil, errors.New("product name cannot be empty")
	}
	product := models.NewProduct(name, price, sku)
	err := s.repo.Save(product)
	if err != nil {
		return nil, err
	}
	y := s.xy(product)
	println(y)
	return product, nil
}

func (s *ProductService) GetProduct(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}
