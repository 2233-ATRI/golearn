package main

import "sync"

//func main() {
//	// 创建一个容量为2的切片
//	array := make([]int, 0, 2)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	// 虽然 append 但是没有赋予原来的变量 array
//	_ = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	_ = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	_ = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	fmt.Println("-------")
//	// 赋予回原来的变量
//	array = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	array = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	array = append(array, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	array = append(array, 1, 1, 1, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1, 1)
//	fmt.Println("cap", cap(array), "len", len(array), "array:", array)
//}

type Array struct {
	array []int
	len   int
	cap   int
	lock  sync.Mutex
}

func newArray(len, cap int) *Array {
	if len < 0 {
		return nil
	}

	// 使用正确的顺序创建切片
	aaa := make([]int, len, cap)

	// 创建一个新的 Array 实例
	s := new(Array)
	s.len = len
	s.cap = cap
	s.array = aaa

	return s
}

func (a *Array) APPend() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.len == a.cap {
		newcap := 2 * a.len
		if a.cap == 0 {
			newcap = 1
		}
		newarray := make([]int, newcap, newcap)
		for k, v := range a.array {
			newarray[k] = v
		}
		a.array = newarray
		a.len = a.cap + 1
		a.cap = newcap

	}
	return a.len
}
