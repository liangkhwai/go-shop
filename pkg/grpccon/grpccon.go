package grpccon

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/liangkhwai/go-shop/config"
	authPb "github.com/liangkhwai/go-shop/modules/auth/authPb"
	inventoryPb "github.com/liangkhwai/go-shop/modules/inventory/inventoryPb"
	itemPb "github.com/liangkhwai/go-shop/modules/item/itemPb"
	playerPb "github.com/liangkhwai/go-shop/modules/player/playerPb"
	"github.com/liangkhwai/go-shop/pkg/jwtauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}

	grpcAuth struct {
		secretKey string
	}
)

func (g *grpcAuth) unaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}
	authHeader, ok := md["auth"]
	
	if !ok {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
	}
	
	if len(authHeader) == 0 {
		log.Printf("Error: Metadata not found")
		return nil, errors.New("error: metadata not found")
		
	}
	log.Printf("Metadata: %#v", authHeader)

	claims, err := jwtauth.ParseToken(g.secretKey, string(authHeader[0]))
	if err != nil {
		log.Printf("Error: Parse token failed: %s", err.Error())
		return nil, errors.New("error: token is invalid")
	}
	log.Printf("Claims: %v", claims)
	_ = claims

	return handler(ctx, req)

}

func (g *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}
func (g *grpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Printf("Error: grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: clientConn,
	}, nil

}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcAuth := &grpcAuth{
		secretKey: cfg.ApiSecretKey,
	}

	opts = append(opts, grpc.UnaryInterceptor(grpcAuth.unaryAuthorization))

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("Error: Failed to listen: %v", err)
	}

	return grpcServer, lis
}
