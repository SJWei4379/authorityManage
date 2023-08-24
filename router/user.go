package router

import (
	"authoritymanage/controller"
	"github.com/gin-gonic/gin"
)

func UserRouter(engine *gin.Engine) {
	user := engine.Group("user")
	{
		user.POST("/create", controller.Register)
	}
}
