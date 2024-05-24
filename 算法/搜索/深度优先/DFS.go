package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DFS(root *TreeNode) {
	//没使用队列 而是栈递归 所以未逐层遍历 而是先左根到底 符合 深度优先
	//这个是根左右的前序深度遍历 还可以写其他的顺序
	if root == nil {
		return
	}
	// 前序遍历位置
	fmt.Println(root.Val)
	DFS(root.Left)
	DFS(root.Right)

}

func main() {
	// 创建一个简单的二叉树
	//       1
	//      / \
	//     2   3
	//    / \
	//   4   5
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// 执行深度优先搜索
	DFS(root)
}
