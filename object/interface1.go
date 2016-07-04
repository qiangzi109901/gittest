package main

import (
	"fmt"
)


type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}


type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

type MInteger int

func (a MInteger) Less(b MInteger) bool {
	return a < b
}

type MLesser interface {
	Less(b MInteger) bool
}


func main() {
	var a Integer = 1

	var b LessAdder= &a

	fmt.Println(b.Less(2))

	b.Add(3)

	fmt.Println(b)


	var c MInteger = 2

	var d MLesser = c

	fmt.Println(d.Less(1))


}