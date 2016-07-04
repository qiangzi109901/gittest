package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

type User struct {
	Username string
}

func tpl_data_with_struct() {
	t := template.New("tpl_data_with_struct")
	t, _ = t.Parse("hello, {{.Username}} !\n")
	user := User{Username: "AAA"}
	var bf bytes.Buffer
	t.Execute(&bf, user)
	fmt.Println(string(bf.Bytes()))
}

func tpl_data_with_map() {
	t := template.New("tpl_data_with_map")
	t, _ = t.Parse("hello, {{.username}} !\n")
	userMap := make(map[string]string)
	userMap["username"] = "BBB"
	t.Execute(os.Stdout, userMap)
}

func main() {
	tpl_data_with_struct()
	tpl_data_with_map()
}
