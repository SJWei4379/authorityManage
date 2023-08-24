package controller

import (
	"authoritymanage/model"
	"authoritymanage/service"
	"authoritymanage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err))
		return
	}
	c.JSON(http.StatusOK, service.CreateRole(role))
}

func UpdateRole(c *gin.Context) {
	var role model.Role
	if err := c.Bind(&role); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err))
	}
	c.JSON(http.StatusOK, service.UpdateRole(role))

}

func GetRole(c *gin.Context) {
	name := c.Query("name")
	pageSize := c.Query("pageSize")
	currPage := c.Query("currPage")
	if pageSize == "" && currPage == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
	}
	c.JSON(http.StatusOK, service.GetRole(name, pageSize, currPage))
}
