package main

import "fmt"

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func initDlinkNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkNode{},
		capacity: capacity,
		size:     0,
		head:     initDlinkNode(0, 0),
		tail:     initDlinkNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := initDlinkNode(key, value)
		this.cache[key] = newNode
		this.addToHead(newNode)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
			this.size--
		}
	}

}

// addToHead
// removeNode
// moveToHead
// removeTail

func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	// test with show
	lru := Constructor(2)
	lru.Put(1, 10)
	lru.Put(2, 20)
	lru.Put(3, 30)
	lru.Put(5, 50)
	fmt.Println(lru.Get(2))
	fmt.Println(lru.Get(1))
	fmt.Println(lru.Get(3))
	fmt.Println(lru.Get(5))

}
