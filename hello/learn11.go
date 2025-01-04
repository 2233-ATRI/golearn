package main

import "fmt"

func main() {

	var mymap map[string]int
	fmt.Println(mymap)
	if mymap == nil {
		fmt.Println("kong")
	}
	mymap = make(map[string]int, 3)
	//var mymap = map[string]int{"one": 1, "two": 2, "three": 3} //前key,后value
	mymap["a"] = 1
	mymap["b"] = 2
	mymap["c"] = 3
	fmt.Println(mymap)
	fmt.Println(len(mymap))
	mymap["d"] = 4
	fmt.Println(mymap)
	fmt.Println(len(mymap)) //容量参考切片相同，超出就会给之前同等大小的空间
}
