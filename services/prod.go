package services

import (
	"context"
	"grpc/helper"
	"grpc/pbfiles"
)

type ProdService struct {
}

func (p *ProdService) GetProdStock(ctx context.Context, req *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	var stock int32 = 0
	switch req.Area {
	case pbfiles.Area_A:
		stock = 10
	case pbfiles.Area_B:
		stock = 20
	case pbfiles.Area_C:
		stock = 30
	default:
		stock = 60
	}
	return &pbfiles.ProdResponse{
		ProdStock: stock,
	}, nil
}

func (p *ProdService) GetProdStocks(ctx context.Context, size *pbfiles.QuerySize) (*pbfiles.ProdResponseList, error) {
	products := []*pbfiles.ProdResponse{
		&pbfiles.ProdResponse{ProdStock: 21},
		&pbfiles.ProdResponse{ProdStock: 49},
		&pbfiles.ProdResponse{ProdStock: 81},
		&pbfiles.ProdResponse{ProdStock: 108},
	}
	return &pbfiles.ProdResponseList{
		List: products,
	}, nil
}

func (p *ProdService) GetProdInfo(ctx context.Context, request *pbfiles.ProdRequest) (*pbfiles.Music, error) {
	songs := []string{"秋日的私语", "水边的阿荻丽娜", "梦中的婚礼", "星空", "致爱丽丝", "夜的钢琴曲", "童年"}
	index := helper.GetRandomInt(len(songs) - 1)
	return &pbfiles.Music{
		Id:     int32(index) + 1,
		Type:   pbfiles.Type_Popular,
		Name:   songs[index],
		Price:  helper.GetRandomFloat(100),
		Author: "Richard Clayderman",
	}, nil
}
