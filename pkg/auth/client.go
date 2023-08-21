package auth

import (
	"fmt"

	"flotta-home/mindbond/api-gateway/pkg/auth/pb"
	"flotta-home/mindbond/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.AuthServiceUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
