package main

import (
	"time"
	"fmt"
	"strings"
	"errors"
)

func main() {

	//test01()
	//test02()

	//statistic(test01)
	//statistic(test03)
	//e := test04()
	//fmt.Println(e)
	ts,e := test05()
	fmt.Println(ts)
	fmt.Println(e)
}

//顺序执行
func test01(){
	filenames := [...]string{"1.jpg","2.jpg","3.jpg","4.jpg","5.jpg","6.jpg","7.jpg","8.jpg","9.jpg","10.jpg"}

	thumbs := make([]string, len(filenames))
	for i,v := range filenames {
		f,err := ImageFile(v)
		if err != nil {
			fmt.Println(err)
		}
		thumbs[i] = f
		fmt.Println(f)
	}
	fmt.Println(thumbs)
}

//goruntine执行,不考虑结果

//结果并没有任何输出,因为main线程执行完毕后,goruntine会放弃执行了
func test02() {
	filenames := [...]string{"1.jpg","2.jpg","3.jpg","4.jpg","5.jpg","6.jpg","7.jpg","8.jpg","9.jpg","10.jpg"}

	for _,v := range filenames {
		go func(file string){
			f,err := ImageFile(file)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(f)
		}(v)
	}
}

//goruntine执行,通过chan约束
func test03() {
	filenames := [...]string{"1.jpg","2.jpg","3.jpg","4.jpg","5.jpg","6.jpg","7.jpg","8.jpg","9.jpg","10.jpg"}

	c := make(chan int)
	for _,v := range filenames {
		go func(file string){
			f,err := ImageFile(file)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(f)
			c <- 1
		}(v)
	}

	for range filenames {
		<- c
	}
	close(c)
}

//goroutine执行,并返回是否有失败的
func test04() error{
	filenames := [...]string{"1.jpg","2.jpg","3.jpg","4.jpg","5.jpg","6.jpg","7.jpg","8.jpg","9.jpg","10.jpg"}

	c := make(chan error)

	for _,v := range filenames {

		go func(file string){
			_,err := ImageFile(file)
			c <- err

		}(v)
	}

	var hasError error
	for range filenames {
		e := <- c
		if e != nil {
			hasError = e
			//close(c) //通知
			//break
		}
	}
	return hasError
}


type FileItem struct{
	Filename string
	Err error
}

//goroutine执行,并返回全部的处理过的文件
func test05() (thumbs[]string,hasErr error) {
	filenames := [...]string{"1.jpg","2.jpg","3.jpg","4.jpg","5.jpg","6.jpg","7.jpg","8.jpg","9.jpg","10.jpg"}
	c := make(chan FileItem)
	for _,v := range filenames {
		go func(file string){
			f,err := ImageFile(file)
			c <- FileItem{f,err}
		}(v)
	}

	var hasError error
	for range filenames {
		item := <- c
		if item.Err != nil {
			hasError = item.Err
			//close(c) //通知
			//break
		}
		thumbs = append(thumbs,item.Filename)
	}
	return thumbs,hasError
}








//模拟对图片进行缩略图处理
func ImageFile(infile string) (outfile string,err error) {

	time.Sleep(1000 * time.Millisecond)

	if strings.Contains(infile, "5") {
		return "",errors.New("5 is error")
	}
	return "thumb_" + infile , nil
}


func statistic(f func()){
	t1 := time.Now().Unix()

	f()

	fmt.Println("耗时:", time.Now().Unix() - t1)
}