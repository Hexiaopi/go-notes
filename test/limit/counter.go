package main

import (
	"sync"
	"time"
)

type CounterLimiter struct {
	mux     sync.Mutex
	ctime   time.Time
	cycle   time.Duration
	counter int
	limit   int
}

func NewCounterLimiter(limit int, cycle time.Duration) *CounterLimiter {
	return &CounterLimiter{
		mux:     sync.Mutex{},
		ctime:   time.Now(),
		cycle:   cycle,
		counter: 0,
		limit:   limit,
	}
}

func (c *CounterLimiter) Reset() {
	c.ctime = time.Now()
	c.counter = 0
}

func (c *CounterLimiter) Allow() bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	now := time.Now()
	if now.Sub(c.ctime) >= c.cycle {
		c.Reset()
		return true
	}
	if c.counter >= c.limit {
		return false
	}
	c.counter++
	return true
}
