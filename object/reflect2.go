package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (this Person) showName() {
	ps(this.Name)
}

func (this Person) showAge() {
	ps(this.Age)
}

func (this *Person) showInfo() {
	ps(this.Name, ",", this.Age, " years old")
}

func ps(o ...interface{}) {
	fmt.Println(o...)
}

func main() {
	var p Person
	var p2 *Person

	v1 := reflect.ValueOf(p)
	v2 := reflect.ValueOf(p2)

	ps(v1)
	ps(v2)
	ps(v1.NumMethod())
	ps(v2.NumMethod())

	p = Person{"qiangzi", 22}
	p.showInfo()
}
