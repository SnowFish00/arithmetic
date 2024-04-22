package main

import (
	"fmt"
	"strings"
)

type Elem struct {
	next, prev *Elem
	list       *List
	val        interface{}
}

type List struct {
	root Elem
	len  int
}

// 初始化链表
func (l *List) init_list() *List {
	//以root为分割 将root的prev指向表尾 next指向表头 利用这个哨兵节点
	l.root.prev = &l.root
	l.root.next = &l.root
	l.len = 0
	return l
}

// 新建链表
func (l *List) new_list() *List {
	return new(List).init_list()
}

// 返回表长
func (l *List) len_list() int {
	return l.len
}

// 获取链表头部
func (l *List) get_head() *Elem {
	if l.len == 0 {
		return nil
	}

	return l.root.next

}

// 返回链表尾
func (l *List) get_tail() *Elem {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

// 打印全部链表
func (l *List) get_all() {
	if l.len == 0 {
		return
	}

	//从哨兵节点下一个节点开始即头节点
	head := l.root.next
	var builder strings.Builder
	//这个root相当于  尾<-root->头
	for head != &l.root {
		builder.WriteString(fmt.Sprintf("%v->", head.val))
		head = head.next
	}
	fmt.Println(builder.String())
}

// 根据index获取node
func (l *List) get_node_index(index int) *Elem {
	if l.root.next == nil || index < 0 || index > l.len-1 {
		return nil
	}

	node := l.root.next
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node

}

// 根据val获取node
func (l *List) get_node_target(target interface{}) *Elem {
	if l.root.next == nil {
		return nil
	}

	node := l.root.next

	for node != nil {
		if node.val == target {
			return node
		}
		node = node.next
	}

	return nil
}

// 后插入
func (l *List) insert_back(val interface{}, node *Elem) *Elem {
	if l.root.next == nil {
		return nil
	}

	new_node := &Elem{val: val}

	//new_node 的前后插入
	new_node.prev = node
	new_node.next = node.next

	//原结点指向新结点
	node.next.prev = new_node
	node.next = new_node

	//长度++
	l.len++

	//同步list
	node.list = l

	return new_node

}

// 前插入
func (l *List) insert_front(val interface{}, node *Elem) *Elem {
	if l.root.next == nil {
		return nil
	}

	new_node := &Elem{val: val}

	new_node.prev = node.prev
	new_node.next = node

	node.prev.next = new_node
	node.prev = new_node

	l.len++

	return new_node
}

// 删除 信任node未判断node存在性,其实node是传来的所以信任了
func (l *List) delete(node *Elem) bool {
	if l.root.next == nil {
		return false
	}

	//架空node
	node.prev.next = node.next
	node.next.prev = node.prev

	//清空node GC回收 node的prev与next仍然存在 必须指控置空
	node.prev = nil
	node.next = nil
	node.val = nil
	node.list = nil

	//长度--
	l.len--

	return true
}

// 移动 b 到 a后
func (l *List) move(a_node *Elem, b_node *Elem) bool {
	if l.root.next == nil {
		return false
	}

	//b
	b_node.prev.next = b_node.next
	b_node.next.prev = b_node.prev

	//a
	//注意操作时机 先变换指向在改变值 防止值改变指向错误
	b_node.prev = a_node
	b_node.next = a_node.next

	a_node.next.prev = b_node
	a_node.next = b_node

	return true

}

func main() {
	var l List
	l.init_list()

	l.insert_back(3306, &l.root)
	first := l.get_node_target(3306)
	seconde := l.insert_front("不是哥们!", first)
	third := l.insert_back("啊!", seconde)
	fmt.Println(l.move(third, seconde))
	four := l.insert_back("哈哈哈哈", third)
	l.get_all()
	fmt.Println(l.delete(four))
	l.get_all()

	fmt.Println(l.get_head())
	fmt.Println(l.get_tail())
}
