package server

import (
	"log"

	itemPb "github.com/liangkhwai/go-shop/modules/item/itemPb"
	"github.com/liangkhwai/go-shop/modules/item/itemHandler"
	"github.com/liangkhwai/go-shop/modules/item/itemRepository"
	"github.com/liangkhwai/go-shop/modules/item/itemUsecase"
	"github.com/liangkhwai/go-shop/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)

	}()
	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	//Health Check

	item.GET("", s.healthCheckService)
}
