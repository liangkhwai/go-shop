package server

import (
	"log"

	playerPb "github.com/liangkhwai/go-shop/modules/player/playerPb"
	"github.com/liangkhwai/go-shop/modules/player/playerHandler"
	"github.com/liangkhwai/go-shop/modules/player/playerRepository"
	"github.com/liangkhwai/go-shop/modules/player/playerUsecase"
	"github.com/liangkhwai/go-shop/pkg/grpccon"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)
		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)

	}()
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	//Health Check

	player.GET("", s.healthCheckService)

	player.POST("/player/register", httpHandler.CreatePlayer)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	player.POST("/player/add-money", httpHandler.AddPlayerMoney)
}
