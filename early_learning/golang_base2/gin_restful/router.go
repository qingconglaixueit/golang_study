package main

import (
	"fmt"
	"gin_http/db"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	Name      string `json:"name" form:"name"`
	Telephone string `json:"telephone" form:"telephone"`
}

//插入
func (person *Person) Create() int64 {
	rs, err := db.SqlDB.Exec("INSERT into users (name, telephone) value (?,?)", person.Name, person.Telephone)
	if err != nil {
		log.Fatal(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return id
}

//查询一条记录
func (p *Person) GetRow() (person Person, err error) {
	person = Person{}
	err = db.SqlDB.QueryRow("select id,name,telephone from users where id = ?", p.Id).Scan(&person.Id, &person.Name, &person.Telephone)
	return
}

//查询所有记录
func (person *Person) GetRows() (persons []Person, err error) {
	rows, err := db.SqlDB.Query("select id,name,telephone from users")
	for rows.Next() {
		person := Person{}
		err := rows.Scan(&person.Id, &person.Name, &person.Telephone)
		if err != nil {
			log.Fatal(err)
		}
		persons = append(persons, person)
	}
	rows.Close()
	return
}

//修改
func (person *Person) Update() int64 {
	rs, err := db.SqlDB.Exec("update users set telephone = ? where id = ?", person.Telephone, person.Id)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

//删除一条记录
func Delete(id int) int64 {
	rs, err := db.SqlDB.Exec("delete from users where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}

//index
func IndexUsers(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

//增加一条记录
func AddUsers(c *gin.Context) {
	name := c.Request.FormValue("name")
	telephone := c.Request.FormValue("telephone")
	fmt.Println("name:", name)
	fmt.Println("telephone:", telephone)
	if name == "" {
		msg := fmt.Sprintf("name字段错误")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": msg,
		})
		return
	}
	person := Person{
		Name:      name,
		Telephone: telephone,
	}
	id := person.Create()
	msg := fmt.Sprintf("insert 成功 %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//获得一条记录
func GetOne(c *gin.Context) {
	ids := c.Param("id")
	id, _ := strconv.Atoi(ids)
	p := Person{
		Id: id,
	}
	rs, _ := p.GetRow()
	c.JSON(http.StatusOK, gin.H{
		"result": rs,
	})
}

//获得所有记录
func GetAll(c *gin.Context) {
	p := Person{}
	rs, _ := p.GetRows()
	c.JSON(http.StatusOK, gin.H{
		"list": rs,
	})
}

func UpdateUser(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	telephone := c.Request.FormValue("telephone")
	person := Person{
		Id:        id,
		Telephone: telephone,
	}
	row := person.Update()
	msg := fmt.Sprintf("updated successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

//删除一条记录
func DelUser(c *gin.Context) {
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	row := Delete(id)
	msg := fmt.Sprintf("delete successful %d", row)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", IndexUsers) //http://192.168.2.132:8806

	//路由群组
	users := router.Group("users")
	{
		users.GET("", GetAll)             //http://192.168.2.132:8806/api/v1/users
		users.POST("/add", AddUsers)      //http://192.168.2.132:8806/api/v1/users/add
		users.GET("/get/:id", GetOne)     //http://192.168.2.132:8806/api/v1/users/get/5
		users.POST("/update", UpdateUser) //http://192.168.2.132:8806/api/v1/users/update
		users.POST("/del", DelUser)       //http://192.168.2.132:8806/api/v1/users/del
	}

	departments := router.Group("api/v1/department")
	{
		departments.GET("", GetAll)             //http://192.168.2.132:8806/api/v1/users
		departments.POST("/add", AddUsers)      //http://192.168.2.132:8806/api/v1/users/add
		departments.GET("/get/:id", GetOne)     //http://192.168.2.132:8806/api/v1/users/get/5
		departments.POST("/update", UpdateUser) //http://192.168.2.132:8806/api/v1/users/update
		departments.POST("/del", DelUser)       //http://192.168.2.132:8806/api/v1/users/del
	}

	return router
}
