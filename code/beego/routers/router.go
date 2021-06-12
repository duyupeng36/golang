package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ArticleController{}, "get:List")
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})

	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle;post:Add")
}
