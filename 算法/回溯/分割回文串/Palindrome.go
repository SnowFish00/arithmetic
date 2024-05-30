package main

import (
	"fmt"
)

// 判断一个字符串是否是回文串
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 回溯函数，寻找所有的回文分割
func backtrack(s string, start int, path []string, result *[][]string) {
	if start == len(s) {
		// 如果起始位置到达字符串末尾，将当前路径加入结果
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// 尝试每一个可能的结束位置
	//这里end 结尾是对的因为 s[start:end] 左开右闭
	for end := start + 1; end <= len(s); end++ {
		substr := s[start:end]
		if isPalindrome(substr) {
			// 如果当前子串是回文，将其加入当前路径，并递归处理剩余的字符串
			path = append(path, substr)
			backtrack(s, end, path, result)
			// 回溯，移除当前子串
			path = path[:len(path)-1]
		}
	}
}

// 主函数，寻找字符串的所有回文分割
func partition(s string) [][]string {
	var result [][]string
	var path []string
	backtrack(s, 0, path, &result)
	return result
}

func main() {
	s := "aab"
	result := partition(s)
	for _, partition := range result {
		fmt.Println(partition)
	}
}
