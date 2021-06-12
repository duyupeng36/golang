package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// Model Struct

type User struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"size(100)"`
}

func init() {
	var err error
	// 设置默认的数据库
	err = orm.RegisterDataBase("default", "mysql", "root:dyp1996@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Local")
	if err != nil {
		log.Fatalln("连接数据库失败，失败原因", err)
		return
	} // 连接数据库，给定一个别名

	// register model
	orm.RegisterModel(new(User)) // 注册表到orm

	// create table
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalln("创建表失败，失败原因", err)
		return
	} // 创建表
}
