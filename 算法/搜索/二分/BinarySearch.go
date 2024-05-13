package main

import "fmt"

// BinarySearch 执行二分查找，返回目标值的索引，如果未找到则返回-1
func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1 // 初始化左指针和右指针

	for left <= right {
		mid := left + (right-left)/2 // 计算中间索引，防止整数溢出

		// 检查目标值是否在中间位置
		if arr[mid] == target {
			return mid
		}

		// 如果目标值大于中间值，则忽略左半部分，移动左指针
		if arr[mid] < target {
			left = mid + 1
		} else {
			// 如果目标值小于中间值，则忽略右半部分，移动右指针
			right = mid - 1
		}
	}

	// 目标值不在数组中
	return -1
}

func main() {
	// 示例已排序数组
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 15 // 要查找的目标值

	// 执行二分查找
	result := BinarySearch(arr, target)

	if result != -1 {
		fmt.Printf("元素在索引 %d 处找到\n", result)
	} else {
		fmt.Println("元素在数组中未找到")
	}
}
