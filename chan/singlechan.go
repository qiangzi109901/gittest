package main

import "fmt"

func main() {
	num := make(chan int)
	squares := make(chan int)

	go counter(num)
	go squarer(num, squares)
	print(squares)
}


func counter(out chan <- int) {
	for i:=0;i<10;i++ {
		out <- i
	}
	close(out)
}

func squarer(in <-chan int, out chan <- int) {
	for x := range in{
		out <- x * x
	}
	close(out)
}

func print(in <-chan int) {
	for x := range in{
		fmt.Println(x)
	}
}