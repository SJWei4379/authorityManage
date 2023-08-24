package router

import (
	"authoritymanage/controller"
	"github.com/gin-gonic/gin"
)

func ApiRouter(engine *gin.Engine) {
	api := engine.Group("api")
	{
		api.POST("/create", controller.CreateApi)
		api.PUT("/update", controller.UpdateApi)
		api.GET("/get", controller.GetApi)
		api.GET("/delete", controller.DeleteApi)
	}
}
