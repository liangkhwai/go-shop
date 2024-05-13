package server

import (
	"log"

	"github.com/liangkhwai/go-shop/modules/inventory/inventoryHandler"
	inventoryPb "github.com/liangkhwai/go-shop/modules/inventory/inventoryPb"
	"github.com/liangkhwai/go-shop/modules/inventory/inventoryRepository"
	"github.com/liangkhwai/go-shop/modules/inventory/inventoryUsecase"
	"github.com/liangkhwai/go-shop/pkg/grpccon"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	// gRpc
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.InventoryUrl)
		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)

	}()
	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	//Health Check

	inventory.GET("", s.healthCheckService)

}
