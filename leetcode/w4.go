//package main
//
//// 我们可以为二叉树 T 定义一个 翻转操作 ，如下所示：选择任意节点，然后交换它的左子树和右子树。
////
//// 只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转 等价 于二叉树 Y。
////
//// 这些树由根节点 root1 和 root2 给出。如果两个二叉树是否是翻转 等价 的函数，则返回 true ，否则返回 false 。
//// * Definition for a binary tree node.
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
//	if root1 == nil && root2 == nil {
//		return true
//	}
//	if root1 == nil || root2 == nil {
//		return false
//	}
//	if root1.Val != root2.Val {
//		return false
//	}
//	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) || (flipEquiv(root1.Left, root2.Right)) && (flipEquiv(root1.Right, root2.Left))
//}
//
//// 递归进行查看只要可以实现x左右节点交换之后和y相同或本就相同则都进入下一层对下一层进行
//// 对比，直到叶子节点，中间出现交换不相同则是返回flase
//func main() {
//
//}
package main
