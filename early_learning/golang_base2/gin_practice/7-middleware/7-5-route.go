// 7-5-route.go
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
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
func main() {
	r := gin.Default()
	r.GET("/user", m1, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "index",
		})
	})

	r.Run(":8080")
}