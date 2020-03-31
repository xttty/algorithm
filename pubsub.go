package algorithm

import (
	"sync"
	"time"
)

// Subcrib  订阅者
type Subcrib chan interface{}
type filterFunc func(v interface{}) bool

// Publisher 发布者
type Publisher struct {
	mu        sync.RWMutex
	subs      map[Subcrib]filterFunc
	cacheSize int
	timeout   time.Duration
}

// NewPublisher 新建发布者
func NewPublisher(size int, timeout time.Duration) *Publisher {
	pub := &Publisher{
		subs:      make(map[Subcrib]filterFunc),
		cacheSize: size,
		timeout:   timeout,
	}
	return pub
}

// Subcrib 订阅
func (pub *Publisher) Subcrib(filter filterFunc) Subcrib {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	sub := make(Subcrib, pub.cacheSize)
	pub.subs[sub] = filter
	return sub
}

// Publish 发布消息
func (pub *Publisher) Publish(v interface{}) {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	var wg sync.WaitGroup
	for sub, filter := range pub.subs {
		if filter != nil && !filter(v) {
			continue
		}
		wg.Add(1)
		go func(s Subcrib) {
			defer wg.Done()

			select {
			case <-time.Tick(pub.timeout):
			case s <- v:
			}
		}(sub)
	}
	wg.Wait()
}

// Close 关闭发布者
func (pub *Publisher) Close() {
	pub.mu.Lock()
	defer pub.mu.Unlock()

	for sub := range pub.subs {
		delete(pub.subs, sub)
		close(sub)
	}
}
