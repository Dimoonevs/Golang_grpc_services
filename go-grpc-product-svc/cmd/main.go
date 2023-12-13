package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/go-grpc-product-svc/pkg/config"
	"github.com/Dimoonevs/go-grpc-product-svc/pkg/db"
	"github.com/Dimoonevs/go-grpc-product-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-product-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	handler := db.Init(conf.DBUrl)

	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Product Svc on", conf.Port)

	service := services.Service{
		H: handler,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductServiceServer(grpcServer, &service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
