package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

func testReadXml() {
	file, err := os.Open("file/user.xml")
	if err != nil {
		panic("文件路径不存在")
	}
	defer file.Close()

	xmldoc := xml.NewDecoder(file)

	for {
		t, err := xmldoc.Token()
		if err != nil {
			break
		}
		switch token := t.(type) {

		case xml.StartElement:
			name := token.Name.Local
			fmt.Println(name, "START")

		case xml.CharData:
			c := string([]byte(token))
			fmt.Println("*********", strings.TrimSpace(c) == "")

		case xml.EndElement:
			fmt.Println(token.Name.Local, "END")
		default:

		}
	}
}

func main() {
	testReadXml()
}
