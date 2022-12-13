package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kzmijak/zswod_api_go/ent"
)

var Client *ent.Client; 


func InitDatabase(cfg DatabaseConfig, ctx context.Context) error  {

	client, err := ent.Open(cfg.DriverName, cfg.DSN )
	if err != nil {
		return ErrConnectionFailed
	}

	Client = client;
	
	
	if err := Client.Schema.Create(ctx); err != nil {
		return ErrSchemaCreationFail
	}
	
	return nil;
}
