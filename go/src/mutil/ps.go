package mutil

import (
	"fmt"
)

func ps(args ...interface{}) {
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println("")
}
