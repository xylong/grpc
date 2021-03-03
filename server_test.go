package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc/services"
	"testing"
)

func TestGrpc(t *testing.T) {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	res, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 1})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
