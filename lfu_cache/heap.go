package lfu_cache

type lfuHeapItem struct {
	value interface{}
	freq int
	indexInHeap int
}

type LFUHeapCache struct {
	heap []*lfuHeapItem
	hash map[string]*lfuHeapItem
	cap int
	count int
}

func NewLfuHeapCache (cap int) *LFUHeapCache {
	res := &LFUHeapCache{}
	res.heap = make([]*lfuHeapItem, 0, cap)
	res.hash = make(map[string]*lfuHeapItem, cap)
	res.cap = cap
	return res
}

func (c *LFUHeapCache) Get(key string) (interface{},bool) {
	if res, ok := c.hash[key]; ok  {
		res.freq++
		c.shiftDown(res.indexInHeap)
		return res, true
	}

	return nil, false
}

func (c *LFUHeapCache)Set (key string, value interface{}) {
	if res, ok := c.hash[key]; ok  {
		res.freq++
		res.value = value
		c.shiftDown(res.indexInHeap)
		return
	}
	newItem := &lfuHeapItem{}
	newItem.freq = 1
	newItem.value = value

	// 新建key
	if c.count == c.cap {
		// 超出限制
		c.heap[0] = newItem
		newItem.indexInHeap = 0

	} else {
		// 没超出限制
		c.heap = append(c.heap, newItem)
		newItem.indexInHeap = len(c.heap)-1
		c.shiftUp()
		c.count++
	}
	c.hash[key] = newItem
}
func (c *LFUHeapCache) Del (key string) {
	if _,ok := c.hash[key]; !ok {
		return
	}
	delV := c.hash[key]
	c.heap[len(c.heap)-1].indexInHeap = delV.indexInHeap
	c.heap[len(c.heap)-1], c.heap[delV.indexInHeap] = c.heap[delV.indexInHeap],c.heap[len(c.heap)-1]
	c.heap = c.heap[0:len(c.heap)-1]

	c.shiftDown(delV.indexInHeap)

	delete(c.hash, key)
	c.count--
}

func (c *LFUHeapCache) shiftDown (index int) {
	for index < len(c.heap) {
		minChildIndex := 0
		if 2 * index + 2 < len(c.heap) {
			if c.heap[2 * index + 2].freq < c.heap[2 * index + 1].freq {
				minChildIndex = 2 * index + 2
			} else {
				minChildIndex = 2 * index + 1
			}
		} else if 2 * index + 1 < len(c.heap) {
			minChildIndex = 2*index + 1
		}
		if minChildIndex > 0 && c.heap[index].freq > c.heap[minChildIndex].freq {
			c.heap[index].indexInHeap = minChildIndex
			c.heap[minChildIndex].indexInHeap = index
			c.heap[index], c.heap[minChildIndex] = c.heap[minChildIndex], c.heap[index]
			index = minChildIndex
		} else {
			break
		}
	}
}

func (c *LFUHeapCache) shiftUp () {
	childIndex := len(c.heap)-1
	parentIndex :=  (childIndex-1)/2

	for childIndex > 0 && c.heap[parentIndex].freq > c.heap[childIndex].freq {
		c.heap[childIndex].indexInHeap = parentIndex
		c.heap[parentIndex].indexInHeap = childIndex
		c.heap[childIndex], c.heap[parentIndex] = c.heap[parentIndex], c.heap[childIndex]
		childIndex = parentIndex
		parentIndex = (childIndex-1)/2
	}
}
