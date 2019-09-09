package routers

import (
	"web/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.IndexController{}, "put:Get")

	beego.Router("blog/list", &controllers.BlogController{}, "get:List")
	beego.Router("blog/info", &controllers.BlogController{}, "get:Info")
	beego.Router("blog/add", &controllers.BlogController{}, "post:Add")
	beego.Router("blog/del", &controllers.BlogController{}, "delete:Del")

}
