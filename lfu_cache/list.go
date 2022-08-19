package lfu_cache

type cacheItem struct {
	value interface{}
	prev *cacheItem
	next *cacheItem
	freq *freqListItem
}

type freqListItem struct {
	freq int
	next *freqListItem
	prev *freqListItem
	head *cacheItem
	tail *cacheItem
}

type LfuListCache struct {
	hash map[string]*cacheItem
	freqHead *freqListItem
	freqTail *freqListItem
	count int
	cap int
}

func NewLfuListCache (cap int) *LfuListCache {
	res := &LfuListCache{}
	res.cap = cap
	res.hash = make(map[string]*cacheItem)
	res.freqHead = &freqListItem{}
	res.freqTail = &freqListItem{}
	res.freqHead.next = res.freqTail
	res.freqTail.prev = res.freqHead
	return res
}

func (c *LfuListCache) Get (key string) (interface{}, bool) {
	if res, ok := c.hash[key]; ok {
		c.access(res)
		return res.value, true
	}

	return nil, false
}

func (c *LfuListCache) Del (key string) {
	if _, ok := c.hash[key]; !ok {
		return
	}

	res := c.hash[key]

	c.del(res)

	delete(c.hash, key)
	c.count--
}

func (c *LfuListCache) Set (key string, value interface{}) {
	if res, ok := c.hash[key]; ok {
		res.value = value
		c.access(res)
		return
	}
	//新增
	if c.count == c.cap {
		// 满了要淘汰
		c.del(c.freqHead.next.tail.prev)
	}
	node :=  &cacheItem{}
	node.value = value
	var newFreqItem *freqListItem

	if c.freqHead.next.freq == 1 {
		newFreqItem = c.freqHead.next
	} else {
		newFreqItem = &freqListItem{}
		newFreqItem.freq = 1
		newFreqItem.head = &cacheItem{}
		newFreqItem.tail = &cacheItem{}

		newFreqItem.head.next = newFreqItem.tail
		newFreqItem.tail.prev = newFreqItem.head

		newFreqItem.next = c.freqHead.next
		newFreqItem.prev = c.freqHead
		c.freqHead.next.prev = newFreqItem
		c.freqHead.next = newFreqItem
	}

	node.next = newFreqItem.head.next
	node.prev = newFreqItem.head

	newFreqItem.head.next = node
	node.next.prev = node
	node.freq = newFreqItem
	c.hash[key] = node
}

func (c *LfuListCache) del (node *cacheItem) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil
	if node.freq.head.next == node.freq.tail {
		node.freq.prev.next = node.freq.next
		node.freq.next.prev = node.freq.prev

		node.freq.next = nil
		node.freq.prev = nil
	}
}

func (c *LfuListCache)access (node *cacheItem) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil

	newFreq := node.freq.freq+1
	var newFreqItem *freqListItem
	if node.freq.next.freq != newFreq {
		//newFreq 节点不存在
		newFreqItem = &freqListItem{}
		newFreqItem.freq = newFreq
		newFreqItem.next = node.freq.next
		newFreqItem.prev = node.freq

		newFreqItem.head = &cacheItem{}
		newFreqItem.tail = &cacheItem{}

		newFreqItem.head.next = newFreqItem.tail
		newFreqItem.tail.prev = newFreqItem.head

		node.freq.next.prev = newFreqItem
		node.freq.next = newFreqItem
	} else {
		newFreqItem = node.freq.next
	}
	node.next = newFreqItem.head.next
	node.prev = newFreqItem.head

	newFreqItem.head.next = node
	node.next.prev = node

	//当前频次的列表空了 释放freq 节点
	if node.freq.head.next == node.freq.tail {
		node.freq.prev.next = node.freq.next
		node.freq.next.prev = node.freq.prev

		node.freq.next = nil
		node.freq.prev = nil
	}
	node.freq = newFreqItem
}


