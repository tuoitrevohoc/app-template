package migrations

import "github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data"

func AllMigrations() []data.Migration {
	return []data.Migration{
		RolesCreation{},
		AdminCreation{},
	}
}
