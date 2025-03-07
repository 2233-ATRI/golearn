package main

// 公司计划面试 2n 人。给你一个数组 costs ，其中 costs[i] = [aCosti, bCosti] 。第 i 人飞往 a 市的费用为 aCosti ，飞往 b 市的费用为 bCosti 。
//
// 返回将每个人都飞到 a 、b 中某座城市的最低费用，要求每个城市都有 n 人抵达。
//
//	func twoCitySchedCost(costs [][]int) int {
//		// 计算差值并存储原始索引
//		type person struct {
//			diff int // aCosti - bCosti
//			aCost int
//			bCost int
//		}
//		n := len(costs) / 2
//		people := make([]person, len(costs))
//		for i, cost := range costs {
//			people[i] = person{
//				diff:  cost[0] - cost[1],
//				aCost: cost[0],
//				bCost: cost[1],
//			}
//		}
//
//		// 按差值升序排序
//		sort.Slice(people, func(i, j int) bool {
//			return people[i].diff < people[j].diff
//		})
//
//		// 前n个去A，后n个去B
//		totalCost := 0
//		for i := 0; i < n; i++ {
//			totalCost += people[i].aCost // 去A
//		}
//		for i := n; i < 2*n; i++ {
//			totalCost += people[i].bCost // 去B
//		}
//
//		return totalCost
//	} //轮询即可，先找a地的最小值然后c=c+a,然后排除这个元素之后剩下
//
// b最小的输出c=c+b这样循环直到数组 全被输出
func twoCitySchedCost(costs [][]int) int {
	n := len(costs) / 2
	ac := make([]int, len(costs))
	for i := 0; i < len(costs); i++ {
		ac[i] = costs[i][0] - costs[i][1] // 差值：去A比去B的额外成本
	}

	// 复制costs，避免修改原数组
	ab := make([][]int, len(costs))
	for i := range costs {
		ab[i] = make([]int, 2)
		ab[i][0] = costs[i][0]
		ab[i][1] = costs[i][1]
	}

	// 按差值排序，同时调整ab
	for i := 0; i < len(ab)-1; i++ { // 修正：len(ab)-1，避免越界
		for j := i + 1; j < len(ab); j++ { // 从i+1开始，避免重复比较
			if ac[i] > ac[j] { // 注意：应为>号，确保升序
				// 交换ac
				ac[i], ac[j] = ac[j], ac[i]
				// 同步交换ab
				ab[i], ab[j] = ab[j], ab[i]
			}
		}
	}

	// 前n个去A，后n个去B
	total := 0
	for i := 0; i < n; i++ {
		total += ab[i][0] // 去A的费用
	}
	for i := n; i < 2*n; i++ {
		total += ab[i][1] // 去B的费用
	}

	return total
}
func main() {

}
