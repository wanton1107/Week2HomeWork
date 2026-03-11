package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// LockHomeWork1
// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，
// 每个协程对计数器进行1000次递增操作，最后输出计数器的值
func LockHomeWork1() {
	num := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				num++
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

// LockHomeWork2
// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，
// 最后输出计数器的值
func LockHomeWork2() {
	var num int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&num, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(num)
}
