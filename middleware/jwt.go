package middleware

import (
	"authoritymanage/model"
	"authoritymanage/utils"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 生成结构体
type CustomClaims struct {
	User               model.User `json:"user" bson:"user"`
	jwt.StandardClaims            // jwt中标准格式,主要是设置token的过期时间
}

// 定义密钥
var Secret = []byte("Wsj@admin")

// 生成token
func CreateToken(user model.User) (string, error) {
	//获取token，前两部分，默认签名方式为HS256
	token := jwt.NewWithClaims(jwt.SigningMethodES256, CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(), //签名生效时间
			//ExpiresAt: time.Now().Unix() + 60*60*24*30, //30添小时过期
			Issuer: "wsj", //签发人，
		},
	})
	//根据密钥生成加密token，token完整三部分 头部，载荷，签证
	tokenString, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}

	//存入redis
	err = utils.Redis{}.SetValue(tokenString, user.Name, 60*60*48)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

// 解析token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//使用签名解析用户传入的token,获取载荷部分数据
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	//在redis中查看token是否过期
	_, err = utils.Redis{}.GetValue(tokenString)
	if err != nil {
		return nil, errors.New("token过期")
	}
	if token != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("token无效")
	}
	return nil, errors.New("token无效")
}

// JWTAuth中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取token,这里从request header取
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token缺失", Data: " "})
			//终止
			c.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{Code: 401, Message: "token过期", Data: err.Error()})
			//终止
			c.Abort()
			return
		}
		//将用户信息存储在上下文
		c.Set("user", claims.User)
		//重新存入redis
		err = utils.Redis{}.SetValue(token, claims.User.Name, 60*60*48)
		if err != nil {
			fmt.Println(err.Error())
		}
		//继续下面的操作
		c.Next()
	}

}
