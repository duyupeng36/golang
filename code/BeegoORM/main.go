package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"log"
	"main/models"
)

func Insert(u *models.User) {
	o := orm.NewOrm() // 获取default数据库

	// 插入数据库
	id, err := o.Insert(u)
	if err != nil {
		log.Fatalln("插入数据失败，失败原因：", err)
		return
	}
	log.Print("插入数据成功，返回数据id为：", id)
}

func Select(name string) (u *models.User, err error) {
	o := orm.NewOrm()
	u = &models.User{Name: name}
	err = o.Read(u, "name")
	if err != nil {
		log.Fatalln("数据查找失败，失败原因：", err)
		return nil, err
	}
	return
}

func Update(oldName, newName string) {
	o := orm.NewOrm()
	u, err := Select(oldName)
	if err != nil {
		return
	}

	u.Name = newName
	num, err := o.Update(u)
	if err != nil {
		log.Fatalln("修改数据失败，失败原因：", err)
		return
	}
	log.Println("影响的数据行数", num)
}

func Delete(name string) {
	o := orm.NewOrm()
	u, err := Select(name)
	if err != nil {
		return
	}
	num, err := o.Delete(u)
	if err != nil {
		log.Fatalln("删除数据失败，失败原因：", err)
		return
	}
	log.Println("影响的数据行数：", num)
}

func SelectWthRe() {
	o := orm.NewOrm()
	var users []*models.User
	qs := o.QueryTable("user")
	num, err := qs.Filter("name", "dyp").All(&users)
	if err != nil {
		log.Fatalln("关联查询失败，失败原因: ", err)
		return
	}
	log.Println("查询出的数据条数为", num)
}

func ExecuteSQL(sql string) {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		log.Fatalln("sql语句执行失败，失败原因: ", err)
		return
	}
	log.Println("查询到的数据条数为: ", num)
	for _, term := range maps {
		fmt.Println(term["id"], ":", term["name"])
	}
}

func InsertMany() {
	o := orm.NewOrm()
	users := []models.User{
		{Name: "slene"},
		{Name: "astaxie"},
		{Name: "unknown"},
	}
	successNums, err := o.InsertMulti(100, users)
	if err != nil {
		log.Fatalln("插入失败，成功条数: ", successNums)
		return
	}
	log.Println("成功条数: ", successNums)
}

func main() {
	//Insert(&models.User{
	//	Name: "dyp",
	//})
	//Insert(&models.User{
	//	Name: "dyp",
	//})
	//Insert(&models.User{
	//	Name: "dyp",
	//})
	//SelectWthRe()
	//ExecuteSQL("select * from user")
	InsertMany()
}
