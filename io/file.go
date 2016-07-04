package main

import (
	"fmt"
	"os"
)

func testRead(path string){
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}	
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(stat)
	size := stat.Size()

	a := make([]byte, size)
	file.Read(a)
	fmt.Println(string(a))
		
}

func testRead2(path string){
	file,err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	a := make([]byte, 10)
	var b []byte
	for {
		n,err := file.Read(a)
		if err == nil {
			fmt.Println("读取字节数：", n)
			fmt.Println("读取内容：", string(a[0:n]))
			b = append(b,a[0:n]...)
		} else {
			fmt.Println("结束",n, err)
			break
		}
	}

	fmt.Println(string(b))
}



func main() {

	testRead2("../golang.readme")

}