package main

import (
	"fmt"
	"strings"
)

type Ring struct {
	prev, next *Ring
	val        interface{}
}

// 初始化
func (r *Ring) init_ring() *Ring {
	r.prev = r
	r.next = r
	return r
}

// 下一结点
func (r *Ring) Next() *Ring {
	if r == nil {
		return r.init_ring()
	}

	return r.next
}

// 上一结点
func (r *Ring) Prev() *Ring {
	if r == nil {
		return r.init_ring()
	}

	return r.prev
}

// 当前指针的前/后n位的结点
func (r *Ring) move_to(n int) *Ring {
	if r == nil {
		return r.init_ring()
	}

	switch {
	//操作改变了r的指针位置，但是不耽误循环链表的相互地址引用所以原循环链表并未改变结构
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}

	return r
}

// 创建n位的结点
func create_nodes(n int) *Ring {
	if n <= 0 {
		return nil
	}

	//头节点 实际上循环链表谁都是头和尾
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		//&Ring{prev: p}相当于prev 而后 p指向下一个位置坐实了 prev: p
		p.next = &Ring{prev: p}
		p = p.next
	}

	//首尾闭环
	r.prev = p
	p.next = r

	return r
}

// 合并两个结点所代表的循环链表
// 将r看作尾 将s看作头
func (r *Ring) link(s *Ring) *Ring {
	//r的下一结点 以及 看作 r所代表的链表的"头"
	n := r.Next()
	if s != nil {
		//p的上一个结点 以及 p的 "尾"
		p := s.Prev()

		//在结点本身上相互 引用
		r.next = s
		s.prev = r

		//看作头尾相互引用
		n.prev = p
		p.next = n

	}

	//返回的n为废弃元素的头结点 因为 n := r.Next()
	//在作为 结点时返回的是node 即废弃的头节点 作为头结点 使其n.prev头为 p尾又链接起来了链表
	return n
}

// 切割再合并实现删除部分n个元素的link 假设信任非空
func (r *Ring) unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}

	//从n+1位置丢弃
	return r.link(r.move_to(n + 1))
}

// rings长度
func (r *Ring) len_ring() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.next; p != r; p = p.next {
			n++
		}
	}

	return n
}

// 打印循环链表元素
func (r *Ring) list_ring() {
	var builder strings.Builder
	if r != nil {
		builder.WriteString(fmt.Sprintf("->%v->", r.val))
		for p := r.next; p != r; p = p.next {
			builder.WriteString(fmt.Sprintf("%v->", p.val))
		}
		builder.WriteString(fmt.Sprintf("%v<-", r.val))

		fmt.Println(builder.String())
	}

}

// (拓展)功能通用Do
func (r *Ring) Do(f func(any)) {
	if r != nil {
		f(r.val)
		for p := r.Next(); p != r; p = p.next {
			f(p.val)
		}
	}
}

func main() {
	ring := create_nodes(5)

	len := ring.len_ring()
	fmt.Println("环形链表中的元素个数:", len)

	for i := 0; i < len; i++ {
		ring.val = i
		ring = ring.Next()
	}

	ring.list_ring()

	fmt.Println(ring.val)
	fmt.Println(ring.move_to(3).val)
	/*TODO 通过副本指针(函数传入的指针是副本)进行操作，只要不修改指针本身，不会影响到原始指针所指向的内容 即不取它的成员赋值，
	所以move_to没有改变ring 当然用ring = 接一下就可以了*/
	fmt.Println(ring.val)
	ring.list_ring()

	//删除node的废弃 链表
	ring.unlink(2).list_ring()
	//删除后的链表
	ring.list_ring()

	//---

	ring_b := create_nodes(6)

	len_b := ring_b.len_ring()
	fmt.Println("环形链表中的元素个数:", len_b)

	for i := 999; i > 999-len_b; i-- {
		ring_b.val = i
		ring_b = ring_b.Prev()
	}

	ring_b.list_ring()

	ring.link(ring_b).list_ring()
}
