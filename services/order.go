package services

import (
	"context"
	"fmt"
	"grpc/pbfiles"
	"grpc/pbfiles/model"
)

type OrderService struct {
}

func (s *OrderService) NewOrder(ctx context.Context, order *model.Order) (*pbfiles.OrderResponse, error) {
	fmt.Println(order)
	return &pbfiles.OrderResponse{
		Status:  "ok",
		Message: "success",
	}, nil
}
