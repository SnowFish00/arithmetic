package main

import "fmt"

// insertionSort 函数实现了插入排序算法，对指定间隔的元素进行排序
func insertionSort(arr []int, start, gap int) {
	n := len(arr)
	// 对从 start + gap 开始，间隔为 gap 的元素进行插入排序
	//把 0 类比为 start 1位gap 你再看
	for i := start + gap; i < n; i += gap {
		key := arr[i]
		j := i - gap
		// 插入排序
		for j >= start && arr[j] > key {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = key
	}
}

// shellSort 函数实现了希尔排序算法
func shellSort(arr []int) {
	n := len(arr)
	// 使用希尔增量序列
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对每个增量进行插入排序
		for i := 0; i < gap; i++ {
			insertionSort(arr, i, gap)
		}
	}
}

func main() {
	// 示例数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	// 调用希尔排序算法对数组进行排序
	shellSort(arr)
	// 打印排序后的数组
	fmt.Println("排序后的数组:", arr)
}
