package main

import "fmt"

func main() {
	

	var a int = 10

	var b int

	for b = 0; b<10; b++ {

		if b < 3 {
			continue
		}
		fmt.Printf("value of b:%d\n", b)

		if b > 5 {
			break
		}
	}
	fmt.Println("")

	b = 20
	for a < b {
		fmt.Printf("value of a:%d\n", a)
		a ++
	}
	fmt.Println(b)
	fmt.Println(a)



	nums := [6]int{1,2,3,4}

	for i,x := range nums {
		fmt.Printf("nums index : %d ,value is : %d \n", i, x)
	}


	var c int = 0
	var d int = 10
	LOOP: for c < d {
		if c == 5 {
			c++
			goto LOOP
		}
		fmt.Printf("value of c is : %d \n", c)
		c ++
	}


	var e = 0

	for true {
		if e == 100 {
			fmt.Println("e is 100")
			break
		}
		e ++
	}


	for {
		fmt.Println("退出")
		break;
	}


	fmt.Println(fo1())

}


func fo1() int{

	for i,_ := range make([]int, 10) {
		if i == 5 {
			return i
		}
	}
	return 100
}