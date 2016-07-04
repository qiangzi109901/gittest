package main

import (
	"fmt"
)

func main() {

	num := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			fmt.Println(x, "num")
			num <- x
		}
		close(num)
	}()

	go func() {
		for x := range num {
			squares <- x * x
		}
		close(squares)
	}()

	for {
		x, ok := <-squares
		if ok {
			fmt.Println(x)
		} else {
			break
		}
	}

}
