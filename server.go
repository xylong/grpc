package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/services"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	certificate,err:=tls.LoadX509KeyPair("config/server.pem", "config/server.key")
	if err!=nil {
		log.Fatal(err)
	}
	certPool:=x509.NewCertPool()
	ca, err:=ioutil.ReadFile("config/ca.pem")
	if err!=nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds:=credentials.NewTLS(&tls.Config{
		Certificates:                []tls.Certificate{certificate},
		ClientAuth:                  tls.RequireAndVerifyClientCert,
		ClientCAs:                  certPool,

	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
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
