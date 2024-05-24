package main

import (
	"fmt"
)

// findMinSubArray 查找和至少为给定值的最短连续子数组
func findMinSubArray(nums []int, sum int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	minLength := n + 1 // 初始化为不可能的最大值
	windowSum := 0
	start, end := 0, 0

	for end < n {
		windowSum += nums[end] // 增大窗口
		for windowSum >= sum {
			if end-start+1 < minLength { // 更新最小长度
				minLength = end - start + 1
			}
			windowSum -= nums[start] // 缩小窗口
			start++                  // 窗口起始位置右移
		}
		end++ // 窗口结束位置右移
	}

	if minLength > n {
		return 0 // 如果没有找到符合条件的子数组，返回0
	}
	return minLength
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	sum := 7
	fmt.Println("Minimum length subarray is:", findMinSubArray(nums, sum))
}
