package main

import (
	_ "kiyama/routers"

	"path/filepath"
	"strings"
	"html/template"
	"fmt"
	"github.com/yosssi/ace"

	"github.com/astaxie/beego"
	"github.com/Joker/jade"
	"kiyama/utils"
	"os"
	"kiyama/models"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.Debug = true
	initArgs()
	beego.AddFuncMap("TimeToString", models.TimeToString)
	beego.Run()
}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-initmd" {
			models.RefreshMarkdown()
			os.Exit(0)
		}
	}

}

func iniTemplate() {
	beego.AddTemplateEngine("ace", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		aceOptions := &ace.Options{DynamicReload: true, FuncMap: funcs}
		aceBasePath := filepath.Join(root, "base")
		aceInnerPath := filepath.Join(root, strings.TrimSuffix(path, ".ace"))

		tpl, err := ace.Load(aceBasePath, aceInnerPath, aceOptions)

		if err != nil {
			return nil, fmt.Errorf("error loading ace template: %v", err)
		}

		return tpl, nil
	})
	beego.AddTemplateEngine("jade", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		jadePath := filepath.Join(root, path)
		content, err := utils.ReadFile(jadePath)
		fmt.Println(content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		tpl, err := jade.Parse("name_of_tpl", content)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		tmp := template.New("Person template")
		tmp, err = tmp.Parse(tpl)
		if err != nil {
			return nil, fmt.Errorf("error loading jade template: %v", err)
		}
		fmt.Println(tmp)
		return tmp, err

	})
}

