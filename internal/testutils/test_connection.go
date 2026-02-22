package testutils

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TestConnection struct {
}

func (tc *TestConnection) Connection() (*pgxpool.Pool, error) {
	ctx := context.Background()
	adminTestDns := "postgres://root:gagueroferra21@localhost:5432/postgres?sslmode=disable"

	adminTestPool, adminTestErr := pgxpool.New(ctx, adminTestDns)
	if adminTestErr != nil {
		return nil, adminTestErr
	}

	if poolErr := adminTestPool.Ping(ctx); poolErr != nil {
		return nil, poolErr
	}

	defer adminTestPool.Close()

	exists, existsErr := CheckExists(ctx, adminTestPool, "ecommerce_test")
	if existsErr != nil {
		return nil, existsErr
	}

	if !exists {
		log.Print("DATABASE NOT FOUND. CREATING....")
		if createDbErr := CreateDb(ctx, adminTestPool, "ecommerce_test"); createDbErr != nil {
			return nil, fmt.Errorf("error creating database. Error: %w", createDbErr)
		}
	} else {
		log.Print("DATABASE FOUND.CONNECTING.....")
	}

	prodDns := "postgres://root:gagueroferra21@localhost:5432/ecommerce_test?sslmode=disable"

	pool, err := pgxpool.New(ctx, prodDns)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func CheckExists(ctx context.Context, pool *pgxpool.Pool, dbName string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1);`

	err := pool.QueryRow(ctx, query, dbName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CreateDb(ctx context.Context, pool *pgxpool.Pool, dbName string) error {

	query := "CREATE DATABASE " + pgx.Identifier{dbName}.Sanitize()
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func DropDatabase(ctx context.Context) error {

	adminDns := "postgres://root:gagueroferra21@localhost:5432/postgres?sslmode=disable"

	adminPool, err := pgxpool.New(ctx, adminDns)
	if err != nil {
		return err
	}
	defer adminPool.Close()

	_, err = adminPool.Exec(ctx, `
		SELECT pg_terminate_backend(pid)
		FROM pg_stat_activity
		WHERE datname = $1
		AND pid <> pg_backend_pid();
	`, "ecommerce_test")
	if err != nil {
		return err
	}

	query := "DROP DATABASE IF EXISTS " + pgx.Identifier{"ecommerce_test"}.Sanitize()

	_, err = adminPool.Exec(ctx, query)
	return err
}
