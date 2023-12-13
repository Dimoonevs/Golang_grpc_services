package auth

import (
	"fmt"
	pb "my_project/pkg/auth/pb"
	conf "my_project/pkg/config"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *conf.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Could not connect: %d", err)
	}
	return pb.NewAuthServiceClient(cc)
}
