package main

import "fmt"

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func CreateBinaryTreeNode(data int) *BinaryTreeNode {
	return &BinaryTreeNode{
		Data:  data,
		Left:  nil,
		Right: nil,
	}
}

// 插入结点
func (node *BinaryTreeNode) Insert(root *BinaryTreeNode, val int) bool {
	per := root

	for per != nil {
		if per.Data > val {
			if per.Left != nil {
				per = per.Left
			} else {
				per.Left = CreateBinaryTreeNode(val)
				return true
			}
		} else {
			if per.Right != nil {
				per = per.Right
			} else {
				per.Right = CreateBinaryTreeNode(val)
				return true
			}
		}
	}

	return false
}

// 层数打印
func (node *BinaryTreeNode) BreadthFirstSearch() []int {
	if node == nil {
		return nil
	}
	//结果切片
	var result []int
	per := node
	//节点切片
	node_list := []*BinaryTreeNode{per}

	for len(node_list) > 0 {
		result = append(result, node_list[0].Data)
		//没遇见叶子结点则添加根结点
		if node_list[0].Left != nil {
			node_list = append(node_list, per.Left)
		}
		if node_list[0].Right != nil {
			node_list = append(node_list, per.Right)
		}

		//弹出队列第一个元素 根
		node_list = node_list[1:]

		// 更新当前节点为队列的第一个节点
		if len(node_list) > 0 {
			per = node_list[0]
		}
	}

	return result
}

// 前序遍历
func (node *BinaryTreeNode) PreOrder(root *BinaryTreeNode) {
	if root != nil {
		fmt.Println(root.Data)
		root.PreOrder(root.Left)
		root.PreOrder(root.Right)
	}
}

// 中序遍历
func (node *BinaryTreeNode) InOrder(root *BinaryTreeNode) {
	if root != nil {
		root.InOrder(root.Left)
		fmt.Println(root.Data)
		root.InOrder(root.Right)
	}
}

// 后序遍历
func (node *BinaryTreeNode) PostOrder(root *BinaryTreeNode) {
	if root != nil {
		root.PostOrder(root.Left)
		root.PostOrder(root.Right)
		fmt.Println(root.Data)
	}
}

// 获得树高
func (node *BinaryTreeNode) GetHeight(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	l := node.GetHeight(root.Left)
	r := node.GetHeight(root.Right)

	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

// 打印叶子结点
func (node *BinaryTreeNode) FindLeaf(root *BinaryTreeNode) {

	//排除根结点或者到了叶子结点 为空
	if root != nil {
		if root.Left == nil && root.Right == nil {
			fmt.Println(root.Data)
		}
		root.FindLeaf(root.Left)
		root.FindLeaf(root.Right)

	}

}

// 根据值查找node
func (node *BinaryTreeNode) FindValueNode(root *BinaryTreeNode, target int) *BinaryTreeNode {
	if root != nil {
		if target == root.Data {
			return root
		} else {
			//递归函数所有的条件都要有 return 使得逻辑链接起来递归
			if found := root.FindValueNode(root.Left, target); found != nil {
				//若左子树找到了则没必要找又子树了
				return found
			}
			//左子树没找到找右子树
			return root.FindValueNode(root.Right, target)
		}
	}
	return nil
}

func main() {
	var node *BinaryTreeNode
	// 创建一个根节点
	node = CreateBinaryTreeNode(10)
	li := []int{9, 11, 8, 5, 6, 4, 12, 15, 18, 17}
	// 创建一个二叉树
	for _, val := range li {
		node.Insert(node, val)
	}
	ret := node.BreadthFirstSearch()
	fmt.Println("层级打印如下")
	fmt.Println(ret)
	fmt.Println("先序遍历")
	node.PreOrder(node)
	fmt.Println("中序遍历")
	node.InOrder(node)
	fmt.Println("后序遍历")
	node.PostOrder(node)
	fmt.Println("树高")
	res := node.GetHeight(node)
	fmt.Println(res)
	fmt.Println("叶子结点")
	node.FindLeaf(node)
	fmt.Println("值为17的结点")
	ref := node.FindValueNode(node, 17)
	fmt.Println(ref)
}
