package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 注释掉后异常 _ 调用初始化函数
)

func insertPiceInfo(db *sql.DB) {
	// ？ 作为占位符号
	sqlInfo := "insert into user(name,age)values(?,?)"
	ret, err := db.Exec(sqlInfo, "小猪11号", 22)
	if err != nil {
		fmt.Println("Exec err : ", err)
		return
	}

	//插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId err : ", err)
		return
	}

	fmt.Println("LastInsertId == ", id)

	//本次操作影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err : ", err)
		return
	}
	fmt.Println("rows == ", rows)
}

func deletePiceInfo(db *sql.DB) {
	sqlInfo := "delete from user where id=2"
	ret, err := db.Exec(sqlInfo)
	if err != nil {
		fmt.Println("Exec err : ", err)
		return
	}

	//本次操作影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err : ", err)
		return
	}
	fmt.Println("rows == ", rows)

}

func updatePiceInfo(db *sql.DB) {
	sqlInfo := "update user set name='土猪飞天号' where id=4"
	ret, err := db.Exec(sqlInfo)
	if err != nil {
		fmt.Println("Exec err : ", err)
		return
	}

	//本次操作影响的行数
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected err : ", err)
		return
	}
	fmt.Println("rows == ", rows)

}

type myInfo struct {
	id   int
	name string
	age  int
}

func selectInfo(db *sql.DB) {
	sqlInfo := "select * from user"
	rows, err := db.Query(sqlInfo)
	if err != nil {
		fmt.Println("Exec err : ", err)
		return
	}
	defer rows.Close()
	//输出查询出来的行数

	for rows.Next() {
		var u myInfo
		rows.Scan(&u.id, &u.name, &u.age)
		fmt.Printf("id = %d, name = %s, age = %d\n", u.id, u.name, u.age)
	}

}

//预处理 插入数据操作
func prepareInfo(db *sql.DB) {
	sqlInfo := "insert into user (name,age)values(?,?)"

	stmt, err := db.Prepare(sqlInfo)
	if err != nil {
		fmt.Println("Exec err : ", err)
		return
	}

	ret, err := stmt.Exec("花猪2", 28)
	if err != nil {
		fmt.Println("stmt Exec err : ", err)
		return
	}
	ret, err = stmt.Exec("花猪3", 28)
	if err != nil {
		fmt.Println("stmt Exec err : ", err)
		return
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("stmt Exec err : ", err)
		return
	}
	fmt.Println("rows = ", rows)

}

//事务处理
func trasaction(db *sql.DB) {

	//开启一个事务

	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("Begin err :%v", err)
		return
	}

	sqlStr := "update user set name='小猫住6' where id=?"
	_, err = tx.Exec(sqlStr, 9)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("Exec err :%v", err)
		return
	}

	sqlStr = "update user set name='小猫住7' where id=?"
	_, err = tx.Exec(sqlStr, 6)
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("Exec err :%v", err)
		return
	}

	//提交事务
	err = tx.Commit()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Printf("Commit err :%v", err)
		return
	}

	fmt.Println("commit success ")
}

func main() {
	//打开mysql 数据库 进行连接 , 必须要ping 通 才算是连接上mysql 数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4")
	if err != nil {
		fmt.Println("Open err : ", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping err : ", err)
		return
	}

	defer db.Close()
	//insertPiceInfo(db)
	//deletePiceInfo(db)
	//updatePiceInfo(db)
	//selectInfo(db)

	//prepareInfo(db)
	trasaction(db)


}
