package integration

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

type LRUCache struct {
	cap int

	elements map[int]*list.Element // 记录了 key 和 element 的对应关系，element 中存储的是下面的 pair 结构
	l        *list.List            // 按时间顺序存储了所有的 element
}

// 在进行删除操作时，需要维护两个数据结构：删除双向链表中要淘汰的 value；删除 map 中要淘汰的 value 所对应的 key
// 如果双向链表的 Value 中仅存储 value 不存储 key，那么删除 map 中的 key 时还需要遍历 map，找到对应的 key 再删除，时间复杂度就是 O(1) 了
type pair struct {
	k, v int
}

func LRUCacheConstructor(capacity int) LRUCache {
	return LRUCache{
		cap:      capacity,
		elements: make(map[int]*list.Element),
		l:        list.New(),
	}
}

func (cache *LRUCache) Get(key int) int {
	if element, exists := cache.elements[key]; exists {
		// 需要将找到的数据放到链表头部
		cache.l.MoveToFront(element)
		return element.Value.(pair).v
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if element, exists := cache.elements[key]; exists {
		element.Value = pair{k: key, v: value}
		cache.l.MoveToFront(element)
	} else {
		element := cache.l.PushFront(pair{k: key, v: value})
		cache.elements[key] = element
	}
	if cache.l.Len() > cache.cap {
		element := cache.l.Back()
		cache.l.Remove(element)
		delete(cache.elements, element.Value.(pair).k)
	}
}

// for test!
func (cache *LRUCache) String() string {
	values := make([]string, 0, cache.cap)
	cur := cache.l.Front()
	for cur != nil {
		values = append(values, strconv.Itoa(cur.Value.(pair).v))
		cur = cur.Next()
	}
	return strings.Join(values, " -> ")
}

func TestLRUCache(t *testing.T) {
	cache := LRUCacheConstructor(2)
	cache.Put(1, 1)
	assert.Equal(t, "1", cache.String())
	cache.Put(2, 2)
	assert.Equal(t, "2 -> 1", cache.String())
	assert.Equal(t, 1, cache.Get(1))
	assert.Equal(t, "1 -> 2", cache.String())

	cache.Put(3, 3)
	assert.Equal(t, "3 -> 1", cache.String())
	assert.Equal(t, -1, cache.Get(2))
	cache.Put(4, 4)
	assert.Equal(t, "4 -> 3", cache.String())
	assert.Equal(t, -1, cache.Get(1))
	assert.Equal(t, 3, cache.Get(3))
	assert.Equal(t, "3 -> 4", cache.String())
	assert.Equal(t, 4, cache.Get(4))
	assert.Equal(t, "4 -> 3", cache.String())
}
