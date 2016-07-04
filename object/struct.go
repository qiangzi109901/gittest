package main

import (
	"fmt"
)

//定义结构、属性有序
type Rect struct {
	x, y          int
	width, height float64
}

//为结构定义方法
func (a *Rect) Area() float64 {
	return a.width * a.height
}

func main() {

	//四种初始化的方式
	r1 := new(Rect)
	r2 := &Rect{}
	r3 := &Rect{100, 100, 20.6, 20.5}
	r4 := &Rect{width: 100, height: 200}

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
	fmt.Println(r4)

	fmt.Println(r3.Area()) //422.3
	fmt.Println(r4.Area()) //20000

	//重新构造
	r4 = new(Rect)
	fmt.Println(r4)

	r4 = nil
	fmt.Println(r4)
}
