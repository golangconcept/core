package infrastructure

import (
	"ddd/domain"
	"errors"
	"sync"
)

type InMemoryOrderRepository struct {
	orders map[string]*domain.Order

	mu sync.Mutex
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: make(map[string]*domain.Order),
	}
}

func (repo *InMemoryOrderRepository) Save(order *domain.Order) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.orders[order.ID] = order
	return nil
}

func (repo *InMemoryOrderRepository) FindByID(id string) (*domain.Order, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	order, exists := repo.orders[id]

	if !exists {
		return nil, errors.New("order not founnd")
	}
	return order, nil
}
