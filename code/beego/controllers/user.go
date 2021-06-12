package controllers

import (
	"beego/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
)

type RegisterController struct {
	beego.Controller
}

// Get get请求返回注册页面
func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

// Post 请求执行注册逻辑
func (c *RegisterController) Post() {
	name := c.GetString("userName")
	password := c.GetString("password")

	if name == "" || password == "" {
		c.Ctx.WriteString("用户名或密码不能空")
		return
	}
	o := orm.NewOrm()
	u := models.User{Name: name, Password: password}
	qs := o.QueryTable("user")
	if qs.Filter("name", name).Exist() {
		c.Ctx.WriteString("用户名以存在，跳转到登录")
		c.Redirect("/login", 302)
		return
	}

	id, err := o.Insert(&u)
	if err != nil {
		c.Ctx.WriteString("注册失败")
		log.Fatalln("注册失败，失败原因: ", err)
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("注册成功, 注册id：%d\n跳转到登录", id))
	log.Println("注册成功, 注册id：", id)
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	name := c.GetString("userName")
	password := c.GetString("password")

	if name == "" || password == "" {
		c.Ctx.WriteString("用户名或密码不能空")
		return
	}
	o := orm.NewOrm()
	u := models.User{Name: name}
	qs := o.QueryTable("user")
	if !qs.Filter("name", name).Exist() {
		c.Ctx.WriteString("用户名不存在，跳转到注册")
		c.Redirect("/register", 302)
		return
	}
	err := o.Read(&u, "name")
	if err != nil {
		c.Ctx.WriteString("登录失败")
		log.Fatalln("登录失败")
		return
	}

	if u.Password != password {
		c.Ctx.WriteString("用户名或密码不正确")
		return
	}
	c.Ctx.WriteString("登录成功，跳转到首页")
}
