// 7-6-route.go
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
func indexHandler(c *gin.Context) {
	fmt.Println("index in ...")
	c.JSON(http.StatusOK, gin.H{
		"msg": "indx",
	})
}
//定义一个中间件:统计耗时
func m1(c *gin.Context)  {
	fmt.Println("m1 in ....")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数 执行indexHandler函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Println("cost:%v\n", cost)
	fmt.Println("m1 out")
}
func m2(c *gin.Context)  {
	fmt.Println("m2 in ....")
	//c.Next() //调用后续的处理函数
	c.Abort() //阻止后续调用
	//return   //return 立即结束m2函数
	//m1 in ....
	//m2 in ....
	//cost:%v
	// 0s
	//m1 out
	fmt.Println("m2 out")
}
func main() {
	r := gin.Default()
	r.GET("/user", m1,m2, func(c *gin.Context) {  //可以单独多个路由
		c.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})
	//[GIN-debug] Listening and serving HTTP on :8080
	//m1 in ....
	//m2 in ....
	//m2 out
	//cost:%v
	// 0s
	//m1 out
	r.Run(":8080")
}