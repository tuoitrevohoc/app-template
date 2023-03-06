package bootstrap

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data"
	"github.com/tuoitrevohoc/app-template/api/internal/config"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent"
	"github.com/tuoitrevohoc/app-template/api/pkg/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient(config config.Configurations, migrator *data.Migrator) (*ent.Client, error) {
	client, err := ent.Open(dialect.MySQL, config.DBURL)

	if err != nil {
		log.Print("Error connecting to database")
		return nil, err
	}

	err = client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	)

	if err != nil {
		return nil, err
	}

	err = migrator.Execute(
		context.Background(),
		client,
	)

	return client, err
}