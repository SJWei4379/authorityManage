package router

import (
	"authoritymanage/middleware"
	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	engine := gin.Default()
	//路由分组
	engine.Use(middleware.Log(), middleware.Cors()) //跨域

	engine.POST("login")

	engine.Use(middleware.JWTAuth()) //权限

	UserRouter(engine)

	RoleRouter(engine)

	ApiRouter(engine)

	return engine
}
