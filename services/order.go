package services

import (
	"context"
	"fmt"
	"grpc/pbfiles"
)

type OrderService struct {
}

func (s *OrderService) NewOrder(ctx context.Context, request *pbfiles.OrderRequest) (*pbfiles.OrderResponse, error) {
	fmt.Println(request.Order)
	return &pbfiles.OrderResponse{
		Status:  "ok",
		Message: "success",
	}, nil
}
