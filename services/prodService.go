package services

import (
	"context"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.Area {
	case Area_A:
		stock = 10
	case Area_B:
		stock = 20
	case Area_C:
		stock = 30
	default:
		stock = 60
	}
	return &ProdResponse{
		ProdStock: stock,
	}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	products := []*ProdResponse{
		&ProdResponse{ProdStock: 21},
		&ProdResponse{ProdStock: 49},
		&ProdResponse{ProdStock: 81},
		&ProdResponse{ProdStock: 108},
	}
	return &ProdResponseList{
		List: products,
	}, nil
}
