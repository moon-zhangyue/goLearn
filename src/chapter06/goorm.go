package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
//orm会默认根据结构体创建table ， orm采用的是linux命名方式  即小写加下划线，且会在名字后面加 s
//会创建user_infos 表
type UserInfo struct {
	gorm.Model
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

//  truncate table 表名
//  数据库需要提前创建好  例如mygorm
//  parseTime是查询结果是否⾃动解析为时间
//  loc是MySQL的时区设置
//  charset是编码方式
func main() {
	fmt.Println("try open mysql connection....")
	db, err := gorm.Open("mysql", "root:123456@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("successful")
	defer db.Close()
	// 自动迁移
	//若该表不存在则创建该表，若该表存在且结构体发生变化则更新表结构
	db.AutoMigrate(&UserInfo{})
	u1 := UserInfo{gorm.Model{}, 1, "xiaozhu", "man", "playing"}
	// 创建记录
	result := db.Create(&u1)
	fmt.Println("result:", result.RowsAffected)
	// 查询
	var u = new(UserInfo)
	//查询一条记录
	db.First(u)
	fmt.Printf("First: %#v\n", u)
	//按照条件查询
	var uu UserInfo
	db.Find(&uu, "name=?", "xiaozhu")
	fmt.Printf("Find: %#v\n", uu)
	// 更新
	db.Model(&uu).Update("hobby", "sing")
	// 删除 ， 此处删除记录，是不会将数据表中的数据删除掉，而是deleted_at 会更新删除时间
	db.Delete(&uu)
}
