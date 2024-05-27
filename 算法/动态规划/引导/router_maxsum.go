package main

import (
	"fmt"
)

func main() {
	/*
		7
		3 8
		8 1 0
		2 7 4 4
		4 5 2 6 5
		从上到下选择一条路，使得经过的数字之和最大。
		路径上的每一步只能往下或者右下走。
	*/

	D := [][]int{
		{7},
		{3, 8},
		{8, 1, 0},
		{2, 7, 4, 4},
		{4, 5, 2, 6, 5},
	}
	n := len(D) // n表示层数
	maxSum := getMaxSum(D, n, 0, 0)
	fmt.Println("最大路径和为:", maxSum)
}

// 非动态规划 递归完成 (终止条件 操作 入参递归重复)
func getMaxSum(D [][]int, n, i, j int) int {
	if i == n-1 { // 到达最后一行
		return D[i][j]
	}
	x := getMaxSum(D, n, i+1, j)   // 向下
	y := getMaxSum(D, n, i+1, j+1) // 向右下
	return max(x, y) + D[i][j]     // 返回两条路径中较大的一个加上当前节点的值
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
