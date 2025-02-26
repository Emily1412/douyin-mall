package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"product-service/api"
	"product-service/config"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	repo, err := repository.NewProductRepository(cfg.Database.DSN)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}

	productService := service.NewProductService(repo)
	productHandler := handler.NewProductHandler(productService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterProductCatalogServiceServer(s, productHandler)

	log.Printf("Server is listening on port %d", cfg.Server.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
