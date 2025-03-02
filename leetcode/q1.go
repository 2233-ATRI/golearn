package main

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//
// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func preordercomparson(p *TreeNode, q *TreeNode) bool {
	if p.Val == q.Val {
		preordercomparson(p.Left, q.Left)
		preordercomparson(p.Right, q.Right)
	} else if p.Val != q.Val {
		return false
	}
	return true
}
func inordercomnparson(p *TreeNode, q *TreeNode) bool {
	preordercomparson(p.Left, q.Left)
	if p.Val == q.Val {
		preordercomparson(p.Right, q.Right)
	} else if p.Val != q.Val {
		return false
	}
	return true
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return preordercomparson(p, q) && inordercomnparson(p, q)
} //对两个树进行一次中序和一次先序遍历，得到的结果相同则相同
func main() {

}
