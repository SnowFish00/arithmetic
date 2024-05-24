package main

import (
	"fmt"
)

// greedyCoinChange 使用贪心算法实现硬币找零
func greedyCoinChange(coins []int, amount int) []int {
	// 对硬币面额进行降序排序
	sortDescending(coins)
	result := []int{}

	for _, coin := range coins {
		// 循环，直到金额为0
		for amount >= coin {
			result = append(result, coin)
			amount -= coin
		}
	}

	// 如果没有找到合适的硬币组合，返回空数组
	if amount != 0 {
		return []int{}
	}

	return result
}

// sortDescending 降序排序 选择排序
func sortDescending(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	coins := []int{25, 10, 5, 1} // 硬币面额
	amount := 37                 // 需要找零的金额
	change := greedyCoinChange(coins, amount)
	fmt.Printf("Coins to return: %v\n", change)
}
