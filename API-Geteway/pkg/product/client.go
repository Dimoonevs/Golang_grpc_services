package product

import (
	"fmt"
	conf "my_project/pkg/config"
	pb "my_project/pkg/product/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(config *conf.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(config.ProductSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Printf("Could not connect: %d", err)
	}
	return pb.NewProductServiceClient(cc)
}
