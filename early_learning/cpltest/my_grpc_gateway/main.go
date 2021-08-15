package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {好问题
	v1 := router.Group("v1")
	{
		v1.POST("login", func(context *gin.Context) {
			context.String(http.StatusOK, "v1 login")
		})
	}
}
