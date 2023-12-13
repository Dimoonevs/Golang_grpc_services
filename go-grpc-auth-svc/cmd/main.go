package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/config"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/db"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/repository"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/services"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	log.Println(c.DBUrl)
	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 100,
	}

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)
	repo := repository.AuthPostgres{
		Handler: h,
		Jwt:     jwt,
	}

	serv := services.Service{
		Repo: &repo,
		Jwt:  jwt,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &serv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
