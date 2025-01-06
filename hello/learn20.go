package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("/goword/hello", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("error is ", err)
		return
	}

	var r io.Reader
	r = tty

	var w io.Writer
	w = r.(io.Writer)
	w.Write([]byte("hello world"))
	
}
