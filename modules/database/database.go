package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kzmijak/zswod_api_go/ent"
	"github.com/kzmijak/zswod_api_go/models/role"
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

	if err := seedRoles(ctx); err != nil {
		return ErrSchemaCreationFail
	}
	
	return nil;
}

func seedRoles(ctx context.Context) error {
	rolesCount, err := Client.Role.Query().Count(ctx)
	if err != nil {
		return ErrCouldNotQuery
	}

	if rolesCount > 0 {
		return nil
	}

	roles := []string{ 
		role.Admin.String(), 
		role.LegalGuardian.String(), 
		role.Student.String(), 
		role.Teacher.String(), 
		role.Unknown.String(),
	 }

	bulk := make([]*ent.RoleCreate, len(roles))

	for i, id := range roles {
		bulk[i] = Client.Role.Create().SetID(id)
	}

	_, err = Client.Role.CreateBulk(bulk...).Save(ctx)

	return err
}