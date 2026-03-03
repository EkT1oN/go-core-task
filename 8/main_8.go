package main

import (
	"sync/atomic"
)

type CustomWaitGroup struct {
	counter int64
	sem     chan struct{}
}

func (c *CustomWaitGroup) Add(delta int) {
	newCounter := atomic.AddInt64(&c.counter, int64(delta))
	if newCounter == 0 {
		select {
		case c.sem <- struct{}{}:
		default:
		}
	}
}

func (c *CustomWaitGroup) Done() {
	c.Add(-1)
}

func (c *CustomWaitGroup) Wait() {
	if atomic.LoadInt64(&c.counter) == 0 {
		return
	}
	<-c.sem
}

func NewCustomWaitGroup() *CustomWaitGroup {
	return &CustomWaitGroup{
		counter: 0,
		sem:     make(chan struct{}, 1),
	}
}

func main() {

}
