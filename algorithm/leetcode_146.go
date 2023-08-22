package main

/*
*146. LRU缓存机制
 */
type LRUCache struct {
	size, capacity           int
	cache                    map[int]*MyListNode
	protectHead, protectTail *MyListNode
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		size:        0,
		capacity:    capacity,
		cache:       map[int]*MyListNode{},
		protectHead: &MyListNode{},
		protectTail: &MyListNode{},
	}
	lruCache.protectHead.next = lruCache.protectTail
	lruCache.protectTail.pre = lruCache.protectHead
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if listNode, ok := this.cache[key]; ok {
		this.moveToTail(listNode)
		return listNode.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 存在则返回
	if listNode, ok := this.cache[key]; ok {
		listNode.value = value
		this.moveToTail(listNode)
		return
	}
	if this.size == this.capacity {
		headNode := this.protectHead.next
		this.remove(this.protectHead)
		delete(this.cache, headNode.key)
		this.size--
	}
	listNode := &MyListNode{
		key:   key,
		value: value,
	}
	this.addToTail(listNode)
	this.cache[key] = listNode
	this.size++
}

type MyListNode struct {
	key, value int
	pre, next  *MyListNode
}

// 删除指定节点
func (this *LRUCache) remove(listNode *MyListNode) {
	listNode.next.pre = listNode.pre
	listNode.pre.next = listNode.next

	listNode.pre = nil
	listNode.next = nil
}

// protectTail是一个空指针 只是为了控制更好的塞入节点
func (this *LRUCache) addToTail(listNode *MyListNode) {
	this.protectTail.pre.next = listNode
	listNode.pre = this.protectTail.pre

	listNode.next = this.protectTail
	this.protectTail.pre = listNode
}

func (this *LRUCache) moveToTail(listNode *MyListNode) {
	this.remove(listNode)
	this.addToTail(listNode)
}
