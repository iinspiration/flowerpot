package routers

import (
	"flowerpot/controllers"
	"flowerpot/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitCategoryRoute(routerGroup *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.InitMemberRepositoryDB(db)
	controller := controllers.InitMemberController(repo)

	group := routerGroup.Group("/members")
	{
		group.POST("", controller.Create)
		group.GET("", controller.List)
		group.GET("/:id", controller.Get)
		group.PATCH("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}

	// fe
	feGroup := routerGroup.Group("/members/active")
	{
		feGroup.GET("", controller.ActiveList)
	}

}
