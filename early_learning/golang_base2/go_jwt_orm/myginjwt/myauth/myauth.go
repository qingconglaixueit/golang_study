package myauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my/controller"
	"net/http"
)

//身份认证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//拿到token
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "token为空，请携带token",
				"data":   nil,
			})
			c.Abort()
			return
		}

		fmt.Println("token = ", token)

		//解析出实际的载荷
		j := controller.NewJWT()

		claims, err := j.ParserToken(token)
		if err != nil {
			// token过期
			if err == controller.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token授权已过期，请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// 解析到具体的claims相关信息
		c.Set("claims", claims)
	}
}
