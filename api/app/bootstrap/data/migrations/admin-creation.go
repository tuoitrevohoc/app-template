package migrations

import (
	"context"

	"github.com/tuoitrevohoc/app-template/api/ent"
	"github.com/tuoitrevohoc/app-template/api/ent/role"
	"golang.org/x/crypto/bcrypt"
)

type AdminCreation struct {
}

func (r AdminCreation) Name() string {
	return "admin-creation"
}

func (r AdminCreation) Execute(ctx context.Context, tx *ent.Tx) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), 10)

	if err != nil {
		return err
	}

	adminRole, err := tx.Role.Query().Where(role.NameEQ(role.NameAdministrator)).FirstID(ctx)

	if err != nil {
		return err
	}

	tx.User.Create().
		SetUsername("admin").
		SetPassword(string(hashedPassword)).
		SetRoleID(adminRole).
		SaveX(ctx)
	
	return nil
}
