package bootstrap

import (
	"context"
	"log"

	"entgo.io/ent/dialect"
	"github.com/tuoitrevohoc/app-template/api/app/config"
	"github.com/tuoitrevohoc/app-template/api/ent"
	"github.com/tuoitrevohoc/app-template/api/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func NewEntClient(config config.Configurations) (*ent.Client, error) {
	client, err := ent.Open(dialect.MySQL, config.DBURL)

	if err != nil {
		log.Print("Error connecting to database")
		return nil, err
	}

	err = client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	)

	return client, err
}