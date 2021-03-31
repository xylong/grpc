package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpc/helper"
	"grpc/pbfiles"
	"log"
	"net/http"
)

// GODEBUG=x509ignoreCN=0 go run httpServer.go
func main() {
	gwMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	if err := pbfiles.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwMux, "localhost:8081", opt); err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwMux,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
