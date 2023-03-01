package data

import (
	"context"

	"github.com/tuoitrevohoc/app-template/api/ent"
	"github.com/tuoitrevohoc/app-template/api/ent/migration"
	"go.uber.org/zap"
)

type Migration interface {
	Name() string
	Execute(ctx context.Context, tx *ent.Tx) error
}

type Migrator struct {
	log *zap.Logger
	migrations []Migration
}

func NewMigrator(log *zap.Logger, migrations []Migration) *Migrator {
	return &Migrator{
		log: log,
		migrations: migrations,
	}
}

func (m *Migrator) Execute(ctx context.Context, client *ent.Client) error {
	log := m.log
	log.Info("Start migration")
	
	for _, mig := range m.migrations {
		exist, err := client.Migration.Query().Where(
			migration.MigrationEQ(mig.Name()),
		).Exist(ctx)

		if err != nil {
			return err
		}

		if !exist {
			migLog := log.With(zap.String("migration", mig.Name()))
			migLog.Info("Start migration ")
			tx, err := client.Tx(ctx)

			if err != nil {
				migLog.Info("Error creating a transaction", zap.Error(err))
				return err
			}
			
			err = mig.Execute(ctx, tx)

			if err != nil {
				migLog.Error("Error executing migration", zap.Error(err))
				tx.Rollback()
				return err
			}

			_, err = tx.Migration.Create().
				SetMigration(mig.Name()).
				Save(ctx)

			if err != nil {
				migLog.Error("Error saving migration", zap.Error(err))
				_ = tx.Rollback()
				return err
			}

			err = tx.Commit()

			if err != nil {
				migLog.Error("Error commiting transaction", zap.Error(err))
				_ = tx.Rollback()
				return err
			}

			migLog.Info("Migration completed")
		}
	}
	
	return nil
}
