package handlers

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/services"
)

type OrderHandler struct {
	orderService *services.OrderService
}

func NewOrderHandler(s *services.OrderService) *OrderHandler {
	return &OrderHandler{
		orderService: s,
	}
}

func (h *OrderHandler) HandlePlaceOrder(orderID, userID int, productIDs []int) {
	order, err := h.orderService.PlaceOrder(orderID, userID, productIDs)
	if err != nil {
		fmt.Printf("Order Error: %v\n", err)
		return
	}
	fmt.Printf("Order #%d successfully placed for amount %.2f\n", order.ID, order.TotalAmount)
}

func (h *OrderHandler) PrintOrderSummary(id int) {
	details, err := h.orderService.GetOrderDetails(id)
	if err != nil {
		fmt.Printf("Summary Error: %v\n", err)
		return
	}
	fmt.Println(details)
}
