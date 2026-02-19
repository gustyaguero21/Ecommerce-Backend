package migration

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserMigration struct{}

func (um *UserMigration) Name() string {
	return "creating user migration"
}

func (um *UserMigration) Run(ctx context.Context, pool *pgxpool.Pool) error {
	query := `CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		dni TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		telephone TEXT UNIQUE NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW());
	`
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
