package repository

import (
	err "github.com/Dimoonevs/go-grpc-auth-svc/pkg/errorResp"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
)

type AuthRepo interface {
	Login(req *pb.LoginRequest) (string, *err.ErrorResp)
	Register(req *pb.RegisterRequest) *err.ErrorResp
	ValidateToken(email string) (int64, *err.ErrorResp)
}
