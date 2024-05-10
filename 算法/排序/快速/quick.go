package main

import "fmt"

// 快速排序函数
func quickSort(arr []int) []int {
	// 如果数组长度小于等于1，则返回数组本身，因为一个元素或者空数组都是有序的
	if len(arr) <= 1 {
		return arr
	}

	// 选择基准值，这里选择第一个元素作为基准值
	pivot := arr[0]
	var left, right []int

	// 分区操作：将小于基准值的元素放入左侧列表，大于等于基准值的元素放入右侧列表
	for _, x := range arr[1:] {
		if x < pivot {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
	}

	// 递归地对左侧和右侧列表进行快速排序，然后合并结果
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	// 示例数组
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("原始数组:", arr)

	// 调用快速排序函数
	sortedArr := quickSort(arr)
	fmt.Println("排序后的数组:", sortedArr)
}
