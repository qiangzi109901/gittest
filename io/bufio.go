package main

import (
	"fmt"
	"strings"
	"bufio"
)

func testStringBuffer(){
	s := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	br := bufio.NewReader(s)

	b := make([]byte, 20)
	var rst []byte
	for {
		n, err := br.Read(b)
		if n > 0 {
			rst = append(rst, b[0:n]...)
		} else {
			fmt.Println(err)
			break
		}
	}
	fmt.Println(string(rst))
}

func testFileBuffer(path string){
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	def file.Close()

	br := bufio.NewReader(file)
	item := make([]byte, 20)

	for {
		
	}




}

func main() {
	testStringBuffer()
}