package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UrlMapping(r *gin.Engine) {

	api := r.Group("api/v1/ecommerce")

	api.GET("/statuscheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "healty",
		})
	})

}
