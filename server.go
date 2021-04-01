package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/helper"
	"grpc/pbfiles"
	"grpc/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	pbfiles.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	pbfiles.RegisterOrderServiceServer(rpcServer, new(services.OrderService))

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := rpcServer.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
