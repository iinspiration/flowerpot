package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/1hopin/go-module/utils"
	"gorm.io/gorm"
)

func SetupRouter(router *gin.Engine, db_product *gorm.DB) {

	createRouteCommon(router)

	V1 := router.Group("/flowerpot/v1")
	{
		InitCategoryRoute(V1, db_product)
	}
}

func createRouteCommon(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		utils.LoggerInfo("Healthcheck", "check healthcheck", nil, nil, false)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product Service is running with CICD.",
		})
	})
}
