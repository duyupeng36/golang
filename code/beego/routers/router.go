package routers

import (
	"beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/orm/list", &controllers.OrmController{}, "get:List")
	beego.Router("/orm/insert", &controllers.OrmController{}, "post:Insert")
}
