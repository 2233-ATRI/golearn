// package main
//
// import "fmt"
//
//	type linknode struct {
//		next *linknode
//		val  int
//	}
//
//	func main() {
//		node := new(linknode)
//		node.val = 1
//
//		node1 := new(linknode)
//		node1.val = 2
//		node.next = node1
//		for {
//			if node == nil {
//				return
//			} else {
//				fmt.Println(node.val)
//				node = node.next
//			}
//		}
//	}
package main

type Ring struct {
	next, prev *Ring
	value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

//func newRing(n int) *Ring {
//	if n < 1 {
//		return nil
//	}
//	r := new(Ring)
//	r.init()
//	p := r
//	for i := 0; i < n; i++ {
//		p.next = &Ring{nil, nil, i}
//		p = p.next
//		p.prev = r
//		r = r.next
//	}
//	return r
//}

//func (r *Ring) Next() *Ring {
//	if r.next == nil {
//		return r.init()
//	}
//	r = r.next
//	return r
//}
//
//func (r *Ring) Prev() *Ring {
//	if r.prev == nil {
//		return r.init()
//	}
//	r = r.prev
//	return r
//}

//func (r *Ring) remove(n int) *Ring {
//	if n < 0 {
//		return nil
//	}
//	for ; n > 0; n-- {
//		r.next = r.next.next
//		r.next.prev = r.prev
//
//	}
//	return r
//}

func (r *Ring) Long() int {
	if r == nil {
		return 0
	}

	count := 1
	current := r.next

	// 从下一个节点开始遍历，直到回到起始节点 r
	for current != r {
		count++
		current = current.next
	}

	return count
}

func main() {
	r := new(Ring)
	r.init()

}
