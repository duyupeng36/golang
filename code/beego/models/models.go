package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// Model Struct

type User struct {
	Id       int    `orm:"pk;auto"`
	Name     string `orm:"size(100)"`
	Email    string `orm:"size(255)"`
	Password string `orm:"size(255)"`
}

type Article struct {
	Id       int       `orm:"pk;auto"`
	ArtiName string    `orm:"size(20)"`                                   // 文章名称
	Atime    time.Time `orm:"auto_now;type(datetime)"`                    // 修改时间
	Acount   int       `orm:"default(0);null"`                            // 阅读量
	Acontent string    `orm:"size(500)"`                                  // 内容
	Aimg     string    `orm:"size(100);default(/static/img/default.jpg)"` // 文章图片
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
	orm.RegisterModel(new(User), new(Article)) // 注册表到orm

	// create table
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatalln("创建表失败，失败原因", err)
		return
	} // 创建表
}
