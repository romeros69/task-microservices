package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	grpc3 "google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
	"task-microservices/config"
	grpc2 "task-microservices/internal/controller/grpc"
	v1 "task-microservices/internal/controller/http/v1"
	"task-microservices/internal/usecase"
	"task-microservices/internal/usecase/repo"
	"task-microservices/pkg/httpserver"
	"task-microservices/pkg/postgres"
	"time"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal("Error in creating postgres instance")
	}

	conn, err := grpc3.Dial(cfg.GrpcURL, grpc3.WithInsecure())

	if err != nil {
		log.Fatalf("Cannot dial with gRPC connection: %v", err)
	}
	defer conn.Close()

	rpcClient := grpc2.NewRpcServer(conn)

	taskStatusUC := usecase.NewTaskStatusUseCase(repo.NewTaskStatusRepo(pg))
	taskUC := usecase.NewTaskUseCase(repo.NewTaskRepo(pg), taskStatusUC)

	// http server
	handler := gin.New()
	//
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1.NewRouter(handler, taskUC, rpcClient)

	serv := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
