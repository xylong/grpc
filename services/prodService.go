package services

import (
	"context"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, in *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 100,
	}, nil
}
