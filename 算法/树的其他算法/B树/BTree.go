package main

import (
	"fmt"
)

const order = 3 // B树的阶

type Node struct {
	Keys     []int
	Children []*Node
	IsLeaf   bool
}

type BTree struct {
	Root *Node
}

func NewBTree() *BTree {
	return &BTree{
		Root: &Node{
			IsLeaf: true,
		},
	}
}

// 打印B树
func (t *BTree) Print() {
	t.printNode(t.Root, "")
}

func (t *BTree) printNode(node *Node, prefix string) {
	fmt.Printf("%s", prefix)
	for i := 0; i < len(node.Keys); i++ {
		fmt.Printf("%d ", node.Keys[i])
	}
	fmt.Println()
	for i := 0; i < len(node.Children); i++ {
		if node.Children[i] != nil {
			t.printNode(node.Children[i], prefix+"  ")
		}
	}
}

// 查找节点
func (t *BTree) Search(key int, node *Node) (*Node, int) {
	i := 0
	for i < len(node.Keys) && key > node.Keys[i] {
		i++
	}
	if i < len(node.Keys) && key == node.Keys[i] {
		return node, i
	}
	if node.IsLeaf {
		return nil, -1
	}
	return t.Search(key, node.Children[i])
}

// 分裂节点
func (t *BTree) splitChild(node, child *Node, index int) {
	newNode := &Node{
		Keys:     make([]int, order-1),
		Children: make([]*Node, order),
		IsLeaf:   child.IsLeaf,
	}

	// 将子节点中间的键值放到父节点中
	node.Keys = append(node.Keys, 0)
	copy(node.Keys[index+1:], node.Keys[index:])
	node.Keys[index] = child.Keys[order/2-1]

	// 将子节点右侧的键值和子节点移动到新节点
	copy(newNode.Keys, child.Keys[order/2:])
	child.Keys = child.Keys[:order/2-1]

	if !child.IsLeaf {
		copy(newNode.Children, child.Children[order/2:])
		child.Children = child.Children[:order/2]
	}

	// 将新节点作为子节点添加到父节点
	node.Children = append(node.Children, nil)
	copy(node.Children[index+2:], node.Children[index+1:])
	node.Children[index+1] = newNode
}

// 插入键值
func (t *BTree) Insert(key int) {
	root := t.Root
	if len(root.Keys) == 2*order-1 {
		newRoot := &Node{
			Keys:     []int{root.Keys[order-1]},
			Children: []*Node{root},
			IsLeaf:   false,
		}
		t.splitChild(newRoot, root, 0)
		t.Root = newRoot
	}
	t.insertNonFull(t.Root, key)
}

// 在非满节点中插入键值
func (t *BTree) insertNonFull(node *Node, key int) {
	i := len(node.Keys) - 1
	if node.IsLeaf {
		node.Keys = append(node.Keys, 0)
		for i >= 0 && key < node.Keys[i] {
			node.Keys[i+1] = node.Keys[i]
			i--
		}
		node.Keys[i+1] = key
	} else {
		for i >= 0 && key < node.Keys[i] {
			i--
		}
		i++
		if len(node.Children[i].Keys) == 2*order-1 {
			t.splitChild(node, node.Children[i], i)
			if key > node.Keys[i] {
				i++
			}
		}
		t.insertNonFull(node.Children[i], key)
	}
}

// 删除键值
func (t *BTree) Delete(key int) {
	t.deleteKey(t.Root, key)
	if len(t.Root.Keys) == 0 && !t.Root.IsLeaf {
		t.Root = t.Root.Children[0]
	}
}

// 删除节点中的键值
func (t *BTree) deleteKey(node *Node, key int) {
	i := 0
	for i < len(node.Keys) && key > node.Keys[i] {
		i++
	}
	if node.IsLeaf {
		if i < len(node.Keys) && key == node.Keys[i] {
			copy(node.Keys[i:], node.Keys[i+1:])
			node.Keys = node.Keys[:len(node.Keys)-1]
		}
	} else {
		if i < len(node.Keys) && key == node.Keys[i] {
			t.deleteInternalNode(node, i)
		} else {
			child := node.Children[i]
			if len(child.Keys) >= order {
				t.deleteKey(child, key)
			} else {
				t.fixChild(node, i)
				t.deleteKey(child, key)
			}
		}
	}
}

// 删除内部节点
func (t *BTree) deleteInternalNode(node *Node, index int) {
	key := node.Keys[index]

	// 情况1: 如果左子节点有足够的键值
	if len(node.Children[index].Keys) >= order {
		pred := t.getPredecessor(node, index)
		node.Keys[index] = pred
		t.deleteKey(node.Children[index], pred)
	} else if len(node.Children[index+1].Keys) >= order {
		// 情况2: 如果右子节点有足够的键值
		succ := t.getSuccessor(node, index)
		node.Keys[index] = succ
		t.deleteKey(node.Children[index+1], succ)
	} else {
		// 情况3: 如果左右子节点都只有最小数目的键值
		t.mergeChildren(node, index)
		t.deleteKey(node.Children[index], key)
	}
}

// 获取前驱节点
func (t *BTree) getPredecessor(node *Node, index int) int {
	child := node.Children[index]
	for !child.IsLeaf {
		child = child.Children[len(child.Keys)]
	}
	return child.Keys[len(child.Keys)-1]
}

// 获取后继节点
func (t *BTree) getSuccessor(node *Node, index int) int {
	child := node.Children[index+1]
	for !child.IsLeaf {
		child = child.Children[0]
	}
	return child.Keys[0]
}

// 合并子节点
func (t *BTree) mergeChildren(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index+1]

	child.Keys = append(child.Keys, node.Keys[index])
	child.Keys = append(child.Keys, sibling.Keys...)
	if !child.IsLeaf {
		child.Children = append(child.Children, sibling.Children...)
	}

	node.Keys = append(node.Keys[:index], node.Keys[index+1:]...)
	node.Children = append(node.Children[:index+1], node.Children[index+2:]...)
}

// 修复子节点
func (t *BTree) fixChild(node *Node, index int) {
	if index > 0 && len(node.Children[index-1].Keys) >= order {
		t.borrowFromLeft(node, index)
	} else if index < len(node.Keys) && len(node.Children[index+1].Keys) >= order {
		t.borrowFromRight(node, index)
	} else if index > 0 {
		t.mergeChildren(node, index-1)
	} else {
		t.mergeChildren(node, index)
	}
}

// 从左兄弟借用键值
func (t *BTree) borrowFromLeft(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index-1]

	child.Keys = append([]int{node.Keys[index-1]}, child.Keys...)
	if !child.IsLeaf {
		child.Children = append([]*Node{sibling.Children[len(sibling.Keys)]}, child.Children...)
	}
	node.Keys[index-1] = sibling.Keys[len(sibling.Keys)-1]
	sibling.Keys = sibling.Keys[:len(sibling.Keys)-1]
}

// 从右兄弟借用键值
func (t *BTree) borrowFromRight(node *Node, index int) {
	child := node.Children[index]
	sibling := node.Children[index+1]

	child.Keys = append(child.Keys, node.Keys[index])
	if !child.IsLeaf {
		child.Children = append(child.Children, sibling.Children[0])
	}
	node.Keys[index] = sibling.Keys[0]
	copy(sibling.Keys, sibling.Keys[1:])
	sibling.Keys = sibling.Keys[:len(sibling.Keys)-1]
	if !sibling.IsLeaf {
		copy(sibling.Children, sibling.Children[1:])
		sibling.Children = sibling.Children[:len(sibling.Children)-1]
	}
}

// 主函数
func main() {
	btree := NewBTree()

	// 插入键值
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}
	for _, key := range keys {
		btree.Insert(key)
		fmt.Printf("After inserting %d:\n", key)
		btree.Print()
		fmt.Println()
	}

	// 删除键值
	deleteKeys := []int{5, 12, 17, 20}
	for _, key := range deleteKeys {
		btree.Delete(key)
		fmt.Printf("After deleting %d:\n", key)
		btree.Print()
		fmt.Println()
	}

	// 查找键值
	findKeys := []int{10, 30, 7, 100}
	for _, key := range findKeys {
		node, index := btree.Search(key, btree.Root)
		fmt.Printf("Search for %d: ", key)
		if index != -1 {
			fmt.Printf("Found at index %d in node %+v\n", index, node)
		} else {
			fmt.Println("Not found")
		}
	}
}
