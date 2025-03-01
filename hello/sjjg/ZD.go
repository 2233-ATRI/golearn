// //package main
// //
// //import "fmt"
// //
// //func main() {
// //	m := make(map[string]int, 4)
// //	m["a"] = 1
// //	m["b"] = 2
// //	m["c"] = 3
// //	m["d"] = 4
// //	fmt.Println(m)
// //
// //}
//
// package main
//
// import "fmt"
//
//	func main() {
//		// 新建一个容量为4的字典 map
//		m := make(map[string]int64, 4)
//		// 放三个键值对
//		m["dog"] = 1
//		m["cat"] = 2
//		m["hen"] = 3
//		fmt.Println(m)
//		// 查找 hen
//		which := "hen"
//		v, ok := m[which]
//		if ok {
//			// 找到了
//			fmt.Println("find:", which, "value:", v)
//		} else {
//			// 找不到
//			fmt.Println("not find:", which)
//		}
//		// 查找 ccc
//		which = "ccc"
//		v, ok = m[which]
//		if ok {
//			// 找到了
//			fmt.Println("find:", which, "value:", v)
//		} else {
//			// 找不到
//			fmt.Println("not find:", which)
//		}
//	}
package main

import "sync"

type SET struct {
	m    map[int]struct{}
	len  int
	lock sync.Mutex
}

func NEWset(cap int) *SET {
	temp := make(map[int]struct{}, cap)
	return &SET{
		m: temp,
	}
}

func (s *SET) Add(elem int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[elem] = struct{}{}
	//s.len++
	s.len = len(s.m)
}

func (s *SET) Remove(elem int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		return
	}
	delete(s.m, elem)
	s.len--
}

func (s *SET) has(item int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.m[item]
	return ok
}
