package main

import "sync"

type sxlist struct {
	head *sxnode
	tail *sxnode
	len  int
	lock sync.Mutex
}
type sxnode struct {
	pre  *sxnode
	next *sxnode
	val  interface{}
}

func (s *sxlist) get() interface{} {
	return s.tail.val
}

func (s *sxlist) qqget(val interface{}) interface{} {
	return s.tail.pre.val
}

// GetValue 获取节点值
func (node *sxnode) GetValue() interface{} {
	return node.val
}

// GetPre 获取节点前驱节点
//func (node *ListNode) GetPre() *ListNode {
//	return node.pre
//}
//// GetNext 获取节点后驱节点
//func (node *ListNode) GetNext() *ListNode {
//	return node.next
//}
//// HashNext 是否存在后驱节点
//func (node *ListNode) HashNext() bool {
//	return node.pre != nil
//}
//// HashPre 是否存在前驱节点
//func (node *ListNode) HashPre() bool {
//	return node.next != nil
//}
//// IsNil 是否为空节点
//func (node *ListNode) IsNil() bool {
//	return node == nil
//}
//// Len 返回列表长度
//func (list *DoubleList) Len() int {
//	return list.len
//}

func main() {

}
