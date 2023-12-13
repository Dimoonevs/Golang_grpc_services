package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/repository"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/utils"
)

type Service struct {
	pb.UnimplementedAuthServiceServer
	Jwt  utils.JwtWrapper
	Repo repository.AuthRepo
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	err := s.Repo.Register(req)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	token, err := s.Repo.Login(req)

	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Service) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return nil, err
	}
	id, err := s.Repo.ValidateToken(claims.Email)
	fmt.Println(id)
	if err != nil && id == 0 {
		return nil, err
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: id,
	}, nil
}
