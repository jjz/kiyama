package models

import (
	"github.com/russross/blackfriday"
	"html/template"
	"kiyama/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
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
	unsafe := blackfriday.MarkdownBasic([]byte(article.Markdown))
	//html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	article.Html = template.HTML(string(unsafe))
	fmt.Println(article.Markdown)

	return nil
}
func FileToMarkdown(filePath string) (*Article) {
	ArticleIndex++
	str, _ := utils.ReadFile(filePath)
	fi, _ := os.Open(filePath)
	baseFileName := filepath.Base(fi.Name())
	fileName := strings.Replace(baseFileName, ".md", "", 1)
	article := &Article{Id:ArticleIndex, Markdown:str, Title:fileName, FilePath:filePath}
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
	count, err := o.QueryTable(article).Filter("FilePath", fileName).Count()
	return count > 0, err

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
	fmt.Println(md5Str)
	o := orm.NewOrm()
	article := Article{FilePath:filePath}

	err := o.Read(&article, "FilePath")
	fmt.Println(article.Md5)
	if err != nil {
		fmt.Println(err.Error())
		return err

	}
	fmt.Println(md5Str, article.Md5)
	if md5Str != article.Md5 {
		article.Markdown = markdown
		article.Md5 = md5Str

		o.Update(&article, "Markdown", "Md5")
		fmt.Println("update:", article.Md5)
	}
	return nil

}


