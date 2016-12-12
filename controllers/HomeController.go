package controllers

import (
	"github.com/astaxie/beego"
	"kiyama/models"
	"fmt"
)

type HomeController struct {
	beego.Controller

}

func (c *HomeController) Home() {
	c.Data["cateogrys"] = models.Categorys
	c.Data["articles"] = models.GetArticleList(-1)
	c.Layout = "layout.html"
	c.TplName = "home.html"


}

func (c *HomeController) Article() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		fmt.Println("id error")
	}
	c.Data["article"] = models.GetArticleById(id)
	c.Data["cateogrys"] = models.Categorys
	c.Layout = "layout.html"
	c.TplName = "article.html"
}
func (c *HomeController)Categroy() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		fmt.Println("id error ")
	}
	c.Data["cateogrys"] = models.Categorys
	c.Data["articles"] = models.GetArticleList(id)
	c.Layout = "layout.html"
	c.TplName = "category.html"

}