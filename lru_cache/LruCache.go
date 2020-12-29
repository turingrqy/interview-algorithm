package lru_cache

import "renqiyang/interview/list"
//只有双向链表的插入和删除是o(1) 单俩表头尾插是o(1)中间插是o(n) 因为需要找
type LruCache struct {
	hashMap map[int]*list.ListNode2
	head *list.ListNode2
	tail *list.ListNode2
	cap int
	size int
}

func Construct(cap int) *LruCache {
	return &LruCache{
		make(map[int]*list.ListNode2,cap),
		nil,
		nil,
		cap,
		0,
	}
}

func (self *LruCache) Get(k int) int {
	if node,ok := self.hashMap[k];ok {
		if self.size == 1 || node == self.tail {
			return node.Value
		}

		if self.head == node {
			self.head = self.head.Next
			self.head.Prev = nil
			self.tail.Next = node
			node.Prev = self.tail
			self.tail = node
			self.tail.Next = nil

		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			self.tail.Next = node
			node.Prev = self.tail
			self.tail = node
			self.tail.Next = nil
		}

		return node.Value
	}
	return -1
}

func(self *LruCache)  Put (k,store int) {
	if node,ok := self.hashMap[k];ok {
		node.Value = store
		if self.head == node {
			self.head = self.head.Next
			self.head.Prev = nil
			self.tail.Next = node
			node.Prev = self.tail
			self.tail = node
			self.tail.Next = nil

		} else {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			self.tail.Next = node
			node.Prev = self.tail
			self.tail = node
			self.tail.Next = nil
		}
	} else {
		newItem := &list.ListNode2{
			k,
			store,
			nil,
			nil,
		}
		self.tail.Next = newItem
		newItem.Prev = self.tail
		if self.size == self.cap {
			delete(self.hashMap,self.head.Key)
			self.head = self.head.Next
			self.head.Prev = nil
		} else {

			self.size++
		}
	}
}

