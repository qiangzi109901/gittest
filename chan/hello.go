package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	c := make(chan int, 1)
	go func(){
		c <- 1
		fmt.Println("第一次发送成功")
		c <- 2
		fmt.Println("第二次发送成功")
		c <- 3
		fmt.Println("第三次发送成功")
		c <- 4
		fmt.Println("第四次发送成功")
	}()

 	fmt.Println("准备接收")
	k := <- c
	fmt.Println("第一次接收", k)
	k = <- c
	fmt.Println("第二次接收", k)
	k = <- c
	fmt.Println("第三次接收", k)
	k = <- c
	fmt.Println("第四次接收", k)
}


func testc01(){
	c := make(chan int)

	go func() {
		fmt.Println("goroutine ok")
		c <- 1
		close(c)
	}()
	<-c
	fmt.Println("main ok")
}

func testc02(){
	c := make(chan int)

	go func() {

		for i:=0;i<10;i++{
			fmt.Println("goroutine ok")
			c <- i
		}
	}()
	for range [10]int{} {
		<-c
		fmt.Println("main ok")
	}

}

func testc03(){

	c := make(chan int)

	go func(){
		c <- 10
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}

func testc04(){
	c := make(chan []int)

	go func(){
		c <- makenums(1)
		//close(c)
	}()


	for v := range c{
		fmt.Println(v)
		for _,num := range v {
			fmt.Println(num)
			//c <- makenums(num)
		}
	}
}

func makenums(n int) []int{
	if n > 100 {
		return nil
	}
	rst := make([]int,10)
	for i:=0;i<10;i++{
		rand.Seed(time.Now().UnixNano())
		jns := n * 10
		rst[i] = rand.Intn(jns) + 1
	}
	return rst
}