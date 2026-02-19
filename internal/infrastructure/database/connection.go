package database

import (
	"context"
	"ecommerce-backend/internal/infrastructure/config"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnection struct {
	Config *config.Config
}

func (dc *DBConnection) Connection() (*pgxpool.Pool, error) {
	ctx := context.Background()
	adminDns := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		dc.Config.DB.User,
		dc.Config.DB.Pass,
		dc.Config.DB.Host,
		dc.Config.DB.Port)

	adminPool, adminErr := pgxpool.New(ctx, adminDns)
	if adminErr != nil {
		return nil, adminErr
	}

	if poolErr := adminPool.Ping(ctx); poolErr != nil {
		return nil, poolErr
	}

	defer adminPool.Close()

	exists, existsErr := checkExists(ctx, adminPool, dc.Config.DB.Name)
	if existsErr != nil {
		return nil, existsErr
	}

	if !exists {
		log.Print("DATABASE NOT FOUND. CREATING....")
		if createDbErr := createDb(ctx, adminPool, dc.Config.DB.Name); createDbErr != nil {
			return nil, fmt.Errorf("error creating database. Error: %w", createDbErr)
		}
	} else {
		log.Print("DATABASE FOUND.CONNECTING.....")
	}

	prodDns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dc.Config.DB.User,
		dc.Config.DB.Pass,
		dc.Config.DB.Host,
		dc.Config.DB.Port,
		dc.Config.DB.Name)

	pool, err := pgxpool.New(ctx, prodDns)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func checkExists(ctx context.Context, pool *pgxpool.Pool, dbName string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1);`

	err := pool.QueryRow(ctx, query, dbName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func createDb(ctx context.Context, pool *pgxpool.Pool, dbName string) error {

	query := "CREATE DATABASE " + pgx.Identifier{dbName}.Sanitize()
	_, err := pool.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
