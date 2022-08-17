package cache

import (
	"fmt"
	"strings"
)

// Cache - Struct that implements LRU Cache
type Cache struct {
	capacity   int
	size       int
	head, tail *Node
	index      map[string]*Node
}

func NewCache(capacity int) *Cache {
	head := NewNode().Set("head", true)
	tail := NewNode().Set("tail", true).SetPrevious(head)
	head = head.SetNext(tail)

	return &Cache{
		capacity: capacity,
		size:     0,
		head:     head,
		tail:     tail,
		index:    make(map[string]*Node),
	}
}

// Cache.Cache() - Prints the cache data
func (cache *Cache) Cache() string {
	data := fmt.Sprintf("{\"capacity\": %d, \"size\": %d, \"isFull\": %v, data:[", cache.capacity, cache.size, cache.IsFull())
	node := cache.head.next
	for node.next != nil {
		data += fmt.Sprintf("%s,", node.Node())
		node = node.next
	}
	data = strings.TrimRight(data, ", ")
	return fmt.Sprintf("%s]}", data)
}

func (cache *Cache) Capacity() int {
	return cache.capacity
}

func (cache *Cache) Delete(key string) {
	node, ok := cache.index[key]
	if !ok {
		return
	}
	cache.remove(node)
}

func (cache *Cache) remove(node *Node) { //TODO: return error
	if node == nil || node.IsHead() || node.IsTail() {
		return
	}

	if node.previous != nil {
		node.previous.SetNext(node.next)
	}

	if node.next != nil {
		node.next.SetPrevious(node.previous)
	}

	delete(cache.index, node.Key)
	if cache.size > 0 {
		cache.size -= 1
	}
}

func (cache *Cache) Get(key string) (interface{}, bool) {
	node, ok := cache.index[key]
	if !ok {
		return nil, false
	}
	cache.remove(node) // remove from old position
	newNode := node.Clone().SetPrevious(cache.head).SetNext(cache.head.next)
	cache.head.next.SetPrevious(newNode)
	cache.head.SetNext(newNode)
	cache.size += 1

	return newNode.Value, true
}

func (cache *Cache) IsFull() bool {
	return !(cache.size < cache.capacity)
}

// Cache.Purge() - removes all data in the cache
func (cache *Cache) Purge() {
	cache.index = make(map[string]*Node)
	cache.head.SetNext(cache.tail)
	cache.tail.SetPrevious(cache.head)
	cache.size = 0
}

func (cache *Cache) Put(key string, value interface{}) {
	if cache.IsFull() {
		cache.remove(cache.tail.previous)
	}

	node, ok := cache.index[key]
	if ok {
		cache.remove(node)
	}

	newNode := NewNode().
		Set(key, value).
		SetPrevious(cache.head).
		SetNext(cache.head.next)

	cache.head.next.SetPrevious(newNode)
	cache.head.SetNext(newNode)

	cache.index[key] = newNode
	cache.size += 1
}

func (cache *Cache) Size() int {
	return cache.size
}
