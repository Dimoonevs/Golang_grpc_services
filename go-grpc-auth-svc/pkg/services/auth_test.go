package services

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/models"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/pb"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/repository"
	"github.com/Dimoonevs/go-grpc-auth-svc/pkg/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func setupServerAndDB(t *testing.T) (context.Context, pb.AuthServiceClient, func()) {
	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	userDb := make(map[string]models.User)
	user1 := models.User{
		Id:       int64(1),
		Email:    "test@example.com",
		Password: "$2a$05$Cerpkotud1AU3cA2FvUN9u9.SDWwzv1hvn95mnEF5B2upYtyuABsW",
	}
	user2 := models.User{
		Id:       int64(2),
		Email:    "test2@example.com",
		Password: "$2a$05$9rp4jTdunzvC.v6w/wxPTeH5HtLFhH7kVE/Xj5s7lo0.N6oURs0pe",
	}
	userDb[user1.Email] = user1
	userDb[user2.Email] = user2

	srv := grpc.NewServer()
	t.Cleanup(func() {
		srv.Stop()
	})
	jwt := utils.JwtWrapper{
		SecretKey:       "r43t18sc",
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}
	srvc := &Service{
		UnimplementedAuthServiceServer: pb.UnimplementedAuthServiceServer{},
		Jwt:                            jwt,
		Repo:                           &repository.FakeDB{UsersFakeDB: userDb, Jwt: jwt},
	}

	pb.RegisterAuthServiceServer(srv, srvc)

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("srv.Serve %v", err)
		}
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("Failed to dial bufnet %v", err)
	}
	client := pb.NewAuthServiceClient(conn)

	return ctx, client, func() {
		for k := range userDb {
			delete(userDb, k)
		}

		srv.Stop()
		lis.Close()
	}
}
func TestService_Regsiter(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res, err := client.Register(ctx, &pb.RegisterRequest{
		Email:    "test3@example.com",
		Password: "password",
	})
	expected := &pb.RegisterResponse{
		Status: http.StatusCreated,
	}
	require.NoError(t, err)
	require.Equal(t, expected.Status, res.Status)

}
func TestService_RegsiterError(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res, err := client.Register(ctx, &pb.RegisterRequest{
		Email:    "test2@example.com",
		Password: "password",
	})
	fmt.Println(err)
	expect := "rpc error: code = Unknown desc = E-Mail already exists"
	require.Error(t, err)
	require.Equal(t, expect, err.Error())
	require.Empty(t, res)

}
func TestService_Login(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res, err := client.Login(ctx, &pb.LoginRequest{
		Email:    "test2@example.com",
		Password: "password",
	})
	require.NoError(t, err)
	require.NotEmpty(t, res.Token)

}
func TestService_LoginError(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res1, errEmail := client.Login(ctx, &pb.LoginRequest{
		Email:    "test4@example.com",
		Password: "password",
	})
	res2, errPassword := client.Login(ctx, &pb.LoginRequest{
		Email:    "test2@example.com",
		Password: "password1",
	})

	require.Equal(t, "rpc error: code = Unknown desc = User not found", errEmail.Error())
	require.Empty(t, res1)
	require.Equal(t, "rpc error: code = Unknown desc = Password is incorrect", errPassword.Error())
	require.Empty(t, res2)

}

func TestService_Validate(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res, err := client.Validate(ctx, &pb.ValidateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzMyMjY3OTAsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6MiwiRW1haWwiOiJ0ZXN0MkBleGFtcGxlLmNvbSJ9.wnVORkrmLzq71TuPIWvumXugH_CryEFXun4Wkzzqsgo",
	})

	require.NotEmpty(t, res)
	require.Empty(t, err)

}
func TestService_ValidateError(t *testing.T) {
	ctx, client, clienup := setupServerAndDB(t)
	defer clienup()

	res, err := client.Validate(ctx, &pb.ValidateRequest{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTAzMjkwMjMsImlzcyI6ImdvLWdycGMtYXV0aC1zdmMiLCJJZCI6NCwiRW1haWwiOiJkaW1hQG1haWwuY29tIn0.hwfkpJKmPUx11ttNWkGRO4H09FhO2jeuFhbxSvb8QKs",
	})

	require.NotEmpty(t, err)
	require.Equal(t, "rpc error: code = Unknown desc = User not found", err.Error())
	require.Empty(t, res)

}
