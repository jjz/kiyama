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
	page, _ := c.GetInt(models.PAGE, 1)

	articles := models.GetArticlesBypPage(page, models.PAGE_SIZE)
	c.Data["cateogrys"] = models.Categorys
	c.Data["articles"] = articles
	c.Data["page"] = page
	c.Data["previous"] = page - 1
	c.Data["hasNext"] = len(articles) >= models.PAGE_SIZE
	c.Data["next"] = page + 1
	c.Layout = "layout.html"
	c.TplName = "home.html"

}

func (c *HomeController) Article() {
	id, _ := c.GetInt(":id", -1)
	if id == -1 {
		fmt.Println("id error")

	}
	article, err := models.GetArticleById(id)

	models.UpdateArticleView(id)
	if err != nil {
		c.Data["article"] = "没有了"
	}
	article.ToSafeHtml()
	fmt.Println(article.Html)
	c.Data["article"] = article
	c.Data["next"] = id + 1
	c.Data["previous"] = id - 1
	c.Data["cateogrys"] = models.Categorys
	c.Data["id"] = id
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