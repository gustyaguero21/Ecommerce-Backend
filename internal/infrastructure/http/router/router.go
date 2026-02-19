package router

import (
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	UrlMapping(router)

	return router
}
