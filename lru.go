package cache

import (
	"container/list"
	"sync"
)

type (
	item struct {
		Key   string
		Value interface{}
	}

	Lru struct {
		queue    *list.List
		mutex    *sync.RWMutex
		items    map[string]*list.Element
		capacity int
	}
)

func NewLru(capacity int) *Lru {
	return &Lru{
		queue:    list.New(),
		mutex:    &sync.RWMutex{},
		items:    make(map[string]*list.Element),
		capacity: capacity,
	}
}

func (c *Lru) Set(key string, value interface{}) bool {
	c.mutex.Lock()
	if element, exists := c.items[key]; exists == true {
		c.queue.MoveToFront(element)
		element.Value.(*item).Value = value
		return true
	}
	c.mutex.Unlock()

	if c.queue.Len() == c.capacity {
		c.purge()
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	item := &item{
		Key:   key,
		Value: value,
	}

	element := c.queue.PushFront(item)
	c.items[item.Key] = element

	return true
}

func (c *Lru) purge() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if element := c.queue.Back(); element != nil {
		item := c.queue.Remove(element).(*item)
		delete(c.items, item.Key)
	}
}

func (c *Lru) Get(key string) interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	element, exists := c.items[key]
	if exists == false {
		return nil
	}
	c.queue.MoveToFront(element)

	return element.Value.(*item).Value
}
