package main

import (
	"fmt"
)

/*
1.滑动窗口用来解释连续的数字元素
2.滑动窗口直接或间接有两个指针滑动窗口大小可变
*/

// maxSubArray 使用滑动窗口（Kadane's Algorithm）找到最大子数组和
// maxSubArray 则隐式地使用当前元素位置和之前的累积和
func maxSubArray(nums []int) int {
	// 如果数组为空，返回0
	if len(nums) == 0 {
		return 0
	}

	// 初始化当前子数组和和最大子数组和
	currentSum := nums[0]
	maxSum := nums[0]

	// 从第二个元素开始遍历数组
	for i := 1; i < len(nums); i++ {
		// 如果当前子数组和为负数，则丢弃之前的子数组，从当前元素重新开始(剪枝)
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			// 否则，将当前元素加入当前子数组和中
			currentSum += nums[i]
		}

		// 更新最大子数组和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {
	// 示例数组
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// 找到最大子数组和并打印结果
	fmt.Println("最大子数组和为:", maxSubArray(nums))
}
