package controllers

import (
	"beego/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

// List 后台首页 文章列表展示
func (c *ArticleController) List() {
	o := orm.NewOrm() // 获取默认数据库对象
	var articles []models.Article
	qs := o.QueryTable("article")
	_, err := qs.All(&articles)
	if err != nil {
		c.Data["articles"] = articles
		c.TplName = "index.html"
		return
	}
	c.Data["articles"] = articles
	c.TplName = "index.html"
}

// AddArticle 请求方式为get 添加文章的页面展示
func (c *ArticleController) AddArticle() {
	c.TplName = "add.html"
}

// Add 请求方式为post添加文章
func (c *ArticleController) Add() {
	/*
		1. 获取添加文章的标题和内容
		2. 获取上传的文件，并重命名，一面图片重复，保存文件，并记录文件所在的目录集文件名
		3. 保存文章数据集文件路径
		4. 跳转到文章列表页
	*/

	o := orm.NewOrm() // 获取默认数据库对象

	articleName := c.GetString("articleName") // 文章标题
	content := c.GetString("content")         // 文章内容
	if articleName == "" || content == "" {
		c.Data["articleName"] = articleName
		c.Data["content"] = content
		c.Data["errmsg"] = "文章标题或文章内容不能为空"
		c.TplName = "add.html"
		return
	}

	_, head, err := c.GetFile("uploadname") // 获取前端传过来的文件，返回文件本身，文件头，错误
	if err != nil {
		// 出现错误，没有上传文件
		article := models.Article{
			ArtiName: articleName,
			Acontent: content,
		}

		_, err := o.Insert(&article)
		if err != nil {
			c.Data["articleName"] = articleName
			c.Data["content"] = content
			c.Data["errmsg"] = "文章添加失败，请重新添加"
			c.TplName = "add.html"
			log.Fatalln("保存文章出现错误，错误原因为：", err)
			return
		}
		c.Redirect("/", 302) // 文章上传成功，跳转到文章列表页面
		return
	}

	// 上传了文件
	// 校验上传文件的格式
	fileExt := path.Ext(head.Filename)
	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jpeg" {
		c.Data["articleName"] = articleName
		c.Data["content"] = content
		c.Data["errmsg"] = "文件格式不正确，接收jpg png jpeg格式文件"
		c.TplName = "add.html"
		return
	}

	// 校验文件大小
	if head.Size > 5000000 {
		c.Data["articleName"] = articleName
		c.Data["content"] = content
		c.Data["errmsg"] = "文件大小不能超过 5000000"
		c.TplName = "add.html"
		return
	}

	// 重命名文件 以免文件重复
	fileName := time.Now().Format("2006-01-02 15-04-05.000")
	filePath := "/static/img/" + fileName + fileExt

	// 存储
	err = c.SaveToFile("uploadname", "."+filePath)
	if err != nil {
		c.Data["articleName"] = articleName
		c.Data["content"] = content
		c.Data["errmsg"] = "文件上传失败，请重新上传文件"
		c.TplName = "add.html"
		log.Fatalln("保存文件出现错误，错误原因为：", err)
		return
	}

	article := models.Article{
		ArtiName: articleName,
		Acontent: content,
		Aimg:     filePath,
	}

	_, err = o.Insert(&article)
	if err != nil {
		c.Data["articleName"] = articleName
		c.Data["content"] = content
		c.Data["errmsg"] = "文章添加失败，请重新添加"
		c.TplName = "add.html"
		log.Fatalln("保存文章出现错误，错误原因为：", err)
		return
	}
	c.Redirect("/", 302) // 文章添加成功，跳转到文章列表页面
}
