package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Id int
	Name string
}
func main() {

	var a interface{}

	var b string = "hello"


	fmt.Println(a, b)


	detect(b)

	var c int = 10

	detect(c)


	k1 := reflect.ValueOf(111).Interface()
	fmt.Println(k1)
	fmt.Println(k1 == 111)

	var d int
	d = k1.(int)

	fmt.Println(d)

	u := User{}

	u2 := &User{}


	fmt.Println(reflect.ValueOf(u))
	fmt.Println(reflect.ValueOf(u2))


	tt2 := reflect.ValueOf(u2).Interface()

	vk := tt2.(*User)

	fmt.Println(vk)


	var a11 int
	var a12 int64 = 10

	a11 = int(a12)

	var a13 int64
	a13 = int64(a11)

	fmt.Println(a11, a13)

	fmt.Println(int8(a11),int16(a11),int32(a11),int64(a11))
	fmt.Println(int8(a13),int16(a13),int32(a13),int64(a13))




}

func detect(o interface{}) {
	switch o.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	}
}