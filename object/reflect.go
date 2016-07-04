package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (this Person) showName(){
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

	var a int
	var b string
	var c bool

	ps(reflect.TypeOf(a))
	ps(reflect.TypeOf(b))
	ps(reflect.TypeOf(c))

	ps(reflect.ValueOf(a))
	ps(reflect.ValueOf(b))
	ps(reflect.ValueOf(c))


	var a1 *int = &a

	a2 := reflect.ValueOf(*a1);
	a2.SetInt(10)
	ps(a2)


	var p Person
	var p2 *Person

	ps(p)
	ps(p2)
}
