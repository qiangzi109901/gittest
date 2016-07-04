package main


import (
	"fmt"
)


func testInit(){
	var a []int
	printSlice(a)

	b := make([]int, 10)
	printSlice(b)

	c := make([]int, 10, 20)
	printSlice(c)

	c[9] = 10 //ok
	// c[10] = 11 //err index out of range
	
	d := []int{1,2,3,4,5,6,7,8,9,10}
	e := d[1:3]
	print("e")
	printSlice(e)

	f := d[2:5:8]   //2 表示数组d的开始索引， 5 表示结束索引  8 表示f在d中的最大索引
	print("f")
	printSlice(f)



	ks := make([]int,0)

	ks = append(ks,1)
	ks = append(ks,2)
	ks = append(ks,3)
	ks = append(ks,4)

	for i,v := range ks{
		ks = append(ks,(i+100))
		fmt.Println(v)
	}



}

func testAppend(){
	var a []int
	printSlice(a)
	
	for i := 0; i<10; i++ {
		a = append(a, i)
		printSlice(a)
		printAddress(&a)
		print("")
	}
}

func print(a interface{}){
	fmt.Println(a)
}

func printSlice(a []int){
	fmt.Printf("长度：%d,容量: %d\n", len(a), cap(a))
}

func printStringSlice(a []string){
	fmt.Printf("长度：%d,容量: %d\n", len(a), cap(a))
}

func printAddress(s... *[]int){
	for _,v := range s{
		fmt.Println("address:", &v)
	}
}

//清空切片中的空字符串
func nonempty(strs []string) []string{
	i := 0
	for _,s := range strs{
		if s != "" {
			strs[i] = s
			i ++
		}
	}
	return strs[:i]
}

//测试nonempty
func testNonempty(){
	str := []string{"hello", "", "world"}
	str = nonempty(str)
	printStringSlice(str)
}

//删除切片中的某个索引位置的值
func remove(a []int, index int) []int{
	b := a[:index]
	c := a[index+1:]
	return append(b, c...)
}


func testRemove(){
	a := []int{1,2,3,4,5,6,7,8,9}
	b := remove(a, 3)
	fmt.Println(b)
}



func testCopy(){
	a := []int{1,2,3,4,5,6,7,8}
	//将a[4:]复制到a[3:]，相当于把i>3的值向前移动一位
	copy(a[3:],a[4:]) 
	//1 2 3 4 5 6 7 8
	//1 2 3 5 6 7 8 8
	fmt.Println(a)
}


func testAssign(){
	fmt.Println("******")
	a := make([]int, 5, 10)

	a[0] = 10
	a = append(a, 20)

	fmt.Println(a)
}



func main() {
	testInit()
	testAppend()
	testNonempty()
	testRemove()
	testCopy()
	testAssign()
}