package main

import (
	"fmt"
)

func main() {
	D := [][]int{
		{7},
		{3, 8},
		{8, 1, 0},
		{2, 7, 4, 4},
		{4, 5, 2, 6, 5},
	}
	n := len(D) // n表示层数
	maxSum := make([][]int, n)
	// 初始化maxSum数组
	for i := range maxSum {
		maxSum[i] = make([]int, i+1)
	}

	// 从底部开始向上计算最大路径和
	//自下而上每次计算时舍去较小的和max(a,b) 从而后面和这个相关的计算可能性舍去简化时间复杂度
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if i == n-1 {
				// 最后一行的值就是其本身
				maxSum[i][j] = D[i][j]
			} else {
				// 其他行的值是其下方和右下方两个值的较大者加上当前值
				maxSum[i][j] = D[i][j] + max(maxSum[i+1][j], maxSum[i+1][j+1])
			}
		}
	}

	// 最大路径和是顶部的值
	fmt.Println("最大路径和为:", maxSum[0][0])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
