package main

import "fmt"

/*
双向链表
add() – 从头插入一个数据结点
push() – 从尾插入一个数据结点
del() – 删除一个结点（因为需要比较，所以使用了 compareable 的泛型）
print() – 从头遍历一个链表，并输出值。
   ----------------------------------
   || prev  | data  |  next ||    || prev | data  | next  ||
   ----------------------|----------------->-^
*/

// 定义 node 结构体 当前数据 前一个node 后一个node
type node[T comparable] struct {
	data T
	prev *node[T] // 指向前一个节点
	next *node[T] // 指向后一个节点
}

type list[T comparable] struct {
	head, tail *node[T]
	len        int
}

// 头和尾都是 nil
func (l *list[T]) isEmpty() bool {
	return l.head == nil && l.tail == nil
}

func (l *list[T]) add(data T) {
	n := &node[T]{ // 为什么是 &
		data: data,
		prev: nil,
		next: l.head,
	}
	// 长度为1
	if l.isEmpty() {
		l.head = n
		l.tail = n
	}
	l.head.prev = n
	l.head = n
}

func (l *list[T]) push(data T) {
	n := &node[T]{
		data: data,
		prev: l.tail,
		next: nil,
	}
	if l.isEmpty() {
		l.head = n
		l.tail = n
	}
	l.tail.next = n
	l.tail = n
}
func (l *list[T]) del(data T) {
	for p := l.head; p != nil; p = p.next {
		if data == p.data {

			if p == l.head {
				l.head = p.next
			}
			if p == l.tail {
				l.tail = p.prev
			}
			if p.prev != nil {
				p.prev.next = p.next
			}
			if p.next != nil {
				p.next.prev = p.prev
			}
			return
		}
	}
}
func (l *list[T]) print() {
	if l.isEmpty() {
		fmt.Println("the link list is empty.")
		return
	}
	for p := l.head; p != nil; p = p.next {
		fmt.Printf("[%v] -> ", p.data)
	}
	fmt.Println("nil")
}

func main() {
	var l = list[int]{}
	l.add(1)
	l.add(2)
	l.push(3)
	l.push(4)
	l.add(5)
	l.print() //[5] -> [2] -> [1] -> [3] -> [4] -> nil
	l.del(5)
	l.del(1)
	l.del(4)
	l.print() //[2] -> [3] -> nil
}
