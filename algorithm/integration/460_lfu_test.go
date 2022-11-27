package integration

import "container/list"

type LFUCache struct {
	cap      int
	min      int                   // 记录了最小频次
	elements map[int]*list.Element // 记录了 key 和 element 的对应关系，element 中存储的为下面的 node 结构
	lists    map[int]*list.List    // 记录了 frequency 和 list 的对应关系，list 中的 elements 以时间顺序排序
}

type node struct {
	k, v      int
	frequency int
}

func LFUCacheConstructor(capacity int) LFUCache {
	return LFUCache{
		cap:      capacity,
		min:      0,
		lists:    make(map[int]*list.List),
		elements: make(map[int]*list.Element),
	}
}

func (cache *LFUCache) Get(key int) int {
	oldElement, exists := cache.elements[key]
	if !exists {
		return -1
	}
	curNode := oldElement.Value.(*node)
	// 从原频次 list 中移除元素
	cache.lists[curNode.frequency].Remove(oldElement)
	// 元素访问频次更新
	curNode.frequency++
	// 判断最新频次的 list 是否存在，不存在则需要初始化
	if _, exists := cache.lists[curNode.frequency]; !exists {
		cache.lists[curNode.frequency] = list.New()
	}
	// 将 node 插入到最新频次对应的 list 头部
	l := cache.lists[curNode.frequency]
	newElement := l.PushFront(curNode)
	// 更新 key 对应的 element 地址
	cache.elements[key] = newElement
	// 更新 min
	if curNode.frequency-1 == cache.min && cache.lists[curNode.frequency-1].Len() == 0 {
		cache.min++
	}
	return curNode.v
}

func (cache *LFUCache) Put(key int, value int) {
	if cache.cap == 0 {
		return
	}
	// 如果存在，则更新值和访问频次并返回
	if element, exists := cache.elements[key]; exists {
		node := element.Value.(*node)
		node.v = value
		cache.Get(key)
		return
	}
	// 如果不存在，且缓存满了，删除最小频次 list 的最早的元素
	if cache.cap == len(cache.elements) {
		minList := cache.lists[cache.min]
		backElement := minList.Back()
		delete(cache.elements, backElement.Value.(*node).k)
		minList.Remove(backElement)
	}
	cache.min = 1
	curNode := &node{
		k:         key,
		v:         value,
		frequency: 1,
	}
	if _, exists := cache.lists[1]; !exists {
		cache.lists[1] = list.New()
	}
	l := cache.lists[1]
	newElement := l.PushFront(curNode)
	cache.elements[key] = newElement
}
