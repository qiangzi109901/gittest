
package main

import "fmt"

func main() {
	var a int = 10

	if(a < 20){
		fmt.Println("a小于20")
	} else {
		fmt.Println("a大于或等于20")
	}



	if(a < 0){
		fmt.Println("a小于0")
	} else if(a<10){
		fmt.Println("a大于0，小于10")
	} else if(a<20){
		fmt.Println("a小于20")
	} else{
		fmt.Println("a大于20")
	}


	var b int = 15

	switch b{
		case 10,15: fmt.Println("10 or 15")
		case 20: fmt.Println("20")
		case 30: fmt.Println("30")
		default: fmt.Println("nothing")
	}

	b = 50

	switch{
		case b == 10,b == 50:
			fmt.Println("10 or 50")
		case b == 20:
			fmt.Println(20)
		case b == 30:
			fmt.Println(30)
		default:
			fmt.Println("nothing")
	}

	

	var c int = 10


	// select {
	// 	case c == 10:
	// 		fmt.Println("c is 10 or 100")
	// 	case c == 20:
	// 		fmt.Println("c is 20")
	// 	case c == 30:
	// 		fmt.Println("c is 30")
	// }	
	




}