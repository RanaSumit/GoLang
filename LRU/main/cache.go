package main

import "container/list"
const CACHE_SIZE int = 3
type Key interface{}

type entry struct {
	key Key
	value int
}

type Cache struct {
	max int
	l *list.List
	cache map[interface{}]*list.Element
}

func New(num int) *Cache {
	return &Cache{
		max : num,
		l : list.New(),
		cache : make(map[interface{}]*list.Element),
	}
}

func (c *Cache) Set (key Key, value int) {
	if e, ok := c.cache[key]; ok {
		c.l.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}

	ele := c.l.PushFront(&entry{key, value})
	c.cache[key] = ele
	if c.max != 0 && c.l.Len() > c.max {
		c.RemoveOld()
	}
}

func (c *Cache) Get(key Key) (value int) {
	if c.cache == nil {
		return
	}

	if ele, ok := c.cache[key]; ok {
		c.l.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return

}
func (c *Cache) RemoveOld() {
	if c.cache == nil {
		return
	}

	e := c.l.Back()
	c.removeElement(e)
}

func (c *Cache) removeElement(e *list.Element) {
	c.l.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cache, kv.key)
}








