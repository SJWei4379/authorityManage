package controller

import (
	"authoritymanage/model"
	"authoritymanage/service"
	"authoritymanage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
	}
	c.JSON(http.StatusOK, service.Register(user))
}
