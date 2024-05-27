/*
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
*/

/*
我们先说下递归法。

像走到终点，实际只有 2 条路：

如果是 m*n 的网格，目标就是 (m, n)

（1）走到终点的左边 (m-1, n)

（2）走到终点的上边 (m, n-1)

不同的路径就是上面两条路的和。
*/

package main

import "fmt"

// 递归
func uniquePathsD(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePathsD(m-1, n) + uniquePathsD(m, n-1)
}

// 动态规划
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化第一行和第一列 因为只能向左或者向右移动所以第一排和第一列只能有一种方法所以均设为1
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	// 计算其他位置的路径数量
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			//状态方程f(m, n) = f(m-1, n) + f(m, n-1);
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// 返回右下角的路径数量
	return dp[m-1][n-1]
}

func main() {
	m := 3
	n := 7
	fmt.Println("Unique paths for a", m, "x", n, "grid:", uniquePaths(m, n))
}
