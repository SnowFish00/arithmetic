/*
在 8×8 格的国际象棋上摆放八个皇后，使其不能互相攻击，即任意两个皇后都不能处于同一行、同一列或同一斜线上，问有多少种摆法。
*/

package main

import (
	"fmt"
)

// g_number 用于记录找到的解决方案的数量
var g_number = 0

// EightQueen 是解决8皇后问题的入口函数
func EightQueen() {
	// queens 表示皇后的数量，这里固定为8
	const queens = 8
	// 创建一个长度为queens的切片，用于存储每一列皇后的行索引
	columnIndex := make([]int, queens)
	// 初始化columnIndex，使其为[0, 1, 2, ..., queens-1]
	for i := 0; i < queens; i++ {
		columnIndex[i] = i
	}
	// 调用Permutation函数开始递归排列，并检查每种排列是否是有效的解决方案
	Permutation(columnIndex, queens, 0)
}

// Permutation 是一个递归函数，用于生成所有可能的皇后排列
func Permutation(columnIndex []int, length int, index int) {
	// 如果index等于length，说明已经生成了一个完整的排列
	if index == length {
		// 检查这个排列是否有效
		if Check(columnIndex, length) {
			// 如果是有效解决方案，增加解决方案计数
			g_number++
			// 打印解决方案
			PrintQueen(columnIndex, length)
		}
	} else {
		// 如果index小于length，继续生成排列
		//由于初始化斜线 然后每个递归位置都是向下 所以行不会重复
		for i := index; i < length; i++ {
			// 交换columnIndex[i]和columnIndex[index]，以便生成新的排列
			columnIndex[i], columnIndex[index] = columnIndex[index], columnIndex[i]
			// 递归调用Permutation，处理下一个位置
			Permutation(columnIndex, length, index+1)
			// 交换回来，回溯到上一步的状态
			columnIndex[i], columnIndex[index] = columnIndex[index], columnIndex[i]
		}
	}
}

// Check 检查给定的皇后排列是否有效
func Check(columnIndex []int, length int) bool {
	// 遍历所有皇后
	for i := 0; i < length; i++ {
		// 检查其他皇后
		for j := i + 1; j < length; j++ {
			// 如果两个皇后在同一行或者在同一对角线上，则返回false
			//主副对角线判断
			/*
				主对角线：从左上角到右下角的对角线。对于棋盘上的任意一点(x, y)，它在主对角线上的所有点都有一个共同的特点：x - y的值是相同的。
				例如，点(0,0)、(1,1)、(2,2)等都在同一条主对角线上，因为它们的x - y值都是0。
				次对角线：从右上角到左下角的对角线。对于棋盘上的任意一点(x, y)，它在次对角线上的所有点都有一个共同的特点：x + y的值是相同的。
				例如，点(0,7)、(1,6)、(2,5)等都在同一条次对角线上，因为它们的x + y值都是7。
			*/
			if (i-j == columnIndex[i]-columnIndex[j]) || (j-i == columnIndex[i]-columnIndex[j]) {
				return false
			}
		}
	}
	// 如果没有冲突，返回true
	return true
}

// PrintQueen 打印一个有效的皇后排列
func PrintQueen(columnIndex []int, length int) {
	// 打印解决方案编号
	fmt.Printf("Solution %d\n", g_number)
	// 打印每一列皇后的行索引
	for i := 0; i < length; i++ {
		fmt.Printf("%d\t", columnIndex[i])
	}
	fmt.Println() // 打印换行符，分隔不同的解决方案
}

func main() {
	// 调用EightQueen函数开始解决8皇后问题
	EightQueen()
}
