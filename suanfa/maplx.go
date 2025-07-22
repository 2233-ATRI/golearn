package main

import "fmt"

//type str struct {
//	name, city string
//	age        int
//}

func main() {
	// 正确的声明方式：map的map
	//var mapSlice = make(map[int]map[string]string, 3)
	//
	//// 初始化内层map
	//mapSlice[0] = make(map[string]string, 3)
	//mapSlice[0]["name"] = "王五"
	//mapSlice[0]["password"] = "123456"
	//mapSlice[0]["address"] = "红旗大街"
	//
	//// 添加第二个元素
	//mapSlice[1] = make(map[string]string)
	//mapSlice[1]["name"] = "李四"
	//mapSlice[1]["password"] = "abcdef"
	//
	//// 遍历打印
	//for index, value := range mapSlice {
	//	fmt.Printf("index:%d value:%v\n", index, value)
	//}
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
