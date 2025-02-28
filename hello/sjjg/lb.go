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

func main() {
	r := new(Ring)
	r.init()

}
