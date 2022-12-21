package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "task-microservices/docs"
	"task-microservices/internal/controller/grpc"
	"task-microservices/internal/usecase"
)

func NewRouter(handler *gin.Engine, tc usecase.TaskContract, rpcClient *grpc.RpcClient) {
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := NewAuthMiddle(rpcClient)
	h := handler.Group("/api/v1")
	{
		newTaskRoutes(h, tc, auth)
	}
}
