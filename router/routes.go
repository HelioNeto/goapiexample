package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initilizeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		//mostra o Opening
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "GET Opening",
			})
		})
	}
}
