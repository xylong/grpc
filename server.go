package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := rpcServer.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
