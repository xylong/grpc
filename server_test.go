package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/services"
	"io/ioutil"
	"log"
	"testing"
)

func TestGrpc(t *testing.T) {
	certificate,err:=tls.LoadX509KeyPair("config/client.pem", "config/client.key")
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
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
		ServerName:   "localhost",
	})

	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(creds))
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
