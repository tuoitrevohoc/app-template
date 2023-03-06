package migrations

import "github.com/tuoitrevohoc/app-template/backend/internal/bootstrap/data"

func AllMigrations() []data.Migration {
	return []data.Migration{
		RolesCreation{},
		AdminCreation{},
	}
}
