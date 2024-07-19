package cache

import (
	"container/list"
	"errors"
	"sync"
)

type Cache struct {
	capacity            int
	evictionOrderedList *list.List
	storage             map[string]*list.Element
	mu                  sync.Mutex
}

type entry struct {
	key   string
	value interface{}
}

func NewCache(capacity int) (*Cache, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than zero")
	}

	cache := &Cache{
		capacity:            capacity,
		evictionOrderedList: list.New(),
		storage:             make(map[string]*list.Element),
	}
	return cache, nil
}

func (cache *Cache) Get(key string) (interface{}, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if element, ok := cache.storage[key]; ok {
		cache.evictionOrderedList.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return nil, false
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if element, ok := cache.storage[key]; ok {
		cache.evictionOrderedList.MoveToFront(element)
		element.Value.(*entry).value = value
		return
	}

	element := cache.evictionOrderedList.PushFront(&entry{key, value})
	cache.storage[key] = element

	if cache.evictionOrderedList.Len() > cache.capacity {
		cache.evict()
	}
}

func (cache *Cache) evict() {
	element := cache.evictionOrderedList.Back()
	if element != nil {
		cache.delete(element)
	}
}

func (cache *Cache) delete(element *list.Element) {
	cache.evictionOrderedList.Remove(element)
	delete(cache.storage, element.Value.(*entry).key)
}

func (cache *Cache) Remove(key string) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if element, ok := cache.storage[key]; ok {
		cache.delete(element)
		return true
	}
	return false
}
