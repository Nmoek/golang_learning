/*
 * @file: list.go
 * @brief: 练习Go泛型-结构体泛型
 * @author: Kewin Li
 * @date: 2023年9月30日
 */

package struct_generics

import (
	"fmt"
)

type LinkList[T any] interface {
	AddToTail(val T)     // 将值追加在尾部
	AddToHead(val T)     // 将值追加在头部
	Add(idx int, val T)  // 指定位置追加值
	Delete(idx int) bool // 指定位置删除值
	Count() int          // 返回当前结点数
	Show()               // 展示当前所有结点
}

type myListV1[T any] struct {
	head  *Node[T]
	tail  *Node[T]
	count int
}

func (l *myListV1[T]) AddToTail(val T) {
	node := Node[T]{data: val, next: nil, pre: nil}

	if l.count == 0 {
		l.head = &node
		l.tail = &node
	} else {
		l.tail.next = &node
		node.pre = l.tail
		l.tail = &node
	}
	l.count++
}

func (l *myListV1[T]) AddToHead(val T) {
	node := Node[T]{data: val, next: nil, pre: nil}

	if l.count == 0 {
		l.head = &node
		l.tail = &node
	} else {
		node.next = l.head
		l.head.pre = &node
		l.head = &node
	}
	l.count++
}

func (l *myListV1[T]) Count() int {
	return l.count
}

func (l *myListV1[T]) Delte(idx int) bool {
	if idx < 0 || idx >= l.count {
		fmt.Printf("idx %d out of range! \n", idx)
		return false
	}

	tmp := l.head
	for i := 0; i < idx; i++ {
		tmp = tmp.next
	}

	if tmp == nil {
		return false
	}

	if idx == 0 {
		l.head = tmp.next
	} else if idx == l.count-1 {
		l.tail = tmp.pre
	} else {

		tmp.pre.next = tmp.next
		tmp.next.pre = tmp.pre
		if idx == l.count-1 {
			l.tail = tmp.pre
		}
	}
	l.count--

	fmt.Printf("[delete] idx:%d, val:%v \n", idx, tmp.data)
	return true

}

func (l *myListV1[T]) Show() {

	for first := l.head; first != nil; first = first.next {
		fmt.Printf("%v", first.data)
		fmt.Printf("->")
	}
	fmt.Printf("nil \n")

}

type Node[T any] struct {
	data T
	next *Node[T]
	pre  *Node[T]
}

func UseList() {

	l1 := myListV1[int]{}
	l1.AddToTail(1)
	l1.AddToTail(2)
	l1.AddToTail(3)
	l1.AddToTail(4)
	l1.AddToHead(5)
	l1.Show()
	fmt.Printf("count: %d \n", l1.Count())
	if !l1.Delte(3) {
		fmt.Printf("delete err! \n")
	}
	l1.Show()
	if !l1.Delte(0) {
		fmt.Printf("delete err! \n")
	}
	l1.Show()
	if !l1.Delte(0) {
		fmt.Printf("delete err! \n")
	}
	l1.Show()
}
