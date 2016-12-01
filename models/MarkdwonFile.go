package models

import (
	"io/ioutil"
)

type MarkdownFile struct {
	Path         string
	SubMarkdowns []string
	DirName      string
	SubDir       [] *MarkdownFile
}

func (this *MarkdownFile)ToCategory() (*Category) {
	var articles []*Article
	for _, f := range this.SubMarkdowns {
		article := FileToMarkdown(f)
		articles = append(articles, article)
	}
	CategoryIndex++
	return &Category{Id:CategoryIndex, Name:this.DirName, Articles:articles}

}
func ReadAllMarkdown(path string, dirName string) (markdownFile *MarkdownFile, err error) {
	files, err := ioutil.ReadDir(path)
	var filePaths [] string
	var subMarkdownInfos []*MarkdownFile
	for _, f := range files {
		fileName := f.Name()
		if fileName == "" || fileName[0] == '.' {
			continue

		}
		realPath := path + "/" + fileName
		if f.IsDir() {
			subFileInfo, _ := ReadAllMarkdown(realPath, fileName)
			subMarkdownInfos = append(subMarkdownInfos, subFileInfo)
		} else {
			filePaths = append(filePaths, realPath)

		}
	}
	markdownFile = &MarkdownFile{Path:path, DirName:dirName, SubDir:subMarkdownInfos, SubMarkdowns:filePaths}
	return
}