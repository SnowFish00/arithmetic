package main

import "fmt"

func selectionSort(arr []int) {
	len := len(arr)

	for i := 0; i < len-1; i++ {
		//假定当前位置元素最小
		minindex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minindex] {
				//换下标
				minindex = j
			}
		}

		//将当前最小元素交换到假定的最小元素位置
		arr[i], arr[minindex] = arr[minindex], arr[i]

	}

}

func main() {
	// 待排序的数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	// 调用选择排序函数
	selectionSort(arr)

	fmt.Println("Sorted array:", arr)
}
