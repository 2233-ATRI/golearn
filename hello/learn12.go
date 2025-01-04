package main

import "fmt"

func printmap(m map[string]string) { //这里传递的m是指针，会跟随修改而改变
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	citymap := make(map[string]string)
	citymap["china"] = "beijing"
	citymap["usa"] = "newnork"
	citymap["japan"] = "tokyo" //增
	for k, v := range citymap {
		fmt.Println(k, "   ", v)
	}
	fmt.Println()
	fmt.Println()
	printmap(citymap)
	delete(citymap, "china") //删除
	fmt.Println()
	citymap["usa"] = "dc" //修改

	for k, v := range citymap {
		fmt.Println(k, "   ", v)
	} //查找
	fmt.Println()
	fmt.Println()

	printmap(citymap)

}
