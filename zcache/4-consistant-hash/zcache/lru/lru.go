package lru

import "container/list"

type Cache struct {
	maxBytes  int64                         // 缓存运行使用的最大内存，单位为字节
	nbytes    int64                         // 当前缓存使用的内存，单位为字节
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      //字典
	OnEvicted func(key string, value Value) // 某条记录被删除时触发的回调函数
}

// 双向链表节点的数据类型
type entry struct {
	key   string
	value Value
}

// Value 值可以为任何实现了Value接口的数据类型
type Value interface {
	Len() int //返回值所占用的内存大小
}

// New 初始化缓存
func New(maBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get 获取缓存
func (c *Cache) Get(key string) (value Value, ok bool) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		return element.Value.(*entry).value, true
	}
	return
}

// RemoveOldest 删除最近最近未使用缓存
func (c *Cache) RemoveOldest() {
	point := c.ll.Back()
	if point != nil {
		c.ll.Remove(point)
		kv := point.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add 新增缓存
func (c *Cache) Add(key string, value Value) {
	if element, ok := c.cache[key]; ok {
		c.ll.MoveToFront(element)
		kv := element.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.nbytes > c.maxBytes {
		c.RemoveOldest()
	}
}

// Len 获取缓存数据量
func (c *Cache) Len() int {
	return c.ll.Len()
}
