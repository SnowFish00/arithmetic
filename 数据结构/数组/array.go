package main

import (
	"fmt"
	"math/rand"
)

// 扩展数组长度
func extend(array []int, enlarge int) []int {
	res := make([]int, len(array)+enlarge)
	//拷贝回去
	for index, item := range array {
		res[index] = item
	}
	return res
}

// 插入元素
func insert(array []int, num int, index int) []int {
	//从最后面向前移动腾出index ，数组移动前记得扩容(不然最后一位就寄了) 毕竟是移动逻辑不是append
	//循环外负责移动
	for i := len(array) - 1; i > index; i-- {
		//循环内负责操作
		//以此向后移一位
		array[i] = array[i-1]
	}

	array[index] = num
	return array
}

// 删除元素
func delete(array []int, index int) []int {
	//由当前index 逐个向前覆盖
	for i := index; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}

	//最后一位去除
	array = array[:len(array)-1]

	return array
}

// 修改
func modify(array []int, index int, num int) []int {
	array[index] = num
	return array
}

// 查找
func search(array []int, num int) int {
	for i, it := range array {
		if it == num {
			return i
		}
	}

	return -1
}

// 遍历数组
func all(arry []int) {
	for _, it := range arry {
		fmt.Println(it)
	}
}

// 初始化
func array_init() []int {
	arr := []int{1, 2, 3, 4, 5}
	return arr

}

// 随机访问
func randoms(array []int) int {
	randomIndex := rand.Intn(len(array))
	return array[randomIndex]
}

func main() {
	//初始化
	array := array_init()
	//遍历
	all(array)
	//扩容
	array = extend(array, 1)
	//增加
	array = insert(array, 114514, 2)
	fmt.Println(array)
	//删除
	array = delete(array, 0)
	fmt.Println(array)
	//修改
	array = modify(array, 1, 8848)
	fmt.Println(array)
	//查询
	index := search(array, 8848)
	fmt.Println(index)
	//随机访问
	fmt.Println(randoms(array))
}
