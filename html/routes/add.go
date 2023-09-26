package routes

import (
	"add/controllers"


	"github.com/gin-gonic/gin"
)

func Add(router *gin.Engine, controller controllers.AddController)*gin.Engine {
	router.Static("/static", "./")
	router.POST("/add", controller.Add)
	return router
}
