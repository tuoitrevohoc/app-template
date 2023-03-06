package migrations

import (
	"context"

	"github.com/tuoitrevohoc/app-template/backend/pkg/ent"
	"github.com/tuoitrevohoc/app-template/backend/pkg/ent/role"
)

type RolesCreation struct {
}

func (r RolesCreation) Name() string {
	return "roles-creations"
}

func (r RolesCreation) Execute(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.Role.Create().SetName(role.NameUser).
		SetDescription("Role of a user").
		Save(ctx)

	if err != nil {
		return err
	}

	_, err = tx.Role.Create().SetName(role.NameAdministrator).
		SetDescription("Role of an adminstrator").
		Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
