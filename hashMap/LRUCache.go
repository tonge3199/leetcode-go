package hashMap

/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。 实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
*/
import "container/list"

// Node represents key-value pair stored in list element
type Node struct {
	key   int
	value int
}

// make function as a method of an exported type struct( uppercase letter name)
type LRUCache struct {
	capacity int
	find     map[int]*list.Element
	cache    *list.List
}

func New(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity,
		make(map[int]*list.Element, capacity),
		new(list.List),
	}

	return lru
}

// List func
// Init() *List
// New() *List
// MoveToFront(e *Element)
// MoveToBack(e *Element)
// PushFront(v any)
//
// Element func
// Next() *Element
// Prev() *Element

func (lru *LRUCache) get(key int) int {
	e := lru.find[key]
	if e != nil {
		// recent used key, MoveToFront
		lru.cache.MoveToFront(e)
		// Extract value from Node
		node := e.Value.(*Node)
		return node.value
	}

	return -1
}

func (lru *LRUCache) put(key, value int) {
	if e := lru.find[key]; e != nil {
		// Update existing key: update value and move to front
		node := e.Value.(*Node)
		node.value = value
		lru.cache.MoveToFront(e)
	} else {
		// Add new key-value pair
		node := &Node{key: key, value: value}

		// Check capacity before adding
		if lru.cache.Len() >= lru.capacity {
			// Remove least recently used (back of list)
			lastElement := lru.cache.Back()
			if lastElement != nil {
				lru.cache.Remove(lastElement)
				// Remove from map using key stored in node
				lastNode := lastElement.Value.(*Node)
				delete(lru.find, lastNode.key)
			}
		}

		// Add new element to front and update map
		element := lru.cache.PushFront(node)
		lru.find[key] = element
	}
}
