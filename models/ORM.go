package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"kiyama/utils"
	"fmt"
)

const (
	PAGE = "page"
	PAGE_SIZE = 20
)

func init() {
	mysqlStr := beego.AppConfig.String("mysql")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlStr)
	orm.RegisterModel(new(Article), new(Category))
	orm.RunSyncdb("default", false, true)
	RefreshMarkdown()

}

func RefreshMarkdown() {
	CategoryIndex = 0
	ArticleIndex = 0
	Categorys = []*Category{}

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
		exist, err := CheckCategoryExist(category.Title)
		if err != nil {
			fmt.Println(err.Error())

		}
		if !exist {
			AddCategory(category)
		}
	}
	//todo 遍历所有已经存在的文章，不存在的需要删除
	for _, c := range Categorys {
		for _, a := range c.Articles {

			exist, err := CheckArticle(a.FilePath)
			if err != nil {
				fmt.Println(err.Error())
			}
			if !exist {

				err := AddArticle(a)
				if err != nil {
					fmt.Println(err.Error())
				}

			} else {
				UpdateArticle(a.FilePath, a.Markdown)

			}
		}
	}

}
