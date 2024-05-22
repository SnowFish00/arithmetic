package main

import (
	"fmt"
)

// 数据结构之线性表--顺序表
type SqListInterface interface {
	// 基本操作
	NewSeqList(capacity int) *SqList // 初始化
	// InitList(capacity int)           // 初始化
	ListEmpty() bool // 判空
	ListFul() bool   // 判满
	ListLength() int // 返回数据元素个数
	ClearList()      // 清空
	DestroyList()    // 销毁
	// 元素操作
	ListInsert(index int, elem interface{}) bool // 插入元素
	ListDelete(index int) bool                   // 删除元素
	GetElem(index int) (interface{}, bool)       // 获取元素
	SetElem(elem interface{}, index int) bool    // 更新元素
	LocateELem(elem interface{}) (int, bool)     // 返回第1个值与elem相同的元素的位置若这样的数据元素不存在,则返回值为0
	// 其他操作
	PriorElem(elem interface{}) (interface{}, bool) // 寻找元素的前驱（当前元素的前一个元素）
	NextElem(elem interface{}) (interface{}, bool)  // 寻找元素的后驱（当前元素的后一个元素）
	TraverseList()                                  // 遍历
	Pop() interface{}                               // 从末尾弹出一个元素
	Append(elem interface{}) bool                   // 从末尾插入一个元素
	ExtendCapacity()                                // 扩容
	Reserve()                                       // 反转
}

// SqList 顺序表的结构类型为SqList
// 使用golang语言的interface接口类型创建顺序表
type SqList struct {
	Len         int           // 线性表长度
	Capacity    int           // 表容量
	Data        []interface{} // 指向线性表空间
	ExtendRatio int           // 每次列表扩容的倍数
}

/*基本操作*/

func (l *SqList) InitSeqList(capacity int) {
	l.Len = 0
	l.Capacity = capacity
	l.Data = make([]interface{}, capacity)
	l.ExtendRatio = 2
}

func (l *SqList) ListEmpty() bool {
	return l.Len == 0
}

func (l *SqList) ListFul() bool {
	return l.Len == l.Capacity
}

func (l *SqList) ListLength() int {
	return l.Len
}

func (l *SqList) ClearList() {
	l.Len = 0
	l.Data = nil
}

func (l *SqList) DestroyList() {
	l.Len = 0
	l.Capacity = 0
	l.Data = nil
	l.ExtendRatio = 0
}

/*元素操作*/

func (l *SqList) ListInsert(index int, elem interface{}) bool {
	//负数 满 过容
	if index < 0 || index > l.Capacity || l.ListFul() {
		return false
	}

	//若初始化中无数据那么i+1 与 i 比较 比 i 与 i-1 更好防止-1
	for i := l.Len - 1; i > index; i-- {
		l.Data[i+1] = l.Data[i]
	}

	l.Data[index] = elem
	//长度加一
	l.Len++

	return true
}

func (l *SqList) ListDelete(index int) bool {
	//负数 满 过容
	if index < 0 || index > l.Capacity || l.ListEmpty() {
		return false
	}

	for i := index; i < l.Len-1; i++ {
		l.Data[i] = l.Data[i+1]
	}

	// l.Data = l.Data[:l.Len-1]
	//置空
	l.Data[l.Len-1] = nil
	l.Len--

	return true
}

func (l *SqList) SetElem(index int, elem interface{}) bool {
	//负数 过长
	if index < 0 || index >= l.Len {
		return false
	}

	l.Data[index] = elem
	return true
}

func (l *SqList) GetElem(index int) (interface{}, bool) {
	//负数 过长
	if index < 0 || index >= l.Len {
		return nil, false
	}

	return l.Data[index], true
}

func (l *SqList) LocateELem(elem interface{}) (int, bool) {
	for i, it := range l.Data {
		if elem == it {
			return i, true
		}
	}
	return -1, false
}

func (l *SqList) PriorElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateELem(elem)
	if i == 0 || i == -1 {
		return nil, false
	}

	return l.Data[i-1], true
}

func (l *SqList) NextElem(elem interface{}) (interface{}, bool) {
	i, _ := l.LocateELem(elem)
	if i == l.Len-1 || i == -1 {
		return nil, false
	}
	return l.Data[i+1], true
}

func (l *SqList) TraverseList() {
	for _, it := range l.Data {
		fmt.Println(it)
	}
}

func (l *SqList) Pop() interface{} {
	//空则不出
	if l.ListEmpty() {
		return nil
	}

	result := l.Data[l.Len-1]
	l.Data = l.Data[:l.Len-1]
	l.Len--

	return result
}

func (l *SqList) Append(elem interface{}) bool {
	if l.ListFul() {
		return false
	}

	l.Data = append(l.Data, elem)
	l.Len++
	return true
}

func (l *SqList) ExtendCapacity() {
	l.Data = append(l.Data, make([]interface{}, l.Capacity*(l.ExtendRatio-1))...)
	l.Capacity = len(l.Data)
}

// 反转
func (l *SqList) Reserve() {
	for i := 0; i < l.Len/2; i++ {
		//关于中间值对称的两个值相互交换位置
		tmp := l.Data[l.Len-i-1]
		l.Data[l.Len-i-1] = l.Data[i]
		l.Data[i] = tmp
	}
}

func main() {
	var sq SqList
	//初始化
	sq.InitSeqList(3)
	fmt.Println(sq)
	//判空
	fmt.Println(sq.ListEmpty())
	//判满
	fmt.Println(sq.ListFul())
	//返回元素个数
	fmt.Println(sq.ListLength())

	//插入
	fmt.Println(sq.ListInsert(0, "hello1"))
	fmt.Println(sq.ListInsert(1, "hello2"))
	fmt.Println(sq.ListInsert(2, "hello3"))
	fmt.Println(sq.ListInsert(3, "hello4"))
	fmt.Println(sq)

	//删除
	fmt.Println(sq.ListDelete(0))
	fmt.Println(sq)

	//位置获取
	fmt.Println(sq.GetElem(0))

	//更新
	fmt.Println(sq.SetElem(0, 8848))
	fmt.Println(sq)

	//元素获取
	fmt.Println(sq.LocateELem(8848))

	//前驱后继
	fmt.Println(sq.PriorElem("hello3"))
	fmt.Println(sq.NextElem("hello3"))

	//遍历
	sq.TraverseList()

	//反转
	sq.Reserve()
	fmt.Println(sq)

	//末尾弹出
	fmt.Println(sq.Pop())
	fmt.Println(sq)

	//末尾插入
	fmt.Println(sq.Append("over!"))
	fmt.Println(sq)

	//清空list
	sq.ClearList()
	//销毁
	sq.DestroyList()
}
