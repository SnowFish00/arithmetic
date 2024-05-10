package main

import "fmt"

func bubble(list []int) {
	len := len(list)
	//外层控制循环次数
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			//> 则 大数放在后面
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}

	}

}

func main() {
	// 待排序的数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)

	// 调用冒泡排序函数
	bubble(arr)

	fmt.Println("Sorted array:", arr)
}
