//package main
//
//import (
//	"sync"
//)
//
////func main() {
////
////}
////
////// zhan
////type ARRSTACK struct {
////	Array []string
////	size  int
////	lock  sync.Mutex
////}
////
////func (stack *ARRSTACK) pash(value string) {
////	stack.lock.Lock()
////	defer stack.lock.Unlock()
////	stack.Array = append(stack.Array, value)
////	stack.size += 1
////}
////
////func (stack *ARRSTACK) unp() {
////	stack.lock.Lock()
////	defer stack.lock.Unlock()
////	if stack.size < 1 {
////		return
////	}
////	value := stack.Array[stack.size-1]
////	fmt.Println(value)
////	stack.size -= 1
////}
////
////func (stack *ARRSTACK) peek() string {
////	stack.lock.Lock()
////	defer stack.lock.Unlock()
////	if stack.size == 0 {
////		return "error"
////	}
////	v := stack.Array[stack.size-1]
////	return v
////}
////
////func (stack *ARRSTACK) sizeof() int {
////	return stack.size
////}
//
//type linkstack struct {
//	root *linknode
//	size int
//	lock sync.Mutex
//}
//type linknode struct {
//	next *linknode
//	val  string
//}
//
//func (l *linkstack) push(n string) {
//	l.lock.Lock()
//	defer l.lock.Unlock()
//	if l.root == nil {
//		l.root = new(linknode)
//		l.root.val = n
//		l.size++
//	}
//	oldnode := l.root.next
//	newnode := new(linknode)
//	newnode.val = n
//	l.root.next = newnode
//	newnode.next = oldnode
//	l.size++
//}
//
//func (l *linkstack) pop() string {
//	l.lock.Lock()
//	defer l.lock.Unlock()
//	if l.size == 0 {
//		return "error"
//	}
//	str := l.root.next.val
//	l.root.next = l.root.next.next
//	l.size--
//	return str
//}
//
//func (l *linkstack) peek() string{
//	//l.lock.Lock()
//	//defer l.lock.Unlock()
//	if l.size == 0 {
//
//		return "empty stack"
//	}
//	val := l.root.next.val
//	return val
//
//}
//
//

package main
