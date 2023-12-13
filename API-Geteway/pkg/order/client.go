package order

import (
	"fmt"
	conf "my_project/pkg/config"
	pb "my_project/pkg/order/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(config *conf.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(config.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Printf("Could not connect: %d", err)
	}
	return pb.NewOrderServiceClient(cc)

}
