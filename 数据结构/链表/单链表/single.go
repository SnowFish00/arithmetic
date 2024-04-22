package main

import (
	"fmt"
	"strings"
)

type LNode struct {
	Val  interface{}
	Next *LNode
}

func (l *LNode) InitLnode(val int) {
	l.Val = val
	l.Next = nil
}
func (l *LNode) NewNode(val int) *LNode {
	return &LNode{
		Val:  val,
		Next: nil,
	}
}

func (l *LNode) Insert(node *LNode, p *LNode) {
	//当前node节点的next 即 n1
	n1 := node.Next
	//node 的下一位改为p
	node.Next = p
	//p指向原来的n1
	p.Next = n1
}

func (l *LNode) Delete(node *LNode) {
	if node.Next == nil {
		return
	}

	//连续取两次到n2
	n1 := node.Next
	n2 := n1.Next

	//n0指向n2
	node.Next = n2

}

func (l *LNode) IndexSearch(head *LNode, index int) *LNode {
	//搜索下一个自己是空的不行,删除下一个 next是空的不行
	if head == nil {
		return nil
	}

	for i := 0; i < index; i++ {
		head = head.Next
		if head == nil {
			//提前nil了未找到
			return nil
		}
	}

	return head
}

func (l *LNode) TargetSearch(head *LNode, target int) int {
	index := 0
	// if head == nil{}
	//使用head 判断而不是 head.next 毕竟第一个节点要先判断nil 要写也行记得判断前面 head == nil
	//写代码时应该把自己放在中间节点既要考虑前也要考虑后最后 考虑前 和 后
	for head != nil {
		if head.Val == target {
			return index
		}
		index++
		head = head.Next
	}

	return -1
}

func (l *LNode) ShowNodes(head *LNode) {
	var builder strings.Builder

	for head != nil {
		builder.WriteString(fmt.Sprintf("%v", head.Val) + "->")
		head = head.Next
	}

	fmt.Println(builder.String())

}

func main() {
	//初始化首结点
	var head LNode
	head.InitLnode(9527)

	//新建节点
	n1 := head.NewNode(1)
	n2 := head.NewNode(2)
	n3 := head.NewNode(3)
	n4 := head.NewNode(4)
	n5 := head.NewNode(5)

	//构建链表
	head.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	//打印
	head.ShowNodes(&head)

	//插入
	n6 := head.NewNode(8848)
	head.Insert(&head, n6)
	head.ShowNodes(&head)

	//删除
	head.Delete(head.Next)
	head.ShowNodes(&head)

	//index查询
	fmt.Println(head.IndexSearch(&head, 0))

	//target寻找
	fmt.Println(head.TargetSearch(&head, 8848))
}
