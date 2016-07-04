package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func showField(fieldValue string) string {
	return "'" + fieldValue + "'"
}

func render(tpl string, data interface{}) string {
	//自定义template函数，可直接在模板中使用
	tplFuncs := template.FuncMap{
		"showField": showField,
	}

	t, err := template.New("tpl").Funcs(tplFuncs).Parse(tpl)

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

type User struct {
	Username string
	Password string
}

func main() {

	tpl := "insert into user(username,password) value({{.Username | showField}}, {{.Password | showField}})"
	user := &User{"qiangzi", "123456"}
	rst := render(tpl, user)
	fmt.Println(rst)
}
