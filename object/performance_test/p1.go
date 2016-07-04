package main

import (
	"fmt"
	"strconv"
	"time"
)

type User struct  {
	Id int
	Name string
	Age int
	Gender string

}

func main() {
	statistic(test01)
	statistic(test02)
}


func test01(){
	k1 := make(map[string]User)
	for i:=0;i<3000000;i++ {
		k1["key" + strconv.Itoa(i)] = User{i,"name"+ strconv.Itoa(i),(i+20), "男"}
	}
}

func test02(){

	k2 := make(map[string]*User)
	for i:=0;i<3000000;i++ {
		k2["key" + strconv.Itoa(i)] = &User{i,"name"+ strconv.Itoa(i),(i+20), "男"}
	}
}



func statistic(k func()) {
	t1 := time.Now().Unix()
	fmt.Println(t1)
	k()
	t2 := time.Now().Unix()
	fmt.Println(t2)
	fmt.Println("耗时:" + strconv.FormatInt(t2 - t1, 10))
}
