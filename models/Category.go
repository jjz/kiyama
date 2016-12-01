package models

import (
	"github.com/astaxie/beego"
	"kiyama/utils"
	"fmt"
)

var Categorys []*Category
var CategoryIndex int
var ArticleIndex int

var AllArticles []*Article

type Category struct {
	Id       int
	Name     string
	Articles []*Article
}

func init() {
	InitMarkdown()
}

func InitMarkdown() {
	CategoryIndex = 0
	ArticleIndex = 0
	Categorys = []*Category{}
	AllArticles = []*Article{}

	markdownPath := utils.MergePath(beego.AppConfig.String("markdown_dir"))
	fileInfo, err := ReadAllMarkdown(markdownPath, "默认")
	category := fileInfo.ToCategory()
	Categorys = append(Categorys, category)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for _, f := range fileInfo.SubDir {
		category = f.ToCategory()
		Categorys = append(Categorys, category)
	}
	for _, c := range Categorys {
		for _, a := range c.Articles {
			AllArticles = append(AllArticles, a)
		}
	}

}

func GetArticleList(categoryId int) ([]*Article) {

	var articles []*Article
	if categoryId == -1 {
		for _, c := range Categorys {

			for _, a := range c.Articles {
				articles = append(articles, a)
			}

		}

	} else {
		for _, c := range Categorys {
			if c.Id == categoryId {
				for _, a := range c.Articles {
					articles = append(articles, a)
				}
				break
			}
		}

	}
	fmt.Println("articles", articles)
	return articles

}

func GetArticleById(artcileId int) (*Article) {
	for _, c := range AllArticles {
		if c.Id == artcileId {
			return c
		}
	}
	return nil
}
