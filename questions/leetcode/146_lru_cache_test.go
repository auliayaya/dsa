package leetcode

import (
	"container/list"
	"fmt"
	"testing"
)

type LRUCache struct {
	cache    map[int]*list.Element //define double linked list and hashmap
	lruCache *list.List
	capacity int
}
type Entry struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		lruCache: list.New(),
		cache:    make(map[int]*list.Element),
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if element, found := this.cache[key]; found {
		this.lruCache.MoveToFront(element)
		return element.Value.(*Entry).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, found := this.cache[key]; found {
		entry := element.Value.(*Entry)
		fmt.Printf("Updating key %d: Old Value: %d → New Value: %d\n", key, entry.value, value)
		entry.value = value
		this.lruCache.MoveToFront(element)
	} else {
		if len(this.cache) >= this.capacity {
			lruElement := this.lruCache.Back()
			if lruElement != nil {
				lruEntry := lruElement.Value.(*Entry)
				fmt.Printf("Evicting LRU key %d with value %d\n", lruEntry.key, lruEntry.value)
				delete(this.cache, lruEntry.key)
				this.lruCache.Remove(lruElement)
			}
		}
		fmt.Printf("Inserting key %d with value %d\n", key, value)
		newElement := this.lruCache.PushFront(&Entry{key, value})
		this.cache[key] = newElement
	}
}
func (c *LRUCache) IterateCache() {
	fmt.Println("Cache contents (Most → Least Recently Used):")
	for e := c.lruCache.Front(); e != nil; e = e.Next() {
		entry := e.Value.(*Entry)
		fmt.Printf("Key: %d, Value: %d\n", entry.key, entry.value)
	}
}
func (c *LRUCache) IterateCacheReverse() {
	fmt.Println("Cache contents (Least → Most Recently Used):")
	for e := c.lruCache.Back(); e != nil; e = e.Prev() {
		entry := e.Value.(*Entry)
		fmt.Printf("Key: %d, Value: %d\n", entry.key, entry.value)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestLRUCache(t *testing.T) {
	lruCache := Constructor(3)
	lruCache.Put(1, 1)
	lruCache.Put(2, 2)
	lruCache.Put(3, 3)
	lruCache.Put(4, 4)
	lruCache.Put(5, 5)
	lruCache.IterateCache()
	lruCache.Get(3)
	lruCache.Get(4)
	lruCache.IterateCache()
	lruCache.Put(4, 40)
	lruCache.IterateCache()
	lruCache.Put(3, 30)
	lruCache.IterateCache()
	lruCache.Get(5)
	lruCache.Put(6, 6)
	lruCache.IterateCache()
}
