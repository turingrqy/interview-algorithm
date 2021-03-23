package limit

import (
	"sync"
	"time"
	"container/ring"
)

// 令牌桶算法
type TicketBucket struct {
	StartTime time.Time
	Cap int64
	Interval int64 //ns
	QuotaNum int64
	AvaliableNum int64
	LastTick int64
	Lock sync.Mutex
}

func (self *TicketBucket) Take (count int64) int64 {
	self.Lock.Lock()
	self.AdjustNow(time.Now())
	if self.AvaliableNum < 0 {
		return 0
	}
	if count > self.AvaliableNum {
		self.AvaliableNum = 0
		return self.AvaliableNum
	}
	self.AvaliableNum -= count
	self.Lock.Unlock()
	return count
}

func (self *TicketBucket) AdjustNow (now time.Time)  {
	currentTick := int64(now.Sub(self.StartTime))/self.Interval
	self.AvaliableNum += (currentTick-self.LastTick)*self.QuotaNum
	if self.AvaliableNum > self.Cap {
		self.AvaliableNum = self.Cap
	}
	self.LastTick = currentTick
}
//滑动窗口限流
type SlideWindow1 struct {
	queue [] int64
	lock sync.Mutex
	limit int
	windowTime int64
}

func (self *SlideWindow1) Init(windowTime int64,limit int) {
	self.queue = []int64 {}
	self.windowTime = windowTime * 1000
	self.limit = limit
	self.lock = sync.Mutex{}
}

func (self *SlideWindow1) Check () bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	timenow := time.Now().UnixNano() / 1e6
	if len(self.queue) <= self.limit  {
		self.queue = append(self.queue,timenow)
		return true
	}

	if timenow-self.queue[0] <= timenow {
		return false
	} else {
		self.queue = self.queue[1:]
		self.queue = append(self.queue, timenow)
		return true
	}
}

type SlideWindow2 struct {
	slideTime int
	windowTime int
	bucketNum int
	head *ring.Ring
	limit int
	curCount int
	lock sync.Mutex
}

func (self *SlideWindow2) Init (slideTime int, windowTime int,limit int) {
	self.curCount = 0
	self.limit = limit
	self.windowTime = windowTime
	self.slideTime = slideTime
	self.bucketNum = self.windowTime/self.slideTime
	self.head = ring.New(self.bucketNum)
	for i:= 0;i<self.bucketNum; i++ {
		var tmp int = 0
		self.head.Value = tmp
		self.head = self.head.Next()
	}
	go func () {

		ticker := time.NewTicker(time.Duration(self.slideTime) * time.Millisecond)
		for range ticker.C {
			self.lock.Lock()
			self.head = self.head.Next()
			itemCount := self.head.Value.(int)
			self.limit -= itemCount
			self.head.Value = 0
			self.lock.Unlock()
		}

	}()
}

func (self *SlideWindow2) Check () bool {
	self.lock.Lock()
	defer self.lock.Unlock()
	if self.curCount >= self.limit {
		return false
	}
	self.curCount += 1
	tmpCount := self.head.Value.(int)
	self.head.Value = tmpCount +1
	return true
}


