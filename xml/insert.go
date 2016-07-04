package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ModelItem struct {
	model  string
	method string
	attrs  map[string]string
	tplsql string
}

func (this *ModelItem) String() string {
	return fmt.Sprintf("{\n\tmodel : %s \n\tmethod : %s \n\ttplsql : %s \n}\n", this.model, this.method, this.tplsql)
}

func (this *ModelItem) init() {
	this.attrs = make(map[string]string)
}

func (this *ModelItem) initMethod(model, method string) {
	this.attrs = make(map[string]string)
	this.model = model
	this.method = method
}

//缓存全部sql模板
var tpls []ModelItem

func findMethod(model, method string) *ModelItem {
	for _, item := range tpls {
		if item.model == model && item.method == method {
			return &item
		}
	}
	return nil
}

func pushMethod(item *ModelItem) {
	tpls = append(tpls, *item)
}

func analysis_sql_file(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		panic("文件路径不存在")
	}
	defer file.Close()

	xmldoc := xml.NewDecoder(file)
	modelName := ""
	isEnterElement := false
	modelItem := new(ModelItem)
	modelItem.init()
	for {
		t, err := xmldoc.Token()
		if err != nil {
			break
		}

		isClose := false
		switch token := t.(type) {

		case xml.StartElement:
			name := token.Name.Local
			tname := strings.TrimSpace(name)

			//根节点
			if tname == "sql" {
				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					if attrName == "id" {
						modelName = attr.Value
					}
				}
			} else if tname != "" {
				if modelItem.method != "" {
					pushMethod(modelItem)
				}
				modelItem.initMethod(modelName, name)

				for _, attr := range token.Attr {
					attrName := attr.Name.Local
					attrValue := attr.Value
					modelItem.attrs[attrName] = attrValue
				}
				isEnterElement = true

			}

		case xml.CharData:
			c := string([]byte(token))
			tc := strings.TrimSpace(c)

			if tc != "" && isEnterElement {
				modelItem.tplsql = c
				pushMethod(modelItem)
				modelItem.init()
			}

		case xml.EndElement:
			name := token.Name.Local
			tname := strings.TrimSpace(name)

			if tname == "sql" {
				isClose = true
			} else if tname != "" {
				isEnterElement = false
			}

		default:

		}

		if isClose {
			break
		}
	}
}

func ps(args ...interface{}) {
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}

var tplFuncs template.FuncMap

func init() {
	tplFuncs = map[string]interface{}{
		"q": func(arg string) string {
			return "'" + arg + "'"
		},
	}
}

func render(tpl string, data interface{}) string {

	t, err := template.New("new").Funcs(tplFuncs).Parse(tpl)
	if err != nil {
		panic(err)
	}
	var bf bytes.Buffer
	err = t.Execute(&bf, data)
	if err != nil {
		panic(err)
	}
	return bf.String()
}

type String string

func (this *String) String() string {
	return "'" + string(*this) + "'"
}

type User struct {
	Username String
	Password String
	Age      int
	Gender   String
}

func main() {
	analysis_sql_file("file/insert.xml")
	user := &User{
		Username: "qiangzi",
		Password: "123456",
		Age:      10,
		Gender:   "男",
	}
	for _, item := range tpls {
		rst := render(item.tplsql, user)
		ps(rst)
	}

}
