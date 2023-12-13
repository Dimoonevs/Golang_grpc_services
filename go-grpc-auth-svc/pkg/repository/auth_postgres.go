package repository

import (
	"net/http"

	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/db"
	err "github.com/Dimoonevs/go-grpc-auth-svc/pkg/errorResp"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/models"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/utils"
)

type AuthPostgres struct {
	Handler db.Handler
	Jwt     utils.JwtWrapper
}

func (r *AuthPostgres) Login(req *pb.LoginRequest) (string, *err.ErrorResp) {

	var user models.User

	if result := r.Handler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error != nil {
		return "", &err.ErrorResp{
			StatusResp: http.StatusNotFound,
			ErrMsg:     "User not found",
		}
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return "", &err.ErrorResp{
			StatusResp: http.StatusNotFound,
			ErrMsg:     "Password is incorrect",
		}
	}

	token, _ := r.Jwt.GenerateToken(user)
	return token, nil
}

func (r *AuthPostgres) Register(req *pb.RegisterRequest) *err.ErrorResp {
	var user models.User

	if result := r.Handler.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
		return &err.ErrorResp{
			StatusResp: http.StatusConflict,
			ErrMsg:     "E-Mail already exists",
		}
	}
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)

	r.Handler.DB.Create(&user)
	return nil
}

func (r *AuthPostgres) ValidateToken(email string) (int64, *err.ErrorResp) {
	var user models.User

	if result := r.Handler.DB.Where(&models.User{Email: email}).First(&user); result.Error != nil {
		return 0, &err.ErrorResp{
			StatusResp: http.StatusNotFound,
			ErrMsg:     "User not found",
		}
	}
	return user.Id, nil
}
