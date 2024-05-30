package main

import "fmt"

const n = 4      // 物品数量
const TotCap = 5 // 背包总容量

var w = [5]int{0, 2, 3, 4, 5}   // 物品重量数组，索引0不使用，从1开始
var val = [5]int{0, 3, 4, 5, 6} // 物品价值数组，索引0不使用，从1开始
var bestx = [5]int{}            // 存储最优解，即每个物品是否被选中
var bestval = 0                 // 存储最大价值

// dfs是深度优先搜索函数，用于回溯求解背包问题
// i表示当前考虑的物品索引
// cv表示当前背包的总价值
// cw表示当前背包的总重量
// x是一个局部变量，用于在递归过程中跟踪当前路径上的物品选择状态
func dfs(i, cv, cw int, x [5]int) {
	// 如果已经考虑完所有物品
	if i > n {
		// 如果当前价值大于最大价值，则更新最大价值和最优解
		if cv > bestval {
			bestval = cv
			copy(bestx[:], x[:])
		}
	} else {
		// 枚举物体i所有可能的路径，0表示不选，1表示选
		for j := 0; j <= 1; j++ {
			// 如果选择当前物品，判断是否满足重量约束
			if cw+int(j)*w[i] <= TotCap {
				// 选择当前物品，更新价值和重量
				cw += int(j) * w[i]
				cv += int(j) * val[i]
				x[i] = j // 更新当前解
				// 继续考虑下一个物品
				dfs(i+1, cv, cw, x)
				// 回溯，撤销选择当前物品，更新价值和重量
				cw -= int(j) * w[i]
				cv -= int(j) * val[i]
				x[i] = 0 // 回溯解
			}
		}
	}
}

func main() {
	// 从第一个物品开始，价值和重量都是0，调用dfs函数
	dfs(1, 0, 0, bestx)
	fmt.Println("最大价值为:", bestval)
	fmt.Println("选中的物品为:", bestx[1:])
}
