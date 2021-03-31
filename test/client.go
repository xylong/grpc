package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/services"
	"io/ioutil"
	"log"
)

// GODEBUG=x509ignoreCN=0 go run test/client.go

func main() {
	certificate, err := tls.LoadX509KeyPair("config/client.pem", "config/client.key")
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("config/ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
		ServerName:   "localhost",
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	testGetOne(prodClient)
	//testGetMultiple(prodClient)
}

func testGetOne(client services.ProdServiceClient) {
	res, err := client.GetProdStock(context.Background(), &services.ProdRequest{
		ProdId: 1,
		//Area:   services.Area_C,
		Area: -1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func testGetMultiple(client services.ProdServiceClient) {
	res, err := client.GetProdStocks(context.Background(), &services.QuerySize{Size: 10})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
