package main

import (
	"fmt"
)

type USB interface {
	Name() string
	Functional
}

type Functional interface {
	Function()
}

type AndroidUSB struct {
	name string
}

func (this AndroidUSB) Name() string {
	return this.name
}

func (this AndroidUSB) Function() {
	fmt.Println(this.Name(), "用于将Android手机连接到电脑")
}

type IphoneUSB struct {
	name string
}

func (this IphoneUSB) Name() string {
	return this.name
}

func (this IphoneUSB) Function() {
	fmt.Println(this.Name(), "用于将Iphone手机连接到电脑")
}

func getType(usb USB) {
	switch usb.(type) {
	case AndroidUSB:
		fmt.Println("Android USB Type")
	case IphoneUSB:
		fmt.Println("Iphone USB Type")
	default:
		fmt.Println("Other Type")
	}
}

func getType2(o interface{}) {
	switch o.(type) {
	case AndroidUSB:
		fmt.Println("Android USB Type")
	case IphoneUSB:
		fmt.Println("Iphone USB Type")
	default:
		fmt.Println("Other Type")
	}
}

func main() {
	var usb USB
	usb = AndroidUSB{"Android USB"}
	fmt.Println(usb.Name())
	usb.Function()

	u2 := USB(AndroidUSB{"小米 USB"})
	fmt.Println(u2.Name())
	u2.Function()

	getType(AndroidUSB{})
	getType(IphoneUSB{})

	var a int
	getType2(AndroidUSB{})
	getType2(IphoneUSB{})
	getType2(a)

}
