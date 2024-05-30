package main

import (
	"fmt"
	"sync"
)

// BPItem 代表B+树中的一个键值对
type BPItem struct {
	Key int64
	Val interface{}
}

// BPNode 代表B+树中的一个节点
type BPNode struct {
	MaxKey int64
	Nodes  []*BPNode
	Items  []BPItem
	Next   *BPNode
}

// BPTree 代表整个B+树结构
type BPTree struct {
	mutex sync.RWMutex
	root  *BPNode
	width int
	halfw int
}

// NewBPTree 创建一个新的B+树
func NewBPTree(width int) *BPTree {
	if width < 3 {
		width = 3
	}
	bt := &BPTree{}
	bt.root = NewLeafNode(width)
	bt.width = width
	bt.halfw = (bt.width + 1) / 2
	return bt
}

// NewLeafNode 创建一个新的叶子节点
func NewLeafNode(width int) *BPNode {
	node := &BPNode{}
	node.Items = make([]BPItem, width+1)
	node.Items = node.Items[0:0]
	return node
}

// Get 从B+树中获取一个键对应的值
func (t *BPTree) Get(key int64) interface{} {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	node := t.root
	// 查找叶子节点
	for i := 0; i < len(node.Nodes); i++ {
		if key <= node.Nodes[i].MaxKey {
			node = node.Nodes[i]
			i = 0
		}
	}

	// 在叶子节点中查找键值对
	if len(node.Nodes) == 0 {
		for i := 0; i < len(node.Items); i++ {
			if node.Items[i].Key == key {
				return node.Items[i].Val
			}
		}
	}
	return nil
}

// Set 在B+树中设置一个键值对
func (t *BPTree) Set(key int64, value interface{}) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.setValue(nil, t.root, key, value)
}

// setValue 是一个辅助函数，用于递归设置键值对
func (t *BPTree) setValue(parent *BPNode, node *BPNode, key int64, value interface{}) {
	// 如果当前节点有子节点，则继续递归
	for i := 0; i < len(node.Nodes); i++ {
		if key <= node.Nodes[i].MaxKey || i == len(node.Nodes)-1 {
			t.setValue(node, node.Nodes[i], key, value)
			break
		}
	}

	// 如果是叶子节点，直接添加或更新数据
	if len(node.Nodes) == 0 {
		node.setValue(key, value)
	}

	// 如果节点需要分裂，则分裂节点
	node_new := t.splitNode(node)
	if node_new != nil {
		// 如果父节点不存在，创建一个新的父节点
		if parent == nil {
			parent = NewIndexNode(t.width)
			parent.addChild(node)
			t.root = parent
		}
		// 将新节点添加到父节点
		parent.addChild(node_new)
	}
}

// setValue 在节点中设置一个键值对
func (node *BPNode) setValue(key int64, value interface{}) {
	item := BPItem{key, value}
	num := len(node.Items)
	// 直接添加到节点的适当位置
	if num < 1 {
		node.Items = append(node.Items, item)
		node.MaxKey = item.Key
	} else if key < node.Items[0].Key {
		node.Items = append([]BPItem{item}, node.Items...)
	} else if key > node.Items[num-1].Key {
		node.Items = append(node.Items, item)
		node.MaxKey = item.Key
	} else {
		for i := 0; i < num; i++ {
			if node.Items[i].Key > key {
				node.Items = append(node.Items, BPItem{})
				copy(node.Items[i+1:], node.Items[i:])
				node.Items[i] = item
				return
			} else if node.Items[i].Key == key {
				node.Items[i] = item
				return
			}
		}
	}
}

// splitNode 分裂一个节点
func (t *BPTree) splitNode(node *BPNode) *BPNode {
	if len(node.Nodes) > t.width {
		// 创建新节点
		halfw := t.width / 2
		node2 := NewIndexNode(t.width)
		node2.Nodes = append(node2.Nodes, node.Nodes[halfw:len(node.Nodes)]...)
		node2.MaxKey = node2.Nodes[len(node2.Nodes)-1].MaxKey

		// 修改原节点数据
		node.Nodes = node.Nodes[0:halfw]
		node.MaxKey = node.Nodes[len(node.Nodes)-1].MaxKey

		return node2
	} else if len(node.Items) > t.width {
		// 创建新节点
		halfw := (t.width + 1) / 2
		node2 := NewLeafNode(t.width)
		node2.Items = append(node2.Items, node.Items[halfw:len(node.Items)]...)
		node2.MaxKey = node2.Items[len(node2.Items)-1].Key

		// 修改原节点数据
		node.Items = node.Items[0:halfw]
		node.MaxKey = node.Items[len(node.Items)-1].Key
		node.Next = node2

		return node2
	}

	return nil
}

// NewIndexNode 创建一个新的非叶子节点
func NewIndexNode(width int) *BPNode {
	return &BPNode{
		Nodes: make([]*BPNode, width+1),
	}
}

// addChild 将一个子节点添加到父节点
func (node *BPNode) addChild(child *BPNode) {
	node.Nodes = append(node.Nodes, child)
}

// deleteChild 从父节点中删除一个子节点
func (node *BPNode) deleteChild(child *BPNode) {
	for i, n := range node.Nodes {
		if n == child {
			node.Nodes = append(node.Nodes[:i], node.Nodes[i+1:]...)
			break
		}
	}
}

// Remove 从B+树中删除一个键值对
func (t *BPTree) Remove(key int64) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.deleteItem(nil, t.root, key)
}

// deleteItem 是一个辅助函数，用于递归删除键值对
func (t *BPTree) deleteItem(parent *BPNode, node *BPNode, key int64) {
	for i := 0; i < len(node.Nodes); i++ {
		if key <= node.Nodes[i].MaxKey {
			t.deleteItem(node, node.Nodes[i], key)
			break
		}
	}

	if len(node.Nodes) == 0 {
		// 删除记录后若节点的子项<t.halfw，则从兄弟节点移动记录，或者合并节点
		if node.deleteItem(key) && len(node.Items) < t.halfw {
			t.itemMoveOrMerge(parent, node)
		}
	} else {
		// 若节点的子项<t.halfw，则从兄弟节点移动记录，或者合并节点
		node.MaxKey = node.Nodes[len(node.Nodes)-1].MaxKey
		if len(node.Nodes) < t.halfw {
			t.childMoveOrMerge(parent, node)
		}
	}
}

// deleteItem 在节点中删除一个键值对
func (node *BPNode) deleteItem(key int64) bool {
	for i, item := range node.Items {
		if item.Key == key {
			node.Items = append(node.Items[:i], node.Items[i+1:]...)
			return true
		}
	}
	return false
}

// itemMoveOrMerge 在节点子项数量不足时，从兄弟节点移动记录或合并节点
func (t *BPTree) itemMoveOrMerge(parent *BPNode, node *BPNode) {
	// 获取兄弟节点
	var node1, node2 *BPNode
	for i, n := range parent.Nodes {
		if n == node {
			if i > 0 {
				node1 = parent.Nodes[i-1]
			}
			if i < len(parent.Nodes)-1 {
				node2 = parent.Nodes[i+1]
			}
			break
		}
	}

	// 尝试从左侧兄弟节点移动记录
	if node1 != nil && len(node1.Items) > t.halfw {
		item := node1.Items[len(node1.Items)-1]
		node1.Items = node1.Items[:len(node1.Items)-1]
		node.Items = append([]BPItem{item}, node.Items...)
		parent.MaxKey = parent.Nodes[len(parent.Nodes)-1].MaxKey
		return
	}

	// 尝试从右侧兄弟节点移动记录
	if node2 != nil && len(node2.Items) > t.halfw {
		item := node2.Items[0]
		node2.Items = node2.Items[1:]
		node.Items = append(node.Items, item)
		parent.MaxKey = parent.Nodes[len(parent.Nodes)-1].MaxKey
		return
	}

	// 与左侧兄弟节点合并
	if node1 != nil {
		node.Items = append(node.Items, node1.Items...)
		parent.deleteChild(node1)
	}

	// 与右侧兄弟节点合并
	if node2 != nil {
		node.Items = append(node.Items, node2.Items...)
		parent.deleteChild(node2)
	}
}

// childMoveOrMerge 在节点子节点数量不足时，从兄弟节点移动子节点或合并子节点
func (t *BPTree) childMoveOrMerge(parent *BPNode, node *BPNode) {
	// 获取兄弟节点
	var node1, node2 *BPNode
	for i, n := range parent.Nodes {
		if n == node {
			if i > 0 {
				node1 = parent.Nodes[i-1]
			}
			if i < len(parent.Nodes)-1 {
				node2 = parent.Nodes[i+1]
			}
			break
		}
	}

	// 尝试从左侧兄弟节点移动子节点
	if node1 != nil && len(node1.Nodes) > t.halfw {
		child := node1.Nodes[len(node1.Nodes)-1]
		node1.Nodes = node1.Nodes[:len(node1.Nodes)-1]
		node.Nodes = append(node.Nodes, child)
		parent.MaxKey = parent.Nodes[len(parent.Nodes)-1].MaxKey
		return
	}

	// 尝试从右侧兄弟节点移动子节点
	if node2 != nil && len(node2.Nodes) > t.halfw {
		child := node2.Nodes[0]
		node2.Nodes = node2.Nodes[1:]
		node.Nodes = append(node.Nodes, child)
		parent.MaxKey = parent.Nodes[len(parent.Nodes)-1].MaxKey
		return
	}

	// 与左侧兄弟节点合并
	if node1 != nil {
		node.Nodes = append(node.Nodes, node1.Nodes...)
		parent.deleteChild(node1)
	}

	// 与右侧兄弟节点合并
	if node2 != nil {
		node.Nodes = append(node.Nodes, node2.Nodes...)
		parent.deleteChild(node2)
	}
}

// main 函数用于演示B+树的基本操作
func main() {
	tree := NewBPTree(3)

	// 添加键值对
	tree.Set(1, "value1")
	tree.Set(2, "value2")
	tree.Set(3, "value3")
	tree.Set(4, "value4")
	tree.Set(5, "value5")

	// 获取键值对
	fmt.Println("Get 1:", tree.Get(1)) // 应输出 "value1"
	fmt.Println("Get 2:", tree.Get(2)) // 应输出 "value2"
	fmt.Println("Get 3:", tree.Get(3)) // 应输出 "value3"
	fmt.Println("Get 4:", tree.Get(4)) // 应输出 "value4"
	fmt.Println("Get 5:", tree.Get(5)) // 应输出 "value5"

	// 删除键值对
	tree.Remove(3)

	// 再次获取键值对
	fmt.Println("Get 3 (after remove):", tree.Get(3)) // 应输出 nil
}
