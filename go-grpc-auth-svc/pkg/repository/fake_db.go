package repository

import (
	"fmt"
	"net/http"

	errRes "github.com/Dimoonevs/go-grpc-auth-svc/pkg/errorResp"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/models"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/utils"
)

type FakeDB struct {
	UsersFakeDB map[string]models.User
	Jwt         utils.JwtWrapper
}
type ReqId struct {
	Id int64
}

func (f *FakeDB) Login(req *pb.LoginRequest) (string, *errRes.ErrorResp) {
	var userRes models.User
	if _, ok := f.UsersFakeDB[req.Email]; !ok {
		return "", &errRes.ErrorResp{
			StatusResp: http.StatusNotFound,
			ErrMsg:     "User not found",
		}
	}
	for _, userRes := range f.UsersFakeDB {
		if userRes.Email == req.Email {
			if math := utils.CheckPasswordHash(req.Password, userRes.Password); !math {

				return "", &errRes.ErrorResp{
					StatusResp: http.StatusNotFound,
					ErrMsg:     "Password is incorrect",
				}

			}
		}
	}
	userRes.Email = req.Email
	userRes.Id = f.UsersFakeDB[req.Email].Id
	token, err := f.Jwt.GenerateToken(userRes)
	if err != nil {
		return "", &errRes.ErrorResp{
			StatusResp: http.StatusInternalServerError,
			ErrMsg:     err.Error(),
		}
	}
	return token, nil
}

func (f *FakeDB) Register(req *pb.RegisterRequest) *errRes.ErrorResp {
	for _, user := range f.UsersFakeDB {
		if user.Email == req.Email {
			return &errRes.ErrorResp{
				StatusResp: http.StatusConflict,
				ErrMsg:     "E-Mail already exists",
			}
		}
	}
	f.UsersFakeDB[req.Email] = models.User{
		Id:       3,
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
	}
	return nil
}

func (f *FakeDB) ValidateToken(email string) (int64, *errRes.ErrorResp) {
	_, ok := f.UsersFakeDB[email]
	if !ok {
		return 0, &errRes.ErrorResp{
			StatusResp: http.StatusNotFound,
			ErrMsg:     "User not found",
		}
	}
	fmt.Println("Endpoint reached")

	return f.UsersFakeDB[email].Id, nil
}
