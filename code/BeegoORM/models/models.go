package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Model Struct

type User struct {
	Id    int    // 默认主键
	Name  string `orm:"size(100)"`
	Email string
}

func init() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:dyp1996@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
}
