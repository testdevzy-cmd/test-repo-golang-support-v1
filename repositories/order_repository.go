package repositories

import (
	"errors"

	"github.com/test-repo/test-repo-golang-support-v1/models"
)

type OrderRepository struct {
	orders map[int]*models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int]*models.Order),
	}
}

func (r *OrderRepository) Save(o *models.Order) error {
	if o == nil {
		return errors.New("cannot save nil order")
	}
	r.orders[o.ID] = o
	return nil
}

func (r *OrderRepository) GetByID(id int) (*models.Order, error) {
	o, ok := r.orders[id]
	if !ok {
		return nil, errors.New("order not found")
	}
	return o, nil
}

func (r *OrderRepository) ListByUserID(userID int) []*models.Order {
	var userOrders []*models.Order
	for _, o := range r.orders {
		if o.UserID == userID {
			userOrders = append(userOrders, o)
		}
	}
	return userOrders
}
