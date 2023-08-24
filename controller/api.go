package controller

import (
	"authoritymanage/model"
	"authoritymanage/service"
	"authoritymanage/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err))
	}
	c.JSON(http.StatusOK, service.CreateApi(api))
}

func DeleteApi(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
		return
	} else {
		c.JSON(http.StatusOK, service.DeleteApi(id))
	}
}

func UpdateApi(c *gin.Context) {
	var api model.Api
	if err := c.Bind(&api); err != nil {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", err.Error()))
		return
	}
	c.JSON(http.StatusOK, service.UpdateApi(api))
}

func GetApi(c *gin.Context) {
	//conditions := make(map[string]string)
	method := c.Query("method")
	//if  method != "" {
	//	conditions["method"] = method
	//}
	name := c.Query("name")
	//if name != "" {
	//	//i忽略大小写
	//	conditions["name"] = name
	//}
	pageSize := c.Query("pageSize")
	currPage := c.Query("currPage")
	if pageSize == "" || currPage == "" {
		c.JSON(http.StatusOK, utils.ErrorMess("参数错误", nil))
	}
	c.JSON(http.StatusOK, service.GetApi(method, name, pageSize, currPage))
}
