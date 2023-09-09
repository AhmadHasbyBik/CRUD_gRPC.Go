package main

import (
	"grpc_go/cmd/config"
	"grpc_go/cmd/services"
	productPb "grpc_go/pb/product"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":5000"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listened %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("failed to serve %v", err.Error())
	}
}
