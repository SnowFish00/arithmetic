package main

import "fmt"

/*
我们直接从最底下，最简单，问题规模最小的 f(1) 和 f(2) 开始往上推，直到推到我们想要的答案 f(20)，这就是动态规划的思路，
这也是为什么动态规划一般都脱离了递归，而是由循环迭代完成计算。
*/

func fib(N int) int {
	dp := make([]int, N+1)
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= N; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[N]
}

func main() {
	N := 10 // 示例，计算第10个斐波那契数
	result := fib(N)
	fmt.Println(result)
}
