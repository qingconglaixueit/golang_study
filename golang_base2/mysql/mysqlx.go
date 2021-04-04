package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql" // 注释掉后异常 _ 调用初始化函数
)

var db *sqlx.DB

func insertInfo() {

	sqlStr := "insert into user(name,age)values(?,?)"
	res, err := db.Exec(sqlStr, "小王八", 2)
	if err != nil {
		fmt.Printf("Exec err : %v", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("LastInsertId err : %v", err)
		return
	}
	fmt.Printf("id == %d", id)

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("RowsAffected err : %v", rows)
		return
	}
	fmt.Printf("rows == %d", rows)
	return

}

func main() {

	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("Connect err : %v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)

	//插入数据
	insertInfo()

}
