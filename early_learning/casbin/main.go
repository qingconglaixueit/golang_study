package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
	xd "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// myAuth 拦截器
func myAuth(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		// 获取方法
		act := c.Request.Method
		sub := "root"

		// 判断策略是否已经存在了
		if ok := e.Enforce(sub, obj, act); ok {
			log.Println("Check successfully")
			c.Next()
		} else {
			log.Println("sorry , Check failed")
			c.Abort()
		}
	}
}

func main() {
	// 使用自己定义rbac_db
	// 最后的一个参数咱们写true 默认为false,使用缺省的数据库名casbin,不存在则创建
	a := xd.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/mycasbin?charset=utf8", true)

	e := casbin.NewEnforcer("./rbac_models.conf", a)

	//从DB中 load 策略
	e.LoadPolicy()

	//new 一个路由
	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		log.Println("add a policy")
		if ok := e.AddPolicy("root", "/api/v1/hello", "GET"); !ok {
			log.Println("The strategy already exists")
		} else {
			log.Println("add successfully ...")
		}
	})

	//使用自定义拦截器中间件，每一个接口的访问，都会通过这个拦截器
	r.Use(myAuth(e))
	//创建请求
	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("hello wolrd")
	})

	// 监听 127。0.0.1:8888
	r.Run(":8888")
}
