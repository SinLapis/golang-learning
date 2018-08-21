package main

import (
	"container/list"
	"fmt"
)

func linkedList1() {
	l := new(list.List)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Printf("%d\n", i.Value)
	}
}

type DoubleLinkedNode struct {
	Value int
	Next  *DoubleLinkedNode
	Priv  *DoubleLinkedNode
}

type DoubleLinkeList struct {
	cursor *DoubleLinkedNode
}

func (l *DoubleLinkeList) Push(value int) {
	newNode := new(DoubleLinkedNode)
	newNode.Value = value
	if l.cursor == nil {
		l.cursor = newNode
	} else {
		newNode.Priv = l.cursor
		l.cursor.Next = newNode
		l.cursor = l.cursor.Next
	}
}
func (l *DoubleLinkeList) Front() (firstNode *DoubleLinkedNode) {
	for firstNode = l.cursor; firstNode.Priv != nil; firstNode = firstNode.Priv {
	}
	return
}
func linkedList2() {
	l := new(DoubleLinkeList)
	l.Push(1)
	l.Push(2)
	l.Push(4)
	for i := l.Front(); i != nil; i = i.Next {
		fmt.Printf("%d\n", i.Value)
	}
}
func main() {
	linkedList1()
	linkedList2()
}
