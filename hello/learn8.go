package main

import "fmt"

func main() {
	var myarray [10]int
	myarray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 10; i++ {
		fmt.Println(myarray[i])
	}

	//for index,value :=range myarray{
	//	fmt.Println(index,value)
	//}
	for _, value := range myarray {
		fmt.Println(value)
	}
}
