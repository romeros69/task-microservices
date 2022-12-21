package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"task-microservices/internal/controller/grpc"
	"task-microservices/internal/controller/grpc/gen"
)

type AuthMiddle struct {
	rpcClient *grpc.RpcClient
}

func NewAuthMiddle(rpcClient *grpc.RpcClient) *AuthMiddle {
	return &AuthMiddle{
		rpcClient: rpcClient,
	}
}

func (a *AuthMiddle) authMiddle(c *gin.Context) error {
	auth := c.GetHeader("Authorization")
	if auth == "" {
		errorResponse(c, http.StatusUnauthorized, "error in header format")
	}
	headerParts := strings.Split(auth, " ")
	if len(headerParts) != 2 {
		errorResponse(c, http.StatusUnauthorized, "cannot parse token")
		return fmt.Errorf("invalid header patr")
	}
	if headerParts[0] != "Bearer:" {
		errorResponse(c, http.StatusUnauthorized, "cannot find Bearer")
		return fmt.Errorf("invalid bearer")
	}
	result, err := a.rpcClient.CheckAuthorization(c.Request.Context(), &gen.CheckAuthRequest{Bearer: headerParts[1]})
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "error in checking token")
		return fmt.Errorf("invalid rpc")
	}
	if result.Error != "" {
		errorResponse(c, http.StatusBadRequest, result.Error)
		return fmt.Errorf("invalid result")
	}
	if !result.Result {
		errorResponse(c, http.StatusUnauthorized, "no access")
		return fmt.Errorf("no access")
	}
	return nil
}

type Creds struct {
	Email string `json:"email"`
}

func (a *AuthMiddle) getUserCreds(ctx context.Context, token string) (Creds, error) {
	res, err := a.rpcClient.GetUserCreds(ctx, &gen.UserCredentialsRequest{Bearer: token})
	if err != nil {
		return Creds{}, err
	}
	if res.Error != "" {
		return Creds{}, fmt.Errorf(res.Error)
	}
	return Creds{
		Email: res.Email,
	}, err
}
