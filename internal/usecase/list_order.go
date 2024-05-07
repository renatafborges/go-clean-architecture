package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventListedInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventListedInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (u *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := u.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersOutput []OrderOutputDTO

	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	u.OrderListed.SetPayload(ordersOutput)
	u.EventDispatcher.Dispatch(u.OrderListed)

	return ordersOutput, nil
}
