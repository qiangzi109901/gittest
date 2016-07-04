package main

import (
	"fmt"
	"math"
)


func main() {
	x,y := swap(100,200)
	fmt.Println(x)
	fmt.Println(y)



	var a int = 10
	var b int = 20

	fmt.Printf("before swap:a is %d\n", a)
	fmt.Printf("after swap:b is %d\n", b)

	swapTwo(a, b)

	fmt.Printf("after swapTwo:a is %d\n", a)
	fmt.Printf("after swapTwo:b is %d\n", b)


	swapTwo2(&a, &b)


	fmt.Printf("after swapTwo2:a is %d\n", a)
	fmt.Printf("after swapTwo2:b is %d\n", b)


	mk := func(num float64) float64 {
		return math.Sqrt(num)
	}

	fmt.Println(mk(64))


	fmt.Println("test closure programmer")

	nextNumber := getSequence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber2 := getSequence()
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())
	fmt.Println(nextNumber2())


	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(sum2([]int{1,2,3,4,5,6,7,8,9,10}))
	logsum(1,2,3,4,5,6,7,8,9,10)

}

func swap(x,y int)(int,int){
	return y,x
}



func getSequence() func() int{
	// var i int = 0
	i := 0
	return func() int{
		i ++
		return i
	}
}



func swapTwo(x,y int){
	var temp int
	temp = x
	x = y
	y = temp
}


func swapTwo2(x *int, y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}





//可变参数
func sum(nums ...int) int{
	var sums int = 0
	for _, v := range nums{
		sums += v
	}
	return sums
}

//传数组
func sum2(nums []int) int{
	var sums int = 0
	for _, v := range nums{
		sums += v
	}
	return sums
}

//可变参数传递给其它函数
func logsum(nums ...int){
	total := sum(nums...)
	fmt.Print("求和:")
	for i, v := range nums {
		fmt.Print(v)
		if i<len(nums) - 1 {
			fmt.Print(" + ")
		} else{
			fmt.Print(" = ")
		}
	}
	fmt.Println(total)
}