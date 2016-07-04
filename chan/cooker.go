package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//c1()
	c3()
	//c4_2()
	//c3()
	//c4_2()
}





func getRandom() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

//一个厨师一个服务员一盘菜
func c1(){

	food := make(chan int)

	//厨师做菜
	go func(){
		k := getRandom()
		fmt.Println("厨师出菜", k)
		food <- k
	}()

	//服务员传菜
	mf := <- food
	fmt.Println("服务员传菜:", mf)

}


//多个厨师一个服务员一盘菜

func c2(){

	food := make(chan int)

	//厨师做菜
	go func() {
		for i:=0;i<5;i++ {
			k:= getRandom()
			fmt.Printf("厨师%d出菜:%d\n",i,k)
			food <- k
		}
	}()

	//服务员传菜
	mf := <- food
	fmt.Println("服务员传菜", mf)
}


//一个厨师多个服务员一盘菜
func c3(){

	food := make(chan int)

	//厨师做菜
	go func(){
		k:= getRandom()
		fmt.Println("厨师出菜:",k)
		food <- k
	}()

	//服务员传菜
	for i:=0;i<5;i++{
		mf := <- food
		fmt.Println("服务员传菜", mf)
	}
}

//多个厨师多个服务员一盘菜
func c4(){


	food := make(chan int)

	//厨师做菜
	go func() {
		for i:=0;i<5;i++ {
			k:= getRandom()
			fmt.Printf("厨师出菜:%d\n",k)
			food <- k
		}
	}()

	//服务员传菜
	for i:=0;i<5;i++{
		mf := <- food
		fmt.Println("服务员传菜", mf)
	}

}


func c4_2(){


	food := make(chan int)

	//厨师做菜
	go func() {
		for i:=0;i<5;i++ {
			k:= getRandom()
			fmt.Println("厨师出菜:",k)
			food <- k
		}
		close(food)
	}()

	//服务员传菜
	for mf := range food{
		fmt.Println("服务员传菜:", mf)
	}

}

//一个厨师一个服务员多盘菜
func c5(){

}