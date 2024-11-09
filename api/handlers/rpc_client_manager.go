package handlers

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"ryg-api-gateway/conf"
	pbu "ryg-api-gateway/gen_proto/user_service"
)

type RpcClientManager struct {
	User pbu.UserServiceClient
}

func NewRpcClientManager(c *conf.Config) *RpcClientManager {
	userConn, err := grpc.NewClient(c.UserService.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	userClient := pbu.NewUserServiceClient(userConn)

	return &RpcClientManager{
		User: userClient,
	}
}
