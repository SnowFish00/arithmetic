package main

import (
	"fmt"
)

// 定义哈希表中的链表节点
type HashNode struct {
	key   string
	value int
	next  *HashNode
}

// 定义哈希表
type HashTable struct {
	bucketSize int
	buckets    []*HashNode
}

// 创建哈希表
func NewHashTable(size int) *HashTable {
	return &HashTable{
		bucketSize: size,
		buckets:    make([]*HashNode, size),
	}
}

// 哈希函数 质数处理 取模对应bucket
func (ht *HashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = 31*hash + int(char)
	}
	return hash % ht.bucketSize
}

// 插入键值对 (链表法解决冲突)
func (ht *HashTable) Insert(key string, value int) {
	index := ht.hash(key)
	newNode := &HashNode{
		key:   key,
		value: value,
		next:  nil,
	}

	// 如果当前桶为空，直接插入
	if ht.buckets[index] == nil {
		ht.buckets[index] = newNode
	} else {
		// 如果不为空，遍历链表
		current := ht.buckets[index]
		for current != nil {
			// 如果键已经存在，更新值
			if current.key == key {
				current.value = value
				return
			}
			// 移动到下一个节点
			if current.next == nil {
				break
			}
			current = current.next
		}
		// 插入新节点到链表末尾
		current.next = newNode
	}
}

// 查找键对应的值
func (ht *HashTable) Search(key string) (int, bool) {
	index := ht.hash(key)
	current := ht.buckets[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

// 删除键值对
func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)
	current := ht.buckets[index]
	var prev *HashNode = nil
	for current != nil {
		if current.key == key {
			if prev == nil {
				// 删除的是链表头节点
				ht.buckets[index] = current.next
			} else {
				// 删除的是链表中间的节点
				prev.next = current.next
			}
			return true
		}
		prev = current
		current = current.next
	}
	return false
}

func main() {
	hashTable := NewHashTable(10)

	// 插入键值对
	hashTable.Insert("apple", 5)
	hashTable.Insert("banana", 6)
	hashTable.Insert("orange", 7)

	// 查找键对应的值
	if value, found := hashTable.Search("apple"); found {
		fmt.Printf("apple: %d\n", value)
	}

	// 删除键值对
	if deleted := hashTable.Delete("banana"); deleted {
		fmt.Println("banana deleted")
	}
}
