package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close() // 注意这行代码要写在上面err判断的下面
	//需要关掉 下面还没有使用 链接 就关闭了  如果需要关闭 需要在处理完sql后再关闭
}

//func main() {
//Insert
/*r, err := Db.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
if err != nil {
	fmt.Println("exec failed1, ", err)
	return
}
id, err := r.LastInsertId()
if err != nil {
	fmt.Println("exec failed2, ", err)
	return
}

fmt.Println("insert succ:", id)*/

//Select
/*var person []Person
err := Db.Select(&person, "select user_id, username, sex, email from person where user_id=?", 1)
if err != nil {
	fmt.Println("exec failed, ", err)
	return
}

fmt.Println("select succ:", person)*/

//Update
/*res, err := Db.Exec("update person set username=? where user_id=?", "stu0003", 1)
if err != nil {
	fmt.Println("exec failed, ", err)
	return
}
row, err := res.RowsAffected()
if err != nil {
	fmt.Println("rows failed, ", err)
}
fmt.Println("update succ:", row)*/

//Delete
/*res, err := Db.Exec("delete from person where user_id=?", 1)
if err != nil {
	fmt.Println("exec failed, ", err)
	return
}

row, err := res.RowsAffected()
if err != nil {
	fmt.Println("rows failed, ", err)
}

fmt.Println("delete succ: ", row)*/

//事务
//	conn, err := Db.Begin() //开启事务
//	if err != nil {
//		fmt.Println("begin failed :", err)
//		return
//	}
//
//	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback() //回滚
//		return
//	}
//	id, err := r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	fmt.Println("insert succ:", id)
//
//	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	id, err = r.LastInsertId()
//	if err != nil {
//		fmt.Println("exec failed, ", err)
//		conn.Rollback()
//		return
//	}
//	fmt.Println("insert succ:", id)
//
//	conn.Commit() //提交
//}
