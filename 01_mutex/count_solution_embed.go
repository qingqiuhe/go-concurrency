package main

import (
	"fmt"
	"sync"
)

type Counter struct{
	CounterType int
	Name string

	mu sync.Mutex
	count uint64
}

func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Counter() uint64{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}	

func main() {
	var counter Counter
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 1e5; j++ {
				counter.Incr()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Counter())
}
