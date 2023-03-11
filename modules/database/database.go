package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/ent/user"
	"github.com/kzmijak/zswod_api_go/utils/encryption"
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


	if cfg.Env == "dev" {
		if err := seedAdmin(ctx); err != nil {
			return ErrSchemaCreationFail
		}
	}
	
	return nil;
}

func seedAdmin(ctx context.Context) error {
	adminsCount, err := Client.User.Query().Where(user.RoleEQ(user.RoleAdmin)).Count(ctx)
	if err != nil {
		return ErrCouldNotQuery
	}

	if adminsCount > 0 {
		return nil
	}

	adminPassword, _ := encryption.HashString("root")

	err = Client.User.Create().SetID(uuid.UUID{}).SetEmail("root@sporlowd.pl").SetPassword(adminPassword).SetRole(user.RoleAdmin).Exec(ctx)

	return err
}