package main

import (
	"fmt"
)

func main() {
	var stock = 123
	var end = "2020-2-1"
	var url = "http://www.baidu.com"
	var tarl = fmt.Sprintf(url, stock, end)
	fmt.Println(tarl)
}
