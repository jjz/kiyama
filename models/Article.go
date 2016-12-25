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
	Markdown   string `orm:""`
	View       int `orm:"default(0)"`
	Html       template.HTML `orm:"type(text);size(150000)"`
	CreateTime time.Time `orm:"auto_now_add;type(datatime)"`
	UpdateTime time.Time `orm:"auto_now;type(datatiem)"`
	FileName   string
	Md5        string
	Category   *Category `orm:"null;rel(fk)"`
	Deleted    int8 `orm:"default(0)"`
}

func GetArticlesBypPage(page int, pageSize int) (articles [] *Article) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable(article)
	offset := GetOffset(page, pageSize)
	qs.Filter("Deleted", DELETED_NORMAL).Limit(pageSize, offset).OrderBy("-CreateTime").All(&articles)
	return
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
	baseFileName := filepath.Base(fi.Name())
	fileName := strings.Replace(baseFileName, ".md", "", 1)
	article := &Article{Id:ArticleIndex, Markdown:str, Title:fileName, FileName:baseFileName}
	article.ToSafeHtml()
	return article

}

func AddArticle(article *Article) (error) {
	o := orm.NewOrm()
	_, err := o.Insert(article)
	return err

}
func CheckArticle(fileName string) (bool, error) {
	o := orm.NewOrm()
	article := new(Article)
	count, err := o.QueryTable(article).Filter("FileName", fileName).Count()
	return count > 0, err

}

