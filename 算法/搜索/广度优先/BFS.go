package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root} // 初始化队列，并将根节点加入队列
	for len(queue) > 0 {
		// 从队列中取出第一个节点
		node := queue[0]
		queue = queue[1:]
		// 处理当前节点
		fmt.Println(node.Val)
		// 将当前节点的左右子节点加入队列
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
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

	// 执行广度优先搜索
	BFS(root)
}
