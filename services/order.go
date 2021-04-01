package services

import (
	"context"
	"fmt"
	"grpc/pbfiles"
)

type OrderService struct {
}

// http demo:
// {"no":"123456789","price":10.05,"detail":[{"detail_id":1,"order_no":"1","num":1},{"detail_id":2,"order_no":"2","num":2}]}
func (s *OrderService) NewOrder(ctx context.Context, request *pbfiles.OrderRequest) (*pbfiles.OrderResponse, error) {
	if err := request.Order.Validate(); err != nil {
		return &pbfiles.OrderResponse{
			Status:  "error",
			Message: err.Error(),
		}, err
	}

	fmt.Println(request.Order)
	return &pbfiles.OrderResponse{
		Status:  "ok",
		Message: "success",
	}, nil
}
