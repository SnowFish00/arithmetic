package main

import (
	"fmt"
)

// 递归版本 由n到0去减
func minPathSumD(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	return getMin(grid, m-1, n-1)
}

func getMin(grid [][]int, i, j int) int {
	// 起点位置
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	// 第一行，想到达这里，说明前面肯定是水平向右，否则肯定不会在第一行
	if i == 0 {
		return grid[i][j] + getMin(grid, i, j-1)
	}
	// 第一列，说明肯定是垂直到打这里。否则肯定不会在第一列
	if j == 0 {
		return grid[i][j] + getMin(grid, i-1, j)
	}

	//minRow 是到达 (i, j) 位置的所有从上方路径中的最小路径和，而 minColumn 是到达 (i, j) 位置的所有从左边路径中的最小路径和
	// 直接返回当前 + 前面的最小值即可。
	minRow := getMin(grid, i-1, j)
	minColumn := getMin(grid, i, j-1)
	// 找到最小的距离
	return grid[i][j] + min(minRow, minColumn)
}

// 动态规划
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// 初始化dp数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]

	// 初始化第一行和第一列
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// 循环处理剩余的单元格
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 使用状态转移方程
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	// 返回右下角单元格的值
	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSumD(grid)) // 输出应该是最小路径和
	fmt.Println(minPathSum(grid))  // 输出应该是最小路径和
}
