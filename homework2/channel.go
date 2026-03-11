package main

import (
	"fmt"
	"sync"
	"time"
)

// ChannelHomeWork1 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
// 并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
func ChannelHomeWork1() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println(i)
		}
	}()

	wg.Wait()
}

// ChannelHomeWork2 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func ChannelHomeWork2() {
	ch := make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := range ch {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}
