package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

type Map map[string]int

func (m Map) increment(key string) {
	m[key] ++
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(" a < 2")
	} else {
		fmt.Println("a >= 2")
	}

	var m Map = make(Map)
	m["java"] = 1
	m["go"] = 2
	fmt.Println(m)
	m.increment("java")
	fmt.Println(m)


}