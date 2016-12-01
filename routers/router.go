package routers

import (
	"kiyama/controllers"
	"github.com/astaxie/beego"
)

func init() {
	home :=&controllers.HomeController{}
	beego.Router("/",home,"get:Home")
	beego.Router("/article/:id", home, "get:Article")
	beego.Router("/category/:id",home,"get:Categroy")
}
