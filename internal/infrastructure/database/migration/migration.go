package migration

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Migration interface {
	Name() string
	Run(ctx context.Context, pool *pgxpool.Pool) error
}

func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	migrations := []Migration{
		&UserMigration{},
	}

	for _, migration := range migrations {
		fmt.Printf("Running migration: %s\n", migration.Name())

		if err := migration.Run(ctx, pool); err != nil {
			return fmt.Errorf("migration failed [%s]: %w", migration.Name(), err)
		}

		fmt.Printf("Migration completed: %s\n", migration.Name())
	}

	return nil
}
