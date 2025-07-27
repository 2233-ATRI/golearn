package main

import "fmt"

//func main() {
//	var wahtever [5]struct{}
//	for i := range wahtever {
//		defer fmt.Println(i)
//	}
//	for i := range wahtever {
//		defer func() { fmt.Println(i) }()
//	}
//	return
//}

//type Test struct {
//	name string
//}
//
//func (t *Test) Close() {
//	fmt.Println(t.name, " closed")
//}
//func Close(t Test) {
//	t.Close()
//}
//func main() {
//	ts := []Test{{"a"}, {"b"}, {"c"}}
//	for _, t := range ts {
//		defer Close(t)
//	}
//}//得到cba

type test22 struct {
	name string
}

func (t *test22) close() {
	fmt.Println(t.name, " closed")
}
func main() {
	ts := []test22{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.close()
	}
} //闭包导致defer最后只得到最后一次迭代的数值
