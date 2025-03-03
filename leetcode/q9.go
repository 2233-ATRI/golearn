package main

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0 // 伪回文路径计数器

	// DFS 遍历，state 表示路径上数字出现次数的奇偶性
	var dfs func(node *TreeNode, state int)
	dfs = func(node *TreeNode, state int) {
		if node == nil {
			return
		}

		// 更新当前节点值的奇偶状态，使用位运算异或
		state ^= (1 << node.Val)

		// 如果是叶子节点，检查伪回文条件
		if node.Left == nil && node.Right == nil {
			// 检查 state 中 1 的个数是否 ≤ 1
			if countOnes(state) <= 1 {
				count++
			}
			return
		}

		// 递归遍历左右子树
		dfs(node.Left, state)
		dfs(node.Right, state)
	}

	dfs(root, 0) // 从根节点开始，初始状态为 0
	return count
}

// countOnes 计算整数中 1 的个数
func countOnes(n int) int {
	count := 0
	for n > 0 {
		count += n & 1 // 检查最低位是否为 1
		n >>= 1        // 右移一位
	}
	return count
}

// 可以使用深度优先遍历搜索，
// 当得到一个节点没有左右子树时对得到的数据进行查看，
// 主要数据中最多只出现了一个奇数就可以对计数器+1
// 这个过程使用单个位运算码，只要是0则偶数个，1则奇数个
