package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

type Base struct {
	a int
}

func (b *Base) setB(i int) {
	b.a += i
}

type Derive struct {
	Base
	d int
}

func (d *Derive) setD(i int) {
	d.d += i
}

func testBaseAndDerive() {

	var d Derive
	fmt.Println(d.a, d.d)
	d.setB(10)
	d.setD(20)
	fmt.Println(d.a, d.d)

	//总结：在struct内嵌其它结构，就相当于在其内嵌入其它结构的属性和方法,Derive相当于如下所示
	/**
	 type Derive struct {
	 	a int
	 	d int
	 }
	 func (d *Derive) setB(i int) {
		d.a += i
	 }
	 func (d *Derive) setD(i int) {
		d.d += i
	 }
	*/

}

type A struct {
	name string
}

func main() {
	var k []A
	a1 := &A{"A"}
	k = append(k, *a1)

	fmt.Println(k)
}
