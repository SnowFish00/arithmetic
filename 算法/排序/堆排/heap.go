package main

import (
	"fmt"
)

// heapify函数用于将指定节点及其子节点调整为最大堆
func heapify(arr []int, n, i int) {
	largest := i     // 初始化根节点为最大值
	left := 2*i + 1  // 左子节点的索引
	right := 2*i + 2 // 右子节点的索引

	// 如果左子节点存在并且大于根节点，更新最大值索引为左子节点
	//换成arr[left] < arr[smallest] 就是小根堆
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点存在并且大于当前最大值，更新最大值索引为右子节点
	//换成arr[right] < arr[smallest] 就是小根堆
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点，交换根节点和最大值节点的值，并继续调整最大堆
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// 堆排序函数
func heapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	//我们不需要考虑叶子节点，因为它们已经是最大堆（只有一个元素） n/2 倒数第二层 从n/2 -1 处开始构建 倒数第二层的最后一个节点开始向上遍历
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆顶取出元素，并重新调整堆
	//我们要把最大值（即堆顶元素）移到数组的末尾，然后将剩余部分重新调整为最大堆，这样数组的末尾就是已排序部分了
	/*
			具体操作如下：

		将堆顶元素（即数组的第一个元素）与数组中的最后一个元素交换，这样最大值就被放到了数组的最后。
		然后，我们将堆的大小缩小 1（即 heapify 函数的第二个参数为 i，而 i 是数组长度减 1）。
		最后，我们调用 heapify 函数，将剩余的元素重新调整为最大堆。
		通过重复这个过程，每次将最大值移到数组末尾并缩小堆的大小，直到堆中只剩下一个元素，我们就得到了一个按从小到大顺序排列的数组。
	*/

	//如果是小根堆就不需要调换了
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 将当前堆顶（最大值）与当前末尾元素交换
		heapify(arr, i, 0)              // 调整堆
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Unsorted array is:", arr)
	heapSort(arr)
	fmt.Println("Sorted array is:", arr)
}
