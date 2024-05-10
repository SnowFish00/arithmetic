package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i] // 当前要插入的元素
		j := i - 1    // 已排序序列的最后一个元素索引
		// 将当前元素与已排序序列中的元素从后向前逐个比较
		// 找到合适的位置插入当前元素
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // 将较大的元素后移一位
			j--
		}
		arr[j+1] = key // 插入当前元素到合适位置
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	insertionSort(arr)
	fmt.Println("排序后的数组:", arr)
}
