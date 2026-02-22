package main

import (
	"context"
	"ecommerce-backend/internal/infrastructure/config"
	"ecommerce-backend/internal/infrastructure/database"
	"ecommerce-backend/internal/infrastructure/database/migration"
	"ecommerce-backend/internal/infrastructure/database/user"
	"ecommerce-backend/internal/infrastructure/http/router"
	"log"
)

func main() {
	ctx := context.Background()

	cfg, cfgErr := config.LoadEnv()
	if cfgErr != nil {
		log.Fatal(cfgErr)
	}

	dbConn := database.DBConnection{Config: cfg}

	pool, poolErr := dbConn.Connection()
	if poolErr != nil {
		log.Fatal(poolErr)
	}

	if migrateErr := migration.RunMigrations(ctx, pool); migrateErr != nil {
		log.Fatal(migrateErr)
	}

	repo := user.NewUserRepository(pool)

	user, err := repo.GetUserByID(ctx, "1")
	if err != nil {
		log.Print(err)
	}

	log.Print(user)

	router := router.StartServer()

	router.Run(cfg.HTTPConfig.Port)
}
