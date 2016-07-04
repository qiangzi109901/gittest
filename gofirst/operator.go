package main

import "fmt"

func main() {
	
	var a,b int
	a = 10
	b = 20

	a++
	ps(a)

	// ps(a ++)
	ps(b)
	ps(a + b)

	
}

func ps(data int) {
	fmt.Println(data)
}