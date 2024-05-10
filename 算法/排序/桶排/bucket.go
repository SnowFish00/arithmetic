package main

import (
	"fmt"
	"sort"
)

// bucketSort 函数实现桶排序算法
// 参数 arr 是待排序的整数切片
// 返回排序后的整数切片
func bucketSort(arr []int) []int {
	// 如果数组为空或者只有一个元素，则无需排序，直接返回
	if len(arr) <= 1 {
		return arr
	}

	// 设置桶的数量，可以根据具体情况调整
	bucketCount := 5

	// 找到数组中的最大值和最小值
	minValue, maxValue := arr[0], arr[0]
	for _, num := range arr {
		if num < minValue {
			minValue = num
		}
		if num > maxValue {
			maxValue = num
		}
	}

	// 计算桶的范围和大小
	bucketSize := (maxValue-minValue)/bucketCount + 1

	// 初始化桶
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}

	// 将元素放入桶中
	for _, num := range arr {
		index := (num - minValue) / bucketSize
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶中的元素进行排序
	sortedArr := make([]int, 0)
	for _, bucket := range buckets {
		sort.Ints(bucket)
		sortedArr = append(sortedArr, bucket...)
	}

	return sortedArr
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := bucketSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
