package models

import (
	"html/template"
	"kiyama/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/golang-commonmark/markdown"
)

type Article struct {
	Id         int
	Title      string
	Subtitle   string
	Markdown   string `orm:"type(text);size(150000)"`
	View       int `orm:"default(0)"`
	Html       template.HTML `orm:"-"`
	CreateTime time.Time `orm:"auto_now_add;type(datatime)"`
	UpdateTime time.Time `orm:"auto_now;type(datatiem)"`
	FilePath   string
	Md5        string
	Category   *Category `orm:"null;rel(fk)"`
	Status     int8 `orm:"default(0)"`
}

func GetArticlesBypPage(page int, pageSize int) (articles [] *Article) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable(article)
	offset := GetOffset(page, pageSize)
	qs.Filter("Status", STATUS_ENABLE).Limit(pageSize, offset).OrderBy("-UpdateTime").All(&articles)
	return
}

func (article *Article)ToSafeHtml() (error) {
	//unsafe := blackfriday.MarkdownBasic([]byte(article.Markdown))
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	md := markdown.New(markdown.XHTMLOutput(true), markdown.Nofollow(true))
	article.Html = template.HTML(md.RenderToString([]byte(article.Markdown)))
	fmt.Println(article.Markdown)

	return nil
}
func FileToMarkdown(filePath string) (*Article) {
	ArticleIndex++
	str, _ := utils.ReadFile(filePath)
	fi, _ := os.Open(filePath)
	baseFileName := filepath.Base(fi.Name())
	fileName := strings.Replace(baseFileName, ".md", "", 1)
	fmt.Println("title:", fileName)
	article := &Article{Markdown:str, Title:fileName, FilePath:filePath}
	return article

}

func AddArticle(article *Article) (error) {
	o := orm.NewOrm()
	_, err := o.Insert(article)
	return err

}
func CheckArticleWithPath(filePath string) (article *Article, err error) {
	o := orm.NewOrm()
	var articles [] *Article
	_, err = o.QueryTable(article).Filter("FilePath", filePath).All(&articles)
	if err != nil {
		return
	}
	if len(articles) > 0 {
		article = articles[0]
	}
	return

}
func UpdateArticleView(articleId int) (error) {
	o := orm.NewOrm()
	article := Article{Id:articleId}
	o.Read(&article)
	article.View = article.View + 1
	_, err := o.Update(&article, "View")
	return err
}
func UpdateArticle(filePath string, markdown string) (error) {
	md5Str := utils.GetMd5FromFile(filePath)
	o := orm.NewOrm()
	article := Article{FilePath:filePath}
	err := o.Read(&article, "FilePath")
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	if md5Str != article.Md5 {
		article.Markdown = markdown
		article.Md5 = md5Str
		article.UpdateTime = time.Now()
		article.Status = STATUS_ENABLE
		o.Update(&article, "Markdown", "Md5", "UpdateTime", "Status")


	} else {
		article.Status = STATUS_ENABLE
		o.Update(&article, "Status")
		fmt.Println("update ,Status:", article.Md5,article.Status)
	}
	return nil
}

func UpdateAllArticleDeleted() error {
	o := orm.NewOrm()
	o.Raw("UPDATE `article` SET `status`=? WHERE `status`=?", STATUS_DISABLE, STATUS_ENABLE).Exec()
	return nil

}
func UpdateArticleUndeleted(articleId int) error {
	o := orm.NewOrm()
	article := Article{Id:articleId, Status:STATUS_ENABLE}
	_, err := o.Update(&article, "Status")
	return err
}


