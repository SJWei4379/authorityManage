package middleware

import (
	"authoritymanage/global"
	"authoritymanage/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		//ResBody := c.Request.Body
		//bodyByte, err := ioutil.ReadAll(ResBody)
		//if err != nil {
		//	fmt.Println(err.Error())
		//}
		var body interface{}
		//err = json.Unmarshal(bodyByte, &body)
		//if err != nil {
		//	fmt.Println(err)
		//	body = string(bodyByte)
		//}
		//fmt.Println(body)
		cost := time.Since(start)
		user, _ := c.Get("user")

		//type Duser struct {
		//	Name string `json:"name" bson:"name"`
		//}
		//user := Duser{
		//	Name: "jhghjg",
		//}
		//fmt.Println(user)
		log := model.Log{
			User:      user,
			Path:      path,
			Method:    c.Request.Method,
			Status:    c.Writer.Status(),
			Query:     query,
			Body:      body,
			Ip:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			Errors:    c.Errors.ByType(gin.ErrorTypePrivate).String(),
			Cost:      cost.String(),
		}
		if log.Status == 204 {
			return
		}
		if err := global.DB.Model(&model.Log{}).Create(&log).Error; err != nil {
			fmt.Println("日志存储失败")
		}
	}
}
