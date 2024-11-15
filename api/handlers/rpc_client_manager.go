package handlers

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"ryg-api-gateway/conf"
	pbt "ryg-api-gateway/gen_proto/task_service"
	pbu "ryg-api-gateway/gen_proto/user_service"
)

type RpcClientManager struct {
	User      pbu.UserServiceClient
	Task      pbt.TaskServiceClient
	Challenge pbt.ChallengeServiceClient
}

func NewRpcClientManager(c *conf.Config) *RpcClientManager {
	userClient := newUserClient(c)
	taskClient := newTaskClient(c)
	challengeClient := newChallengeClient(c)

	return &RpcClientManager{
		User:      userClient,
		Task:      taskClient,
		Challenge: challengeClient,
	}
}

func newUserClient(c *conf.Config) pbu.UserServiceClient {
	userConn, err := grpc.NewClient(c.UserService.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}
	return pbu.NewUserServiceClient(userConn)
}

func newTaskClient(c *conf.Config) pbt.TaskServiceClient {
	taskConn, err := grpc.NewClient(c.TaskService.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to task service: %v", err)
	}
	return pbt.NewTaskServiceClient(taskConn)
}

func newChallengeClient(c *conf.Config) pbt.ChallengeServiceClient {
	challengeConn, err := grpc.NewClient(c.ChallengeService.Url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to challenge service: %v", err)
	}
	return pbt.NewChallengeServiceClient(challengeConn)
}
