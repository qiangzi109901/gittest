package main

import (
	"fmt"
)

func typeof(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("other")
	}
}

type String string

func main() {
	var a int = 10

	fmt.Printf("address of variable a:%x\n", &a)

	var p *int

	p = &a

	fmt.Printf("address of point p:%x\n", p)
	fmt.Printf("value of point p:%d\n", *p)

	var k *string

	var u string = "hello"
	k = &u

	fmt.Println(k)
	fmt.Println(*k)
	fmt.Println(string(*k))

	// var k2 String = "eeee"
	typeof(string(*k))
}
