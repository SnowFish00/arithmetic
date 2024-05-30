package main

import (
	"fmt"
)

// 定义红黑树的颜色常量
const (
	__RED   = true  // 红色
	__BLACK = false // 黑色
)

// RbTreeColor 代表红黑树节点的颜色
type RbTreeColor bool

// RbTreeKeyType 代表红黑树节点键的类型
type RbTreeKeyType int

// RbTreeValueType 代表红黑树节点值的类型
type RbTreeValueType interface{}

// RbTreeNode 代表红黑树中的节点
type RbTreeNode struct {
	Color  RbTreeColor     // 节点颜色
	Parent *RbTreeNode     // 父节点指针
	Left   *RbTreeNode     // 左子节点指针
	Right  *RbTreeNode     // 右子节点指针
	Key    RbTreeKeyType   // 节点键
	Value  RbTreeValueType // 节点值
}

// RbTree 代表红黑树结构
type RbTree struct {
	Root     *RbTreeNode // 根节点
	sentinel *RbTreeNode // 哨兵节点，表示空节点
	NodeNum  int         // 节点数量
}

// NewRbTree 创建并返回一个新的红黑树实例
func NewRbTree() *RbTree {
	sentinel := &RbTreeNode{}  // 创建一个新的哨兵节点
	sentinel.Color = __BLACK   // 设置哨兵节点的颜色为黑色
	sentinel.Left = sentinel   // 哨兵节点的左子节点指向自己
	sentinel.Right = sentinel  // 哨兵节点的右子节点指向自己
	sentinel.Parent = sentinel // 哨兵节点的父节点指向自己
	sentinel.Value = nil       // 哨兵节点的值为空
	sentinel.Key = -9999       // 设置哨兵节点的键为一个特殊值
	return &RbTree{
		Root:     sentinel, // 红黑树的根节点初始为哨兵节点
		sentinel: sentinel, // 哨兵节点
		NodeNum:  0,        // 初始节点数量为0
	}
}

// FindMinNodeBy 查找以node为根的最小节点
func (node *RbTreeNode) FindMinNodeBy(rbTreeNilNode *RbTreeNode) *RbTreeNode {
	newNode := node                     // 从当前节点开始
	for newNode.Left != rbTreeNilNode { // 一直向左遍历，直到达到哨兵节点
		newNode = newNode.Left
	}
	return newNode // 返回最左侧的节点，即最小节点
}

// FindMaxNodeBy 查找以node为根的最大节点
func (node *RbTreeNode) FindMaxNodeBy(rbTreeNilNode *RbTreeNode) *RbTreeNode {
	newNode := node                      // 从当前节点开始
	for newNode.Right != rbTreeNilNode { // 一直向右遍历，直到达到哨兵节点
		newNode = newNode.Right
	}
	return newNode // 返回最右侧的节点，即最大节点
}

// FindMax 查找红黑树中的最大节点
func (rbTree *RbTree) FindMax() *RbTreeNode {
	return rbTree.Root.FindMaxNodeBy(rbTree.sentinel) // 从根节点开始查找最大节点
}

// FindMin 查找红黑树中的最小节点
func (rbTree *RbTree) FindMin() *RbTreeNode {
	return rbTree.Root.FindMinNodeBy(rbTree.sentinel) // 从根节点开始查找最小节点
}

// LeftRotate 对红黑树进行左旋操作
func (rbTree *RbTree) LeftRotate(node *RbTreeNode) {
	tmpNode := node.Right                // 保存右子节点
	node.Right = tmpNode.Left            // 将右子节点的左子节点设置为当前节点的右子节点
	if tmpNode.Left != rbTree.sentinel { // 如果右子节点的左子节点不是哨兵节点
		tmpNode.Left.Parent = node // 将右子节点的左子节点的父节点设置为当前节点
	}
	tmpNode.Parent = node.Parent        // 将右子节点的父节点设置为当前节点的父节点
	if node.Parent == rbTree.sentinel { // 如果当前节点的父节点是哨兵节点
		rbTree.Root = tmpNode // 将右子节点设置为根节点
	} else if node == node.Parent.Left { // 如果当前节点是其父节点的左子节点
		node.Parent.Left = tmpNode // 将右子节点设置为其父节点的左子节点
	} else { // 如果当前节点是其父节点的右子节点
		node.Parent.Right = tmpNode // 将右子节点设置为其父节点的右子节点
	}
	tmpNode.Left = node   // 将当前节点设置为右子节点的左子节点
	node.Parent = tmpNode // 将当前节点的父节点设置为右子节点
}

// RightRotate 对红黑树进行右旋操作
func (rbTree *RbTree) RightRotate(node *RbTreeNode) {
	tmpNode := node.Left                  // 保存左子节点
	node.Left = tmpNode.Right             // 将左子节点的右子节点设置为当前节点的左子节点
	if tmpNode.Right != rbTree.sentinel { // 如果左子节点的右子节点不是哨兵节点
		tmpNode.Right.Parent = node // 将左子节点的右子节点的父节点设置为当前节点
	}
	tmpNode.Parent = node.Parent        // 将左子节点的父节点设置为当前节点的父节点
	if node.Parent == rbTree.sentinel { // 如果当前节点的父节点是哨兵节点
		rbTree.Root = tmpNode // 将左子节点设置为根节点
	} else if node == node.Parent.Right { // 如果当前节点是其父节点的右子节点
		node.Parent.Right = tmpNode // 将左子节点设置为其父节点的右子节点
	} else { // 如果当前节点是其父节点的左子节点
		node.Parent.Left = tmpNode // 将左子节点设置为其父节点的左子节点
	}
	tmpNode.Right = node  // 将当前节点设置为左子节点的右子节点
	node.Parent = tmpNode // 将当前节点的父节点设置为左子节点
}

// Insert 向红黑树中插入一个新的节点
func (rbTree *RbTree) Insert(key RbTreeKeyType, value RbTreeValueType) {
	rbTree.InsertNewNode(&RbTreeNode{Key: key, Value: value}) // 创建新节点并插入
}

// InsertNewNode 向红黑树中插入一个新的节点，并进行颜色和结构上的调整
func (rbTree *RbTree) InsertNewNode(node *RbTreeNode) {
	newNodeParent := rbTree.sentinel // 初始化新节点的父节点为哨兵节点
	tmpNode := rbTree.Root           // 从根节点开始查找插入位置

	// 搜寻目标节点的父节点
	for tmpNode != rbTree.sentinel { // 遍历树寻找插入位置
		newNodeParent = tmpNode
		if node.Key < tmpNode.Key { // 如果新节点的键小于当前节点的键
			tmpNode = tmpNode.Left // 向左子树遍历
		} else if node.Key > tmpNode.Key { // 如果新节点的键大于当前节点的键
			tmpNode = tmpNode.Right // 向右子树遍历
		} else { // 如果新节点的键等于当前节点的键
			return // 节点已存在，直接返回
		}
	}

	// 将新节点插入到树中
	node.Parent = newNodeParent
	if newNodeParent == rbTree.sentinel { // 如果新节点的父节点是哨兵节点
		rbTree.Root = node // 将新节点设置为根节点
	} else if node.Key < newNodeParent.Key { // 如果新节点的键小于其父节点的键
		newNodeParent.Left = node // 将新节点设置为其父节点的左子节点
	} else { // 如果新节点的键大于其父节点的键
		newNodeParent.Right = node // 将新节点设置为其父节点的右子节点
	}

	// 初始化新节点的左右孩子为哨兵节点，颜色为红色
	node.Left = rbTree.sentinel  // 新节点的左子节点为哨兵节点
	node.Right = rbTree.sentinel // 新节点的右子节点为哨兵节点
	node.Color = __RED           // 新节点的颜色为红色

	// 修复红黑树的颜色和结构
	rbTree.insertFixUp(node) // 调整红黑树以满足其性质
	rbTree.NodeNum++         // 节点数量加一
}

// insertFixUp 修复红黑树的颜色和结构
func (rbTree *RbTree) insertFixUp(node *RbTreeNode) {
	for node != rbTree.Root && node.Parent.Color == __RED { // 当新节点不是根节点且其父节点是红色时，进行调整
		if node.Parent == node.Parent.Parent.Left { // 如果新节点的父节点是祖父节点的左子节点
			uncle := node.Parent.Parent.Right // 获取叔叔节点
			if uncle.Color == __RED {         // 叔叔节点是红色
				node.Parent.Color = __BLACK      // 将父节点颜色设为黑色
				uncle.Color = __BLACK            // 将叔叔节点颜色设为黑色
				node.Parent.Parent.Color = __RED // 将祖父节点颜色设为红色
				node = node.Parent.Parent        // 将当前节点设为祖父节点，继续调整
			} else { // 叔叔节点是黑色
				if node == node.Parent.Right { // 当前节点是父节点的右子节点
					node = node.Parent      // 将当前节点设为父节点
					rbTree.LeftRotate(node) // 对当前节点进行左旋
				}
				node.Parent.Color = __BLACK            // 将父节点颜色设为黑色
				node.Parent.Parent.Color = __RED       // 将祖父节点颜色设为红色
				rbTree.RightRotate(node.Parent.Parent) // 对祖父节点进行右旋
			}
		} else { // 如果新节点的父节点是祖父节点的右子节点
			uncle := node.Parent.Parent.Left // 获取叔叔节点
			if uncle.Color == __RED {        // 叔叔节点是红色
				node.Parent.Color = __BLACK      // 将父节点颜色设为黑色
				uncle.Color = __BLACK            // 将叔叔节点颜色设为黑色
				node.Parent.Parent.Color = __RED // 将祖父节点颜色设为红色
				node = node.Parent.Parent        // 将当前节点设为祖父节点，继续调整
			} else { // 叔叔节点是黑色
				if node == node.Parent.Left { // 当前节点是父节点的左子节点
					node = node.Parent       // 将当前节点设为父节点
					rbTree.RightRotate(node) // 对当前节点进行右旋
				}
				node.Parent.Color = __BLACK           // 将父节点颜色设为黑色
				node.Parent.Parent.Color = __RED      // 将祖父节点颜色设为红色
				rbTree.LeftRotate(node.Parent.Parent) // 对祖父节点进行左旋
			}
		}
	}
	rbTree.Root.Color = __BLACK // 将根节点颜色设为黑色
}

// Delete 从红黑树中删除一个节点
func (rbTree *RbTree) Delete(key RbTreeKeyType) {
	// 查找节点
	node := rbTree.getNode(key)
	if node == rbTree.sentinel {
		return // 如果未找到节点，直接返回
	}

	// 获取替代节点
	willDeletedNode := getReplaceNode(node, rbTree.sentinel)

	// 调整节点关系，将替代节点从树中剥离出来
	willDeletedChildNode := rbTree.sentinel
	if willDeletedNode.Left != rbTree.sentinel {
		willDeletedChildNode = willDeletedNode.Left
	} else if willDeletedNode.Right != rbTree.sentinel {
		willDeletedChildNode = willDeletedNode.Right
	}
	willDeletedChildNode.Parent = willDeletedNode.Parent
	if willDeletedNode.Parent == rbTree.sentinel {
		rbTree.Root = willDeletedChildNode
	} else if willDeletedNode == willDeletedNode.Parent.Left {
		willDeletedNode.Parent.Left = willDeletedChildNode
	} else {
		willDeletedNode.Parent.Right = willDeletedChildNode
	}

	// 如果删除的节点不是替代节点，则用替代节点的值覆盖删除节点的值
	if willDeletedNode != node {
		node.Key = willDeletedNode.Key
		node.Value = willDeletedNode.Value
	}

	// 如果替代节点的颜色为黑色，则修复红黑树的颜色和结构
	if willDeletedNode.Color == __BLACK {
		rbTree.deleteFixUp(willDeletedChildNode)
	}

	// 释放替代节点，帮助GC
	willDeletedNode = nil
	rbTree.NodeNum-- // 节点数量减一
}

// deleteFixUp 修复红黑树的颜色和结构
func (rbTree *RbTree) deleteFixUp(node *RbTreeNode) {
	for node != rbTree.Root && node.Color == __BLACK {
		if node == node.Parent.Left { // 如果当前节点是其父节点的左子节点
			brother := node.Parent.Right // 获取兄弟节点
			if brother.Color == __RED {  // 兄弟节点是红色
				brother.Color = __BLACK        // 将兄弟节点颜色设为黑色
				node.Parent.Color = __RED      // 将父节点颜色设为红色
				rbTree.LeftRotate(node.Parent) // 对父节点进行左旋
				brother = node.Parent.Right    // 更新兄弟节点
			}
			if brother.Left.Color == __BLACK && brother.Right.Color == __BLACK { // 兄弟节点的两个子节点都是黑色
				brother.Color = __RED // 将兄弟节点颜色设为红色
				node = node.Parent    // 将当前节点设为父节点，继续调整
			} else {
				if brother.Right.Color == __BLACK { // 兄弟节点的右子节点是黑色
					brother.Left.Color = __BLACK // 将兄弟节点的左子节点颜色设为黑色
					brother.Color = __RED        // 将兄弟节点颜色设为红色
					rbTree.RightRotate(brother)  // 对兄弟节点进行右旋
					brother = node.Parent.Right  // 更新兄弟节点
				}
				brother.Color = node.Parent.Color // 将兄弟节点颜色设为父节点颜色
				node.Parent.Color = __BLACK       // 将父节点颜色设为黑色
				brother.Right.Color = __BLACK     // 将兄弟节点的右子节点颜色设为黑色
				rbTree.LeftRotate(node.Parent)    // 对父节点进行左旋
				node = rbTree.Root                // 将当前节点设为根节点
			}
		} else { // 如果当前节点是其父节点的右子节点
			brother := node.Parent.Left // 获取兄弟节点
			if brother.Color == __RED { // 兄弟节点是红色
				brother.Color = __BLACK         // 将兄弟节点颜色设为黑色
				node.Parent.Color = __RED       // 将父节点颜色设为红色
				rbTree.RightRotate(node.Parent) // 对父节点进行右旋
				brother = node.Parent.Left      // 更新兄弟节点
			}
			if brother.Left.Color == __BLACK && brother.Right.Color == __BLACK { // 兄弟节点的两个子节点都是黑色
				brother.Color = __RED // 将兄弟节点颜色设为红色
				node = node.Parent    // 将当前节点设为父节点，继续调整
			} else {
				if brother.Left.Color == __BLACK { // 兄弟节点的左子节点是黑色
					brother.Right.Color = __BLACK // 将兄弟节点的右子节点颜色设为黑色
					brother.Color = __RED         // 将兄弟节点颜色设为红色
					rbTree.LeftRotate(brother)    // 对兄弟节点进行左旋
					brother = node.Parent.Left    // 更新兄弟节点
				}
				brother.Color = node.Parent.Color // 将兄弟节点颜色设为父节点颜色
				node.Parent.Color = __BLACK       // 将父节点颜色设为黑色
				brother.Left.Color = __BLACK      // 将兄弟节点的左子节点颜色设为黑色
				rbTree.RightRotate(node.Parent)   // 对父节点进行右旋
				node = rbTree.Root                // 将当前节点设为根节点
			}
		}
	}
	node.Color = __BLACK // 将当前节点颜色设为黑色
}

// getNode 查找具有给定键的节点
func (rbTree *RbTree) getNode(key RbTreeKeyType) *RbTreeNode {
	node := rbTree.Root           // 从根节点开始查找
	for node != rbTree.sentinel { // 遍历树查找目标节点
		if key < node.Key { // 如果目标键小于当前节点的键
			node = node.Left // 向左子树查找
		} else if key > node.Key { // 如果目标键大于当前节点的键
			node = node.Right // 向右子树查找
		} else {
			return node // 找到目标节点
		}
	}
	return rbTree.sentinel // 未找到目标节点，返回哨兵节点
}

// getReplaceNode 获取将要删除的节点的替代者
func getReplaceNode(node, sentinelNode *RbTreeNode) *RbTreeNode {
	if node.Right != sentinelNode { // 如果节点有右子节点
		return node.Right.FindMinNodeBy(sentinelNode) // 返回右子树中的最小节点
	}
	replaceNode := node.Parent                                     // 获取节点的父节点
	for replaceNode != sentinelNode && node == replaceNode.Right { // 如果节点是其父节点的右子节点
		node = replaceNode // 向上遍历树
		replaceNode = replaceNode.Parent
	}
	return replaceNode // 返回替代节点
}

// PrintTree 按层次打印红黑树的结构
func (rbTree *RbTree) PrintTree() {
	printSubTree(rbTree.Root, rbTree.sentinel, 0) // 从根节点开始打印
}

// printSubTree 打印子树的结构
func printSubTree(node, sentinelNode *RbTreeNode, level int) {
	if node != sentinelNode { // 如果节点不是哨兵节点
		printSubTree(node.Right, sentinelNode, level+1) // 先打印右子树
		for i := 0; i < level; i++ {                    // 打印层次
			fmt.Print("    ")
		}
		color := "BLACK"
		if node.Color == __RED {
			color = "RED"
		}
		fmt.Printf("(%s) %d\n", color, node.Key)       // 打印节点
		printSubTree(node.Left, sentinelNode, level+1) // 再打印左子树
	}
}

func main() {
	rbt := NewRbTree()
	rbt.Insert(20, nil)
	rbt.Insert(15, nil)
	rbt.Insert(25, nil)
	rbt.Insert(10, nil)
	rbt.Insert(5, nil)
	rbt.Insert(1, nil)

	rbt.PrintTree()
}
