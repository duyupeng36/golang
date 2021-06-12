package controllers

import (
	models "beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

type OrmController struct {
	beego.Controller
}

// Insert 增加数据
func (c *OrmController) Insert() {
	// 获取提交的数据
	name := c.GetString("name")
	o := orm.NewOrm()
	// 创建对象
	u := models.User{
		Name: name,
	}
	// 插入对象
	id, err := o.Insert(&u)
	if err != nil {
		log.Fatalln("插入数据失败，失败原因", err)
		return
	}
	log.Println("插入数据成功，返回数据id", id)
}

func (c *OrmController) List() {

}
