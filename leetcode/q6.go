package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxdata :=0
	var height func(root *TreeNode) int{
		height = func(root *TreeNode) int{
			if root == nil {
				return 0
			}
			liftheight := height(root.Left)
			rightheight := height(root.Right)
			currHeight := liftheight + rightheight
			if currHeight > maxdata {
				maxdata = currHeight
			}
			return max(liftheight, rightheight) + 1
		}
		height(root)
		return maxdata
	}
} //本质是直接找通过根节点或者最早分歧点进行找两边最深的节点
// 每个节点的深度优先遍历
func main() {

}
