package services

import (
	"context"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 99,
	}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context,size *QuerySize) (*ProdResponseList,error) {
	products:=[]*ProdResponse{
		&ProdResponse{
			ProdStock:21,
		},
		&ProdResponse{
			ProdStock: 49,
		},
		&ProdResponse{
			ProdStock: 81,
		},
		&ProdResponse{
			ProdStock: 108,
		},
	}
	return &ProdResponseList{
		List: products,
	},nil
}