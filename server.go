package main

import (
	"google.golang.org/grpc"
	"grpc/services"
	"net"
)

func main() {
	rpcServer:=grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer,new(services.ProdService))
	listener,_:=net.Listen("tcp", "8080")
	rpcServer.Serve(listener)
}
