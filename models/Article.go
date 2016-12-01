package models

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"kiyama/utils"
	"os"
	"path/filepath"
	"strings"
)

type Article struct {
	Id       int
	Time     string
	Name     string
	Html     template.HTML
	Markdown string
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
	article := &Article{Id:ArticleIndex, Time:"", Markdown:str, Name:fileName}
	article.ToSafeHtml()
	return article

}