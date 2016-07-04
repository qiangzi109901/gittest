package main

import "fmt"

var i int

func init() {
	i = 10
}

func main() {

	fmt.Println(i)

	i := 20

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(i)
}
