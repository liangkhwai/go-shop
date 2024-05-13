package playerHandler

import (
	"context"

	playerPb "github.com/liangkhwai/go-shop/modules/player/playerPb"
	"github.com/liangkhwai/go-shop/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerPb.UnimplementedPlayerGrpcServiceServer
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerUsecase.PlayerUsecaseService) *playerGrpcHandler {
	return &playerGrpcHandler{
		playerUsecase: playerUsecase,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {
	// your code here
	return nil, nil
}


func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context,req *playerPb.FindOnePlayerProfileToRefreshReq) (*playerPb.PlayerProfile, error) {
	// your code here
	return nil, nil

}
func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context,req *playerPb.GetPlayerSavingAccountReq) (*playerPb.GetPlayerSavingAccountRes, error) {
	// your code here
	return nil, nil

}