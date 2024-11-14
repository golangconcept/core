package application

import "ddd/domain"

type OrderService struct {
	OrderRepo domain.OrderRepository
}

func NewOrderService(orderRepo domain.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (service *OrderService) CreateOrder(id, orderDate string) (*domain.Order, error) {
	order := domain.NewOrder(id, orderDate)
	err := service.OrderRepo.Save(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (service *OrderService) AddProducToOrder(orderID string, product *domain.Product) (*domain.Order, error) {
	order, err := service.OrderRepo.FindByID(orderID)
	if err != nil {
		return nil, err
	}
	order.AddProduct(product)
	err = service.OrderRepo.Save(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}
