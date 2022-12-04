package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kzmijak/zswod_api_go/controller"
	"github.com/kzmijak/zswod_api_go/modules/config"
	"github.com/kzmijak/zswod_api_go/modules/database"
	"github.com/kzmijak/zswod_api_go/modules/logger"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Initialize();
	if err != nil {
		fmt.Printf("Failed to load environmental variables %v", err)
		os.Exit(0)
	}
	
	lgr, err := logger.Initialize(ctx, cfg.Logger); 
	if err != nil {
		fmt.Printf("Failed to instantiate logger service %v", err)
		os.Exit(0)
	}
	
	if err := database.InitDatabase(cfg.Database, ctx); err != nil {
		lgr.Fatal("Failed to initialize database", err)
		os.Exit(0)
	}
	fmt.Print(database.Client)


	controller.NewController().
		WithContext(&ctx).
		WithLogger(*lgr).
		Run(cfg.Server.Host);
}

