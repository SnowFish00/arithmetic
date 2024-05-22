package main

import "fmt"

// mergeSort 函数用于对整个数组进行归并排序
func mergeSort(arr []int) []int {
	length := len(arr)
	// 如果数组长度小于 2，直接返回该数组
	if length < 2 {
		return arr
	}
	// 计算数组中间位置
	middle := length / 2
	// 分割数组为左右两部分
	left := arr[0:middle]
	right := arr[middle:]
	// 递归调用 mergeSort 函数对左右两部分进行排序，并合并结果
	return merge(mergeSort(left), mergeSort(right))
}

// merge 函数用于合并两个已排序的子数组
func merge(left []int, right []int) []int {
	var result []int
	// 比较左右两个子数组的元素，并将较小的元素依次添加到结果数组中
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	//下面这两个for为了让奇数长度下左右数组不等长的超出部分衔接的
	// 将剩余的元素添加到结果数组中
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	// 返回合并后的结果数组
	return result
}

func main() {
	// 示例数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	// 调用 mergeSort 函数对数组进行归并排序
	sortedArr := mergeSort(arr)
	// 打印排序后的数组
	fmt.Println("排序后的数组:", sortedArr)
}
