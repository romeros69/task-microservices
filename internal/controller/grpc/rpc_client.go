package grpc

import (
	"context"
	"google.golang.org/grpc"
	"task-microservices/internal/controller/grpc/gen"
)

type RpcClient struct {
	client gen.AuthRegClient
}

func NewRpcServer(rpcConn *grpc.ClientConn) *RpcClient {
	return &RpcClient{
		client: gen.NewAuthRegClient(rpcConn),
	}
}

func (r *RpcClient) CheckAuthorization(ctx context.Context, auReq *gen.CheckAuthRequest) (*gen.CheckAuthResponse, error) {
	return r.client.CheckAuthorization(ctx, auReq)
}

func (r *RpcClient) GetUserCreds(ctx context.Context, ucReq *gen.UserCredentialsRequest) (*gen.UserCredentialsResponse, error) {
	return r.client.GetUserCredentials(ctx, ucReq)
}
