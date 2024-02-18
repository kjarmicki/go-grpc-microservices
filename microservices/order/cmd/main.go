package main

import (
	"log"

	"github.com/kjarmicki/go-grpc-microservices/microservices/order/config"
	"github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/adapters/db"
	"github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/adapters/grpc"
	"github.com/kjarmicki/go-grpc-microservices/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
