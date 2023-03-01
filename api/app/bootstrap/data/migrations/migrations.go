package migrations

import "github.com/tuoitrevohoc/app-template/api/app/bootstrap/data"

func AllMigrations() []data.Migration {
	return []data.Migration{
		RolesCreation{},
		AdminCreation{},
	}
}
