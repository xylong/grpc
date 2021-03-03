package services

import (
	"context"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 100,
	}, nil
}
