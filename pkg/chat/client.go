package chat

import (
	"fmt"

	"flotta-home/mindbond/api-gateway/pkg/chat/pb"
	"flotta-home/mindbond/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ChatServiceClient
}

func InitServiceClient(c *config.Config) pb.ChatServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.ChatServiceUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewChatServiceClient(cc)
}
