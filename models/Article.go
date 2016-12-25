package models

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"kiyama/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id         int
	Title      string
	Subtitle   string
	Markdown   string
	View       int `orm:"default(0)"`
	Html       template.HTML `orm:""`
	CreateTime time.Time `orm:"auto_now_add;type(datatime)"`
	UpdateTime time.Time `orm:"auto_now:type(datatiem)"`
}

func (article *Article)ToSafeHtml() (error) {
	unsafe := blackfriday.MarkdownCommon([]byte(article.Markdown))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	article.Html = template.HTML(string(html))
	return nil

}
func FileToMarkdown(filePath string) (*Article) {
	ArticleIndex++
	str, _ := utils.ReadFile(filePath)
	fi, _ := os.Open(filePath)
	fileName := filepath.Base(fi.Name())
	fileName = strings.Replace(fileName, ".md", "", 1)
	article := &Article{Id:ArticleIndex, Markdown:str, Title:fileName}
	article.ToSafeHtml()
	return article

}

func AddArticle(article *Article) (error) {
	o := orm.NewOrm()
	_, err := o.Insert(article)
	return err

}