package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity, NewList(), make(map[Key]*ListItem, capacity), sync.Mutex{}}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	listElem, ok := c.items[key]
	if ok {
		cacheItemElem, _ := listElem.Value.(*cacheItem)
		cacheItemElem.value = value
		c.queue.MoveToFront(listElem)
	} else {
		newCacheElem := &cacheItem{key: key, value: value}
		newListElem := c.queue.PushFront(newCacheElem)
		c.items[key] = newListElem
		if c.queue.Len() > c.capacity {
			lastListElem := c.queue.Back()
			lastCacheElem, _ := lastListElem.Value.(*cacheItem)
			c.queue.Remove(lastListElem)
			delete(c.items, lastCacheElem.key)
		}
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	listElem, ok := c.items[key]
	if ok {
		cacheElem, _ := listElem.Value.(*cacheItem)
		c.queue.MoveToFront(listElem)
		return cacheElem.value, ok
	}
	return nil, ok
}

func (c *lruCache) Clear() {
	c.Lock()
	defer c.Unlock()

	for key := range c.items {
		c.queue.Remove(c.items[key])
		delete(c.items, key)
	}
}
