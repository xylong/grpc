package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/pbfiles"
	"grpc/pbfiles/model"
	"io/ioutil"
	"log"
	"time"
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

	//prodClient := pbfiles.NewProdServiceClient(conn)
	//testGetOne(prodClient)
	//testGetMultiple(prodClient)
	//testGetMusicInfo(prodClient)

	orderClient := pbfiles.NewOrderServiceClient(conn)
	testNewOrder(orderClient)
}

func testGetOne(client pbfiles.ProdServiceClient) {
	res, err := client.GetProdStock(context.Background(), &pbfiles.ProdRequest{
		ProdId: 1,
		//Area:   services.Area_C,
		Area: -1,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func testGetMultiple(client pbfiles.ProdServiceClient) {
	res, err := client.GetProdStocks(context.Background(), &pbfiles.QuerySize{Size: 10})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func testGetMusicInfo(client pbfiles.ProdServiceClient) {
	res, err := client.GetProdInfo(context.Background(), &pbfiles.ProdRequest{
		ProdId: 1,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}

func testNewOrder(client pbfiles.OrderServiceClient) {
	if res, err := client.NewOrder(context.Background(), &pbfiles.OrderRequest{Order: &model.Order{
		No:        "123456789",
		UserId:    9527,
		Price:     3.5,
		CreatedAt: &timestamp.Timestamp{Seconds: time.Now().Unix()},
	}}); err != nil {
		log.Fatalln(err.Error())
	} else {
		fmt.Println(res)
	}
}
