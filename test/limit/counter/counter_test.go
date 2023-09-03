package main

import (
	"sync"
	"testing"
	"time"
)

func TestCounterLimiter_Allow(t *testing.T) {
	c := NewCounterLimiter(5, time.Second)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			if allow := c.Allow(); allow {
				t.Log("allow")
				time.Sleep(time.Millisecond * 200) //模拟处理耗时
			} else {
				t.Log("unallow")
			}
		}()
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
}

func TestCounterLimiter_Allow2(t *testing.T) {
	c := NewCounterLimiter(5, time.Second)
	var wg sync.WaitGroup
	wg.Add(10)
	worker := func() {
		defer wg.Done()
		if allow := c.Allow(); allow {
			t.Log("allow")
			time.Sleep(time.Millisecond * 200) //模拟处理耗时
		} else {
			t.Log("unallow")
		}
	}
	time.Sleep(time.Millisecond * 900)
	for i := 0; i < 5; i++ {
		go worker()
	}
	time.Sleep(time.Millisecond * 100)
	for i:=0;i<5;i++{
		go worker()
	}
	time.Sleep(time.Millisecond*900)
	wg.Wait()
}
