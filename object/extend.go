package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Println(a.Name, " start eat ...")
}

type Dog struct {
	Name string
	Animal
}

func (d *Dog) Eat() {
	fmt.Println(d.Animal)
	d.Animal.Name = d.Name
	fmt.Println(d.Name, " eat bone")
}

func main() {

	dog := &Dog{Name: "dog"}
	dog.Eat()
}
