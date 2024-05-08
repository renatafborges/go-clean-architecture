package service

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/grpc/pb"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type ListOrderService struct {
	pb.UnimplementedListOrderServiceServer
	ListOrderUseCase usecase.ListOrderUseCase
}

func NewListOrderService(listOrderUseCase usecase.ListOrderUseCase) *ListOrderService {
	return &ListOrderService{
		ListOrderUseCase: listOrderUseCase,
	}
}

func (s *ListOrderService) ListOrder(ctx context.Context, in *pb.Blank) (*pb.OrderList, error) {

	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var ordersResponse []*pb.Order

	for _, order := range orders {
		orderResponse := &pb.Order{
			Id:         order.ID,
			Price:      (float32)(order.Price),
			Tax:        (float32)(order.Tax),
			FinalPrice: (float32)(order.FinalPrice),
		}

		ordersResponse = append(ordersResponse, orderResponse)
	}

	return &pb.OrderList{Orders: ordersResponse}, nil
}
