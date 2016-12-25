package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

var Categorys []*Category
var CategoryIndex int
var ArticleIndex int

type Category struct {
	Id       int
	Title    string
	Articles []*Article `orm:"reverse(many)"`
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

func GetArticleById(artcileId int) (*Article, error) {
	o := orm.NewOrm()
	article := Article{Id:artcileId}
	err := o.Read(&article)
	return &article, err
}

func AddCategory(category *Category) (error) {
	o := orm.NewOrm()
	_, err := o.Insert(category)
	return err
}

func CheckCategoryExist(title string) (bool, error) {
	o := orm.NewOrm()
	category := new(Category)
	count, err := o.QueryTable(category).Filter("Title", title).Count()
	return count > 0, err

}