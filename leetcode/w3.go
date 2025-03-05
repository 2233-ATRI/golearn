package main

import (
	"fmt"
)

//请设计一个函数判断一棵二叉树是否 轴对称 。

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	func maxheight(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//		leftHeight := maxheight(root.Left)
//		rightHeight := maxheight(root.Right)
//		return max(leftHeight, rightHeight)
//	}
//
//	func fillarra(root *TreeNode, arr []*int, index int) {
//		if root == nil || index >= len(arr) {
//			return
//		}
//		arr[index] = &root.Val
//		fillarra(root.Left, arr, index*2+1)
//		fillarra(root.Right, arr, index*2+2)
//	}
//
//	func checkSymmetricTree(root *TreeNode) bool {
//		if root == nil {
//			return true
//		}
//
//		// 获取树的高度以确定数组大小
//		height := maxheight(root)
//		// 完全二叉树的最大节点数为 2^height - 1
//		size := 1<<height - 1
//		arr := make([]*int, size) // 使用指针区分nil和0
//
//		// 将树填充到数组中
//		fillarra(root, arr, 0)
//
//		// 检查每一层的对称性
//		for level := 0; level < height; level++ {
//			start := 1<<level - 1   // 每层起始索引
//			end := 1<<(level+1) - 1 // 每层结束索引
//			if end > size {
//				end = size
//			}
//			for i, j := start, end-1; i < j; i, j = i+1, j-1 {
//				// 获取对称位置的值
//				var leftVal, rightVal int
//				if arr[i] == nil {
//					leftVal = -1 // 用-1表示空节点
//				} else {
//					leftVal = *arr[i]
//				}
//				if arr[j] == nil {
//					rightVal = -1
//				} else {
//					rightVal = *arr[j]
//				}
//				// 对称位置必须相等（包括都为空）
//				if leftVal != rightVal {
//					return false
//				}
//			}
//		}
//		return true
//	} //直接将树放入一个完全二叉树节点数量大小的数组当中
//
// 当二叉树同层对称元素应该相同设对称两个元素之和为n对应数组
// 当中的n-i和i应该是相同的
// 超时
func checkSymmetricTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// 辅助函数：检查两棵子树是否互为镜像
func isMirror(left, right *TreeNode) bool {
	// 如果都为空，返回true
	if left == nil && right == nil {
		return true
	}
	// 如果一个为空一个不为空，返回false
	if left == nil || right == nil {
		return false
	}
	// 检查当前节点值是否相等，且子树是否镜像对称
	return left.Val == right.Val &&
		isMirror(left.Left, right.Right) &&
		isMirror(left.Right, right.Left)
}
func main() {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(node)
	fmt.Println(checkSymmetricTree(&TreeNode{}))
}
