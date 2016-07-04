package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main() {

	var w io.Writer

	w = os.Stdout

	w = new(bytes.Buffer)


	fmt.Println(w)


	// w = time.Second
}