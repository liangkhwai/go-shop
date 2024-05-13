package main

import (
	"context"
	"log"
	"os"

	"github.com/liangkhwai/go-shop/config"
	"github.com/liangkhwai/go-shop/pkg/database"
	"github.com/liangkhwai/go-shop/server"
)

func main() {
	ctx := context.Background()

	// Initialize config

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}

		return os.Args[1]

	}())

	//Database connection
	db := database.DbConn(ctx, &cfg)
	defer db.Disconnect(ctx)

	// Start server
	server.Start(ctx, &cfg, db)

}
