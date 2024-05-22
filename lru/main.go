package main

import (
	"container/list"
	"fmt"
)

func mainA() {
	fmt.Println("enter test lru")

	lru := InitLRU(3)
	lru.Add("1")
	lru.Add("2")
	lru.Add("3")
	lru.Add("5")
	lru.Print()
	fmt.Println(lru.Get("2"))
	lru.Print()
}

type LRU struct {
	limit    int
	eList    *list.List
	elements map[string]*list.Element
}

func InitLRU(size int) *LRU {
	return &LRU{
		limit:    size,
		eList:    list.New(),
		elements: make(map[string]*list.Element, 0),
	}
}

func (l *LRU) Add(key string) {
	if elem, ok := l.elements[key]; ok {
		l.eList.MoveToFront(elem)
		return
	}

	elem := l.eList.PushFront(key)
	l.elements[key] = elem

	if l.eList.Len() > l.limit {
		l.removeBackEle()
	}
}

func (l *LRU) removeBackEle() {
	eleBack := l.eList.Back()
	if eleBack != nil {
		l.eList.Remove(eleBack)
	}
}

func (l *LRU) Get(key string) any {
	elem, ok := l.elements[key]
	if !ok {
		return ""
	}
	l.eList.MoveToFront(elem)
	return elem.Value
}

func (l *LRU) Print() {
	head := l.eList.Front()
	result := make([]any, 0)
	for head != nil {
		result = append(result, head.Value)
		head = head.Next()
	}
	fmt.Println("result", result)
}
