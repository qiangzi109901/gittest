package main

import (
	"fmt"
	"strings"
	"reflect"
)

func main() {
	//定义string
	var a string
	b := "hello"
	fmt.Println(a, b)


	//计算长度
	fmt.Println(len(b))

	//获取某个索引下的字符串
	fmt.Println(string(b[2]))

	//遍历字符串
	for i := range b{
		fmt.Println(string(b[i]))
	}

	//大小写转换
	fmt.Println(strings.ToUpper("golang"))
	fmt.Println(strings.ToLower("GOLANG"))


	//查找字符串/字节
	fmt.Println(strings.Index("hello golang", "o")) //3
	fmt.Println(strings.Index("hello golang", "go")) //6
	fmt.Println(strings.IndexByte("hello 123456", byte(52)))
	//从后查找
	fmt.Println(strings.LastIndex("hello golang", "o")) //7


	//包含
	fmt.Println(strings.Contains("hello golang", "go"))
	fmt.Println(strings.ContainsAny("hello golang", "123o")) //任意包含
	fmt.Println(strings.ContainsRune("hello golang", 'o'))

	//比较
	fmt.Println("hello" == "hello")
	fmt.Println(strings.Compare("hello", "hello")) //0

	//前后缀
	fmt.Println(strings.HasPrefix("music_001.mp3", "music"))
	fmt.Println(strings.HasSuffix("music_002.mp3", "mp3"))

	//分隔
	d := strings.Split("my name is qiangzi", " ")
	fmt.Println(d)//[]string


	//连接
	fmt.Println(strings.Join(d, ",")) //my,name,is,qiangzi

	//重复
	fmt.Println(strings.Repeat("hello ", 2))//"hello hello "

	//替换
	fmt.Println(strings.Replace("welcome to golang world", "o", "O", 1))  //welcOme to golang world
	fmt.Println(strings.Replace("welcome to golang world", "o", "O", 2))  //welcOme tO golang world
	fmt.Println(strings.Replace("welcome to golang world", "o", "O", -1)) //welcOme tO gOlang wOrld

	//处理空白字符
	fmt.Println(strings.Trim(" hello ", " ")) //"hello"
	fmt.Println(strings.Trim("fa hello golang \r\n\n  ", " \nfa")) //"hello golang \r"
	fmt.Println(len(strings.TrimSpace(" welcome \r\n "))) //"welcome", 会清空 '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0

	fmt.Println(strings.TrimLeft(" hello \n", "\n ")) //"hello \n"
	fmt.Println(strings.TrimRight(" hello \n", "\n ")) //" hello"

	//map
	rot := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r - 'A' + 13) % 26
		case r >= 'a' && r <= 'z':
			return 'a' + (r -'a' + 13) % 26
		}
		return r
	}
	fmt.Println(strings.Map(rot, "A B y z")) //N O l m
	fmt.Println("11")


	//其它
	fmt.Println(strings.Title("hello")) //返回Unicode处理的hello
	fmt.Println(strings.TrimSuffix("hello.mp4",".mp4"))
	fmt.Println(strings.TrimPrefix("music_hello.mp4","music_"))

	fmt.Println(reflect.TypeOf("233445"[:3]))


	fmt.Println(getFiledName("user_id"))
}



func getFiledName(columnName string) string {
	cs := strings.Split(columnName, "_")
	fmt.Println(cs)
	for i,v := range cs {
		cs[i] = strings.ToUpper(string(v[0])) + strings.ToLower(v[1:])
	}
	return strings.Join(cs,"")

}