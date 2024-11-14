package main

import (
	"ddd/application"
	"ddd/domain"
	"ddd/infrastructure"
	"fmt"
)

func main() {
	orderRepo := infrastructure.NewInMemoryOrderRepository()
	orderService := application.NewOrderService(orderRepo)

	order, _ := orderService.CreateOrder("order 1", "2024-11-13")

	product := domain.NewProduct("product1", "Laptop", domain.NewMoney(1000, "USD"))
	updatedOrder, _ := orderService.AddProducToOrder(order.ID, product)

	fmt.Printf("Order ID: %s\n", updatedOrder.ID)
	fmt.Printf("Order Total: %.2f %s\n", updatedOrder.Total.Amount, updatedOrder.Total.Currency)
}
