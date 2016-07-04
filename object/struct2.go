package main


import (
	"fmt"
	"bytes"
	"strconv"
)

type Person struct {

	Name string
	Age int

}

func (this Person) show()  {
	fmt.Println(this.Name , ",今年 ", this.Age, '岁')
}

func (this Person) String() string {
	bf := bytes.NewBufferString(this.Name)
	bf.WriteString(",今年")
	bf.WriteString(strconv.Itoa(this.Age))
	bf.WriteString("岁")
	return bf.String()
}

func main() {

	p1 := Person{"张三", 20}
	fmt.Println(p1)



}
