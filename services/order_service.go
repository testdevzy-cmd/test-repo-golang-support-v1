package services

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/models"
	"github.com/test-repo/test-repo-golang-support-v1/repositories"
)

type OrderService struct {
	orderRepo      *repositories.OrderRepository
	userService    *UserService
	productService *ProductService
}

func NewOrderService(o *repositories.OrderRepository, u *UserService, p *ProductService) *OrderService {
	return &OrderService{
		orderRepo:      o,
		userService:    u,
		productService: p,
	}
}

func (s *OrderService) PlaceOrder(orderID, userID int, productIDs []int) (*models.Order, error) {
	_, err := s.userService.GetUser(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user: %v", err)
	}

	var total float64
	for _, pid := range productIDs {
		product, pErr := s.productService.GetProduct(pid)
		if pErr != nil {
			return nil, fmt.Errorf("invalid product %d: %v", pid, pErr)
		}
		total += product.Price
	}

	order := models.NewOrder(orderID, userID, productIDs, total)
	err = s.orderRepo.Save(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) GetOrderDetails(id int) (string, error) {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return "", err
	}

	user, _ := s.userService.GetUser(order.UserID)
	return fmt.Sprintf("Order #%d for %s - Total: %.2f", order.ID, user.FullName, order.TotalAmount), nil
}
