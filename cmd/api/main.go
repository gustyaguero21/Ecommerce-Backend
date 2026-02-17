package main

import "ecommerce-backend/internal/infrastructure/http/router"

func main() {
	server := router.StartServer()

	server.Run(router.Port)
}
